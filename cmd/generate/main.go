package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"
	"sync"

	"github.com/schollz/progressbar/v3"

	"github.com/m-nny/sudoku-solver/sudoku/v2"
)

var seedPtr = flag.Int64("seed", -1, "seed")
var NPtr = flag.Int("n", 100, "# of puzzles to generate")
var hintsPtr = flag.Int("hints", 30, "# of hints")
var workersPtr = flag.Int("workers", 8, "# of parallel workers")
var debugPtr = flag.Bool("debug", false, "enable debugging logs")
var outputDirPtr = flag.String("output", "generated", "output directory")

func main() {
	flag.Parse()
	if *seedPtr != -1 {
		rand.Seed(*seedPtr)
	}
	puzzles, err := generate(*NPtr, *hintsPtr, *workersPtr)
	if err != nil {
		fmt.Printf("Could not generate puzzles: %v\n", err)
	}
	outputFile := path.Join(*outputDirPtr, fmt.Sprintf("n%d-hints%d.csv", *NPtr, *hintsPtr))
	if err := saveDataset(outputFile, puzzles); err != nil {
		fmt.Printf("Could not generate puzzles: %v\n", err)
	}
}

func saveDataset(filepath string, puzzles []string) error {
	if err := os.MkdirAll(path.Dir(filepath), 0700); err != nil {
		return err
	}
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	csvWriter.Write([]string{"puzzle"})
	puzzlesTable := make([][]string, len(puzzles))
	for i, puzzle := range puzzles {
		puzzlesTable[i] = []string{puzzle}
	}
	return csvWriter.WriteAll(puzzlesTable)
}

func generate(n, hints, workers int) ([]string, error) {
	var wg sync.WaitGroup
	jobs := make(chan int, n)
	results := make(chan string, n)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go generateWorker(jobs, results, &wg)
	}
	for i := 0; i < n; i++ {
		jobs <- hints
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	bar := progressbar.Default(int64(n), "generate")
	defer bar.Finish()

	var puzzles []string
	for puzzle := range results {
		puzzles = append(puzzles, puzzle)
		bar.Add(1)
	}
	return puzzles, nil
}

func generateWorker(jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for hints := range jobs {
		puzzle, err := sudoku.Generate(hints)
		if err != nil {
			if *debugPtr {
				fmt.Printf("error generating puzzle: %v", err)
			}
			continue
		}
		results <- puzzle
	}
}
