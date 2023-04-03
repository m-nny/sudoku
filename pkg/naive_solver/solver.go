package sudoku

import "golang.org/x/exp/maps"

func Solve(sGrid string) Grid {
	values := ParseGrid(sGrid)
	return search(values)
}

func search(grid Grid) Grid {
	if grid == nil {
		return nil
	}
	n, pos := fewestPossibilites(grid)
	if n == rank+1 {
		return grid // already solved
	}
	for _, digit := range grid[pos] {
		newValues := make(Grid)
		maps.Copy(newValues, grid)
		if val := search(assign(newValues, pos, digit)); val != nil {
			return val
		}
	}
	return nil
}

func fewestPossibilites(grid Grid) (int, string) {
	min, cell := rank+1, ""
	for _, pos := range squares {
		if 1 < len(grid[pos]) && len(grid[pos]) < min {
			min = len(grid[pos])
			cell = pos
		}
	}
	return min, cell
}

func solved(grid Grid) bool {
	for _, s := range squares {
		if len(grid[s]) != 1 {
			return false
		}
	}
	return true
}
