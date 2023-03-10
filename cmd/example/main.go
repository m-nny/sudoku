package main

import (
	"fmt"

	"github.com/m-nny/sudoku-solver/pkg/sudoku"
)

func main() {
	puzzle := "4.....8.5.3..........7......2.....6.....8.4......1.......6.3.7.5..2.....1.4......"
	solution := "864371259325849761971265843436192587198657432257483916689734125713528694542916378"
	if err := solve(puzzle, solution); err != nil {
		fmt.Printf("Could not solve puzzle: %v\n", err)
	} else {
		fmt.Printf("Puzzle sovled successfully\n")
	}
}

func solve(puzzle, solution string) error {
	puz, ok := sudoku.ParseGrid(puzzle)
	if !ok {
		return sudoku.UnknownErr
	}
	fmt.Printf("Puzzle:\n%v\n", sudoku.ValuesString(puz))

	// prop, err := sudoku.Solve(puz)
	// if err != nil {
	// 	return err
	// }

	// if prop.String() != solution {
	// 	fmt.Printf("Incorrect solution.\nwant=\n%vgot=\n%v\n", solution, prop.PrettyString())
	// }
	// fmt.Printf("Found correct solution.\n%v\n", prop.PrettyString())

	return nil
}
