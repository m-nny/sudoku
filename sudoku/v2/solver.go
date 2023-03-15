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
	for _, digit := range grid.Options(pos) {
		newGrid := grid.Clone()
		if err := assign(newGrid, pos, digit); err != nil {
			continue
		}
		if val, err := search(newGrid); err == nil {
			return val, nil
		}
	}
	return nil, NoSolutionErr
}

func fewestPossibilites(grid Grid) Pos {
	min, cell := RANK+1, Pos(-1)
	for _, pos := range squares {
		posCount := grid[pos].Count()
		if posCount == 2 {
			return pos
		}
		if 1 < posCount && posCount < min {
			min = posCount
			cell = pos
		}
	}
	return cell
}
