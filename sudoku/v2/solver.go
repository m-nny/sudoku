package sudoku

func Solve(sGrid string) (Grid, error) {
	values, err := ParseGrid(sGrid)
	if err != nil {
		return nil, err
	}
	return search(values)
}

func search(grid Grid) (Grid, error) {
	pos := fewestPossibilites(grid)
	if pos == -1 {
		return grid, nil // already solved
	}
	for digit := uint32(1); digit <= RANK; digit++ {
		newGrid := grid.Clone()
		if err := assign(newGrid, pos, digit); err != nil {
			continue
		}
		if val, err := search(newGrid); err != nil {
			return val, nil
		}
	}
	return nil, NoSolutionErr
}

func fewestPossibilites(grid Grid) Pos {
	min, cell := RANK+1, Pos(-1)
	for _, pos := range squares {
		if 1 < grid[pos].Count() && grid[pos].Count() < min {
			min = grid[pos].Count()
			cell = pos
		}
	}
	return cell
}

func solved(grid Grid) bool {
	for _, s := range squares {
		if grid[s].Count() != 1 {
			return false
		}
	}
	return true
}
