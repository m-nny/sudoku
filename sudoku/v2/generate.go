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
	for i := 0; i < hints; i++ {
		pos := Pos(order[i])
		options := grid.Options(pos)
		shuffleSlice(options)
		// fmt.Printf("[%d] pos: %d\n%v\n", i, pos, PrettyString(grid))
		for _, digit := range options {
			newGrid := grid.Clone()
			if err := assign(newGrid, pos, digit); err == nil {
				grid = newGrid
				break
			}
		}
	}
	return CompactString(grid), nil
}

func shuffleSlice(slice []uint32) []uint32 {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
