package sudoku

import "fmt"

func assign(grid Grid, pos Pos, digit uint32) error {
	for otherValue := uint32(1); otherValue <= RANK; otherValue++ {
		if !grid[pos].Contains(otherValue) || otherValue == digit {
			continue
		}
		if err := eliminate(grid, pos, otherValue); err != nil {
			return err
		}
	}
	return nil
}

func eliminate(grid Grid, pos Pos, digit uint32) error {
	b := grid[pos]
	if !b.Contains(digit) {
		return nil // already eliminated
	}
	b.Remove(digit)
	if b.Count() == 0 {
		return fmt.Errorf("removed last digit at %s", pos)
	} else if b.Count() == 1 { // if square is reduced to one value d, then eliminate d from the peers
		leftValue, _ := b.Min()
		for _, peer := range peers[pos] {
			if err := eliminate(grid, peer, leftValue); err != nil {
				return err
			}
		}
	}
	// if a unit u is reduced to only one place for a value d, then put it there
	for _, unit := range squareUnits[pos] {
		var existsIn []Pos
		for _, pos := range unit {
			if grid[pos].Contains(digit) {
				existsIn = append(existsIn, pos)
			}
		}
		if len(existsIn) == 0 {
			return fmt.Errorf("no options left for %s", pos)
		} else if len(existsIn) == 1 {
			if err := assign(grid, existsIn[0], digit); err != nil {
				return err
			}
		}
	}
	return nil
}
