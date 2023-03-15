package sudoku

import "strings"

func assign(grid Grid, pos string, digit rune) Grid {
	otherValues := strings.ReplaceAll(grid[pos], string(digit), "")
	for _, otherValue := range otherValues {
		if val := eliminate(grid, pos, otherValue); val == nil {
			return nil
		}
	}
	return grid
}

func eliminate(grid Grid, pos string, digit rune) Grid {
	if !strings.ContainsRune(grid[pos], digit) {
		return grid // already eliminated
	}
	grid[pos] = strings.ReplaceAll(grid[pos], string(digit), "")
	if len(grid[pos]) == 0 {
		return nil
	} else if len(grid[pos]) == 1 { // if square is reduced to one value d, then eliminate d from the peers
		leftValue := rune(grid[pos][0])
		for _, peer := range peers[pos] {
			if val := eliminate(grid, peer, leftValue); val == nil {
				return nil
			}
		}
	}
	// if a unit u is reduced to only one place for a value d, then put it there
	for _, unit := range units[pos] {
		var existsIn []string
		for _, pos := range unit {
			if strings.ContainsRune(grid[pos], digit) {
				existsIn = append(existsIn, pos)
			}
		}
		if len(existsIn) == 0 {
			return nil
		} else if len(existsIn) == 1 {
			if val := assign(grid, existsIn[0], digit); val == nil {
				return nil
			}
		}
	}
	return grid
}
