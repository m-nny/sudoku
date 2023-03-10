package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/m-nny/sudoku-solver/pkg/sudoku"
)

func main() {
	datasetPtr := flag.String("dataset", "testdata/sudoku_1k.csv", "csv dataset")

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

	var correct, incorrect, errors int
	start := time.Now()

	for i, entry := range sudokus {
		if i%100 == 0 {
			fmt.Printf("[%d/%d] correct: %d incorrent: %d errors: %d\n",
				i, len(sudokus), correct, incorrect, errors)
			fmt.Printf("Time spent %v\n", time.Since(start))
		}

		puzzle, solution := entry[0], entry[1]
		board := sudoku.Solve(puzzle)
		if board == nil {
			errors++
			fmt.Printf("Error reading board:\n%v\n", err)
			break
		}
		if sudoku.CompactString(board) == solution {
			correct++
		} else {
			fmt.Printf("Found incorrect solution for board:\n%v\n%v\n", puzzle, sudoku.PrettyString(board))
			fmt.Printf("Expected:\n%v\n", solution)
			incorrect++
			break
		}
	}
	fmt.Printf("All %d puzzles solved!\ncorrect: %d incorrent: %d errors: %d\n",
		len(sudokus), correct, incorrect, errors)
	fmt.Printf("Total time spent %v\n", time.Since(start))
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
