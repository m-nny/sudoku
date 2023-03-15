package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/m-nny/sudoku-solver/pkg/sudoku"
	"github.com/schollz/progressbar/v3"
)

var datasetPtr = flag.String("dataset", "testdata/sudoku_1k.csv", "csv dataset")
var workerNumPtr = flag.Int("workers", 4, "# of workers")
var parallelOnly = flag.Bool("parallel-only", true, "skip sequential sequential solve")

func main() {
	flag.Parse()

	err := solveDataset(*datasetPtr)
	if err != nil {
		fmt.Printf("Could not solve puzzle: %v\n", err)
	}
}

func solveDataset(dataset string) error {
	sudokus, err := readDataset(dataset)
	if err != nil {
		return err
	}
	fmt.Printf("solveSudokusParallel()")
	resultP := solveSudokusParallel(sudokus, *workerNumPtr)
	fmt.Printf("workerNum %d\n", *workerNumPtr)
	fmt.Printf("All %d puzzles solved!\ncorrect: %d incorrent: %d errors: %d\n",
		resultP.total, resultP.correct, resultP.incorrect, resultP.errors)
	fmt.Println()

	if !*parallelOnly {
		fmt.Printf("solveSudokus()")
		result := solveSudokus(sudokus)
		fmt.Printf("All %d puzzles solved!\ncorrect: %d incorrent: %d errors: %d\n",
			result.total, result.correct, result.incorrect, result.errors)
		fmt.Println()
	}

	return nil
}

func readDataset(filepath string) ([][]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	if _, err = csvReader.Read(); err != nil {
		return nil, err
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

type SolveSudokusResult struct {
	errors    int
	correct   int
	incorrect int
	total     int
}

func solveSudokus(sudokus [][]string) SolveSudokusResult {
	var result SolveSudokusResult
	start := time.Now()
	defer func() {
		fmt.Printf("solveSudokus() finished in %v\n", time.Since(start))
	}()
	result.total = len(sudokus)

	bar := progressbar.Default(int64(result.total))
	for _, entry := range sudokus {
		puzzle, solution := entry[0], entry[1]
		board := sudoku.Solve(puzzle)
		if board == nil {
			result.errors++
			fmt.Printf("Error reading board\n")
			break
		}
		if sudoku.CompactString(board) == solution {
			result.correct++
		} else {
			fmt.Printf("Found incorrect solution for board:\n%v\n%v\n", puzzle, sudoku.PrettyString(board))
			fmt.Printf("Expected:\n%v\n", solution)
			result.incorrect++
			break
		}
		bar.Add(1)
	}
	return result
}

type JobResult struct {
	id       int
	puzzle   string
	solution string
	proposal string
	took     time.Duration
}

func solveSudokusParallel(sudokus [][]string, numWorkers int) SolveSudokusResult {
	var result SolveSudokusResult
	start := time.Now()
	defer func() {
		fmt.Printf("solveSudokusParallel() finished in %v\n", time.Since(start))
	}()
	result.total = len(sudokus)

	var wg sync.WaitGroup
	jobs := make(chan *JobResult, result.total)
	results := make(chan *JobResult, result.total)

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go solveSudokuWorker(jobs, results, &wg)
	}

	for id, entry := range sudokus {
		jobs <- &JobResult{
			id:       id + 1,
			puzzle:   entry[0],
			solution: entry[1],
		}
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	bar := progressbar.Default(int64(result.total))
	for jobResult := range results {
		if jobResult.proposal == "" {
			result.errors++
			fmt.Printf("Error reading board\n")
		} else if jobResult.proposal == jobResult.solution {
			result.correct++
		} else {
			fmt.Printf("Found incorrect solution for board:\n%v\n%v\n", jobResult.puzzle, jobResult.proposal)
			fmt.Printf("Expected:\n%v\n", jobResult.solution)
			result.incorrect++
		}
		bar.Add(1)
	}
	return result
}

func solveSudokuWorker(jobs <-chan *JobResult, results chan<- *JobResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// fmt.Printf("Solving %d\n", job.id)
		start := time.Now()
		board := sudoku.Solve(job.puzzle)
		if board != nil {
			job.proposal = sudoku.CompactString(board)
		}
		job.took = time.Since(start)
		results <- job
	}
}
