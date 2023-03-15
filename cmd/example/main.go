package main

import (
	"fmt"

	"github.com/m-nny/sudoku-solver/sudoku/v1"
)

func main() {
	puzzle := "004300209005009001070060043006002087190007400050083000600000105003508690042910300"
	solution := "864371259325849761971265843436192587198657432257483916689734125713528694542916378"
	if err := solve(puzzle, solution); err != nil {
		fmt.Printf("Could not solve puzzle: %v\n", err)
	} else {
		fmt.Printf("Puzzle sovled successfully\n")
	}
}

func solve(puzzle, solution string) error {
	puz := sudoku.Solve(puzzle)
	if puz == nil {
		return sudoku.UnknownErr
	}
	fmt.Printf("Puzzle:\n%v\n", sudoku.PrettyString(puz))

	fmt.Printf("Found solution.\n%v\n", sudoku.PrettyString(puz))
	if sudoku.CompactString(puz) != solution {
		fmt.Printf("Solve() = %v; want = %v", sudoku.CompactString(puz), solution)
	}

	return nil
}
