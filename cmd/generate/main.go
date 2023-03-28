package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path"

	"github.com/m-nny/sudoku-solver/pkg/parallel_solver"
)

var seedPtr = flag.Int64("seed", -1, "seed")
var NPtr = flag.Int("n", 100, "# of puzzles to generate")
var hintsPtr = flag.Int("hints", 30, "# of hints")
var workersPtr = flag.Int("workers", 8, "# of parallel workers")
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
	results := parallel_solver.Generate(n, hints, workers)
	var puzzles []string
	for _, puzzle := range results {
		if puzzle.Err != nil {
			return puzzles, puzzle.Err
		}
		puzzles = append(puzzles, puzzle.Puzzle)
	}
	return puzzles, nil
}
