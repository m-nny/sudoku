package sudoku

import (
	"math/rand"
	"strings"
)

func Generate(hints int) (string, error) {
	grid, err := ParseGrid(strings.Repeat("0", RANK*RANK))
	if err != nil {
		return "", err
	}
	order := rand.Perm(RANK * RANK)
	grid, err = generateRec(grid, order, 0, hints)
	if err != nil {
		return "", err
	}
	return CompactString(grid), nil
}

func generateRec(oldGrid Grid, order []int, i, hints int) (Grid, error) {
	if hints <= 0 {
		return oldGrid, nil
	}
	pos := Pos(order[i])
	for _, digit := range shuffleSlice(oldGrid.Options(pos))[:3] {
		newGrid := oldGrid.Clone()
		if err := assign(newGrid, pos, digit); err != nil {
			continue
		}
		if val, err := generateRec(newGrid, order, i+1, hints-1); err == nil {
			return val, nil
		}
	}
	return nil, NoSolutionErr
}
