package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/schollz/progressbar/v3"

	"github.com/m-nny/sudoku-solver/sudoku/v2"
)

var NPtr = flag.Int("n", 100, "# of puzzles to generate")
var hintsPtr = flag.Int("hints", 30, "# of hints")
var debugPtr = flag.Bool("debug", false, "enable debugging logs")

func main() {
	flag.Parse()
	n := *NPtr
	hints := *hintsPtr
	puzzles, err := generate(n, hints)
	if err != nil {
		fmt.Printf("Could not generate puzzles: %v\n", err)
	}
	if err := solve(puzzles); err != nil {
		fmt.Printf("Could not solve generated puzzles: %v\n", err)
	}
}

type Stats struct {
	noSolution        int
	mutlipleSolutions int
	ok                int
}

func generate(n, hints int) ([]string, error) {
	bar := progressbar.Default(int64(n), "generate")
	defer bar.Finish()
	var puzzles []string
	for i := 0; i < n; i++ {
		puzzle, err := sudoku.Generate(hints)
		if err != nil {
			return nil, err
		}
		puzzles = append(puzzles, puzzle)
		bar.Add(1)
	}
	return puzzles, nil
}
func solve(puzzles []string) error {
	bar := progressbar.Default(int64(len(puzzles)), "solve")
	defer bar.Finish()
	var stats Stats
	for _, puzzle := range puzzles {
		if *debugPtr {
			fmt.Printf("solving: %s\n%s\n", puzzle, sudoku.PrettySudoku(puzzle))
		}
		_, err := sudoku.Solve(puzzle)
		if errors.Is(err, sudoku.NoSolutionErr) {
			stats.noSolution++
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
