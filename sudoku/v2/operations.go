package sudoku

import "fmt"

func assign(grid Grid, pos Pos, digit uint32) error {
	for _, otherValue := range grid.Options(pos) {
		if otherValue == digit {
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
	// if a unit is reduced to only one place for a value *digit**, then put it there
	for _, unit := range squareUnits[pos] {
		existsIn, ok := Pos(-1), true
		for _, pos := range unit {
			if grid[pos].Contains(digit) {
				if existsIn != -1 {
					ok = false
					break
				}
				existsIn = pos
			}
		}
		if existsIn == -1 {
			return fmt.Errorf("no options left for %s", pos)
		}
		if !ok {
			continue
		}
		if err := assign(grid, existsIn, digit); err != nil {
			return err
		}
	}
	return nil
}
