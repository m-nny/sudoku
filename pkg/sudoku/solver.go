package sudoku

import (
	"errors"
)

func Solve(b *Board) (*Board, error) {
	// fmt.Printf("Solving:\n%v\n", b.PrettyString())
	if !b.Valid() {
		return nil, InvalidBoardErr
	}
	if b.Solved() {
		return b.Copy(), nil
	}
	i, j := FirstEmpty(b)
	// Shouldn't be true, because we checked if board is solved before
	if i == -1 || j == -1 {
		return nil, UnknownErr
	}
	for _, val := range b.options[i][j] {
		prop, err := Solve(assign(b, i, j, val))
		if err == nil {
			return prop, nil
		}
		if errors.Is(err, UnknownErr) {
			return nil, err
		}
	}
	return nil, NoSolutionErr
}

func assign(b *Board, i, j, val int) *Board {
	b = b.Copy()
	for ii := 0; ii < rank; ii++ {
		b.options[ii][j] = Remove(b.options[ii][j], val)
	}
	for jj := 0; jj < rank; jj++ {
		b.options[i][jj] = Remove(b.options[i][jj], val)
	}
	i0, j0 := (i/subrank)*subrank, (j/subrank)*subrank
	for di := 0; di < subrank; di++ {
		for dj := 0; dj < subrank; dj++ {
			ii, jj := i0+di, j0+dj
			b.options[ii][jj] = Remove(b.options[ii][jj], val)
		}
	}
	b.options[i][j] = oneOption(val)
	return b
}

func FirstEmpty(b *Board) (int, int) {
	for i := range b.options {
		for j := range b.options[i] {
			if len(b.options[i][j]) > 1 {
				return i, j
			}
		}
	}
	return -1, -1
}
