package main

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/m-nny/sudoku-solver/pkg/sudoku"
)

func main() {
	dataset := "testdata/sudoku_1k.csv"
	err := solveDataset(dataset)
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

	for i, entry := range sudokus {
		if i%10 == 0 {
			fmt.Printf("[%d/%d] correct: %d incorrent: %d errors: %d\n",
				i, len(sudokus), correct, incorrect, errors)
		}

		puzzle, solution := entry[0], entry[1]
		board, err := sudoku.NewBoard(puzzle)
		if err != nil {
			errors++
			fmt.Printf("Error reading board:\n%v\n%v\n", board.PrettyString(), err)
			break
		}
		prop, err := sudoku.Solve(board)
		if err != nil {
			errors++
			fmt.Printf("Error solving board:\n%v\n%v\n", board.PrettyString(), err)
			break
		}
		if prop.Match(solution) {
			correct++
		} else {
			fmt.Printf("Found incorrect solution for board:\n%v\n%v\n", puzzle, board.PrettyString())
			fmt.Printf("Expected:\n%v\n%v\n", solution, sudoku.MustNewBoard(solution).PrettyString())
			incorrect++
			break
		}
	}
	fmt.Printf("All %d puzzles solved!\ncorrect: %d incorrent: %d errors: %d\n",
		len(sudokus), correct, incorrect, errors)
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
