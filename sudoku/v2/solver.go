package sudoku

func Solve(sGrid string) (Grid, error) {
	values, err := ParseGrid(sGrid)
	if err != nil {
		return nil, err
	}
	return search(values)
}

func search(grid Grid) (Grid, error) {
	n, pos := fewestPossibilites(grid)
	if n == rank+1 {
		return grid, nil // already solved
	}
	for digit := uint32(1); digit <= 9; digit++ {
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

func fewestPossibilites(grid Grid) (int, string) {
	min, cell := rank+1, ""
	for _, pos := range squares {
		if 1 < grid[pos].Count() && grid[pos].Count() < min {
			min = grid[pos].Count()
			cell = pos
		}
	}
	return min, cell
}

func solved(grid Grid) bool {
	for _, s := range squares {
		if grid[s].Count() != 1 {
			return false
		}
	}
	return true
}
