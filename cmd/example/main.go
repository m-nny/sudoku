package main

import (
	"fmt"

	"github.com/m-nny/sudoku-solver/pkg/sudoku"
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
	puz, err := sudoku.NewBoard(puzzle)
	if err != nil {
		return err
	}
	fmt.Printf("Puzzle:\n%v\n", puz.PrettyString())

	prop, err := sudoku.Solve(puz)
	if err != nil {
		return err
	}

	if prop.String() != solution {
		fmt.Printf("Incorrect solution.\nwant=\n%vgot=\n%v\n", solution, prop.PrettyString())
	}
	fmt.Printf("Found correct solution.\n%v\n", prop.PrettyString())

	return nil
}
