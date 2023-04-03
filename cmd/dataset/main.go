package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/m-nny/sudoku-solver/pkg/parallel"
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
	if !*parallelOnly {
		fmt.Printf("solveSudokus()")
		result, err := solveSudokus(sudokus)
		if err != nil {
			return err
		}
		fmt.Printf("All %d puzzles solved!\ncorrect: %d incorrent: %d errors: %d\n",
			result.total, result.correct, result.incorrect, result.errors)
		fmt.Println()
		if result.errors > 0 {
			return nil
		}
	}

	fmt.Printf("solveSudokusParallel()")
	resultP, err := solveSudokusParallel(sudokus, *workerNumPtr)
	if err != nil {
		return err
	}
	fmt.Printf("workerNum %d\n", *workerNumPtr)
	fmt.Printf("All %d puzzles solved!\ncorrect: %d incorrent: %d errors: %d\n",
		resultP.total, resultP.correct, resultP.incorrect, resultP.errors)
	fmt.Println()

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

func solveSudokus(sudokus [][]string) (SolveSudokusResult, error) {
	var result SolveSudokusResult
	start := time.Now()
	defer func() {
		fmt.Printf("solveSudokus() finished in %v\n", time.Since(start).Truncate(time.Millisecond))
	}()
	result.total = len(sudokus)

	bar := progressbar.Default(int64(result.total))
	for _, entry := range sudokus {
		puzzle, solution := entry[0], entry[1]
		board, err := sudoku.Solve(puzzle)
		if board == nil || err != nil {
			result.errors++
			fmt.Printf("Error solving board:\n%v\n", puzzle)
			return result, nil
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
	return result, nil
}

func solveSudokusParallel(sudokus [][]string, numWorkers int) (SolveSudokusResult, error) {
	result := SolveSudokusResult{total: len(sudokus)}
	var jobs []*parallel.SolveResult
	for i, entry := range sudokus {
		jobs = append(jobs, &parallel.SolveResult{
			Id:              i,
			Puzzle:          entry[0],
			CorrectSolution: entry[1],
		})
	}
	results := parallel.Solve(jobs, numWorkers)
	for _, jobResult := range results {
		if jobResult.Solution == "" || jobResult.Err != nil {
			result.errors++
			fmt.Printf("Error solving board\n%v\n", jobResult.Puzzle)
			return result, jobResult.Err
		} else if jobResult.Solution == jobResult.CorrectSolution {
			result.correct++
		} else {
			fmt.Printf("Found incorrect solution for board:\n%v\n%v\n", jobResult.Puzzle, jobResult.Solution)
			fmt.Printf("Expected:\n%v\n", jobResult.CorrectSolution)
			result.incorrect++
		}
	}
	return result, nil
}
