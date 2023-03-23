package main

import (
	"errors"
	"fmt"

	"github.com/schollz/progressbar/v3"

	"github.com/m-nny/sudoku-solver/sudoku/v2"
)

func main() {
	n := 100*1000
	hints := 30
	if err := generate(n, hints); err != nil {
		fmt.Printf("Could not generate puzzles: %v\n", err)
	}
}

type Stats struct {
	noSolution        int
	mutlipleSolutions int
	ok                int
}

func generate(n, hints int) error {
	bar := progressbar.Default(int64(n))
	defer bar.Finish()
	var puzzles []string
	bar.Describe("generate")
	for i := 0; i < n; i++ {
		puzzle, err := sudoku.Generate(hints)
		if err != nil {
			return err
		}
		puzzles = append(puzzles, puzzle)
		bar.Add(1)
	}
	bar.Reset()
	var stats Stats
	bar.Describe("solve")
	for _, puzzle := range puzzles {
		_, err := sudoku.Solve(puzzle)
		if errors.Is(err, sudoku.NoSolutionErr) {
			stats.noSolution++
			// fmt.Printf("no solution: %s\n%s\n", puzzle, sudoku.PrettySudoku(puzzle))
			// return err
		} else if errors.Is(err, sudoku.MultipleSolutionsErr) {
			stats.noSolution++
		} else if err != nil {
			return err
		} else {
			stats.ok++
		}
		bar.Add(1)
	}
	fmt.Printf("total %d\nno solution %d\nmultiple solutions %d\n ok %d\n", len(puzzles), stats.noSolution, stats.mutlipleSolutions, stats.ok)
	return nil
}
