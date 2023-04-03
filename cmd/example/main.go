package main

import (
	"fmt"

	"github.com/m-nny/sudoku-solver/pkg/solver"
)

func main() {
	puzzle := "000000000000010280200034000020000740906000300140203000708000001300000000000000000"
	solution := "671892534534617289289534176823961745956748312147253968798326451315489627462175893"
	if err := solve(puzzle, solution); err != nil {
		fmt.Printf("Could not solve puzzle: %v\n", err)
	} else {
		fmt.Printf("Puzzle sovled successfully\n")
	}
}

func solve(puzzle, solution string) error {
	puz, err := sudoku.Solve(puzzle)
	if err != nil {
		return err
	}
	fmt.Printf("Found solution.\n%v\n", sudoku.PrettyString(puz))
	if sudoku.CompactString(puz) != solution {
		fmt.Printf("Solve() = %v; want = %v", sudoku.CompactString(puz), solution)
	}

	return nil
}
