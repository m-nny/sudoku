package sudoku

import "golang.org/x/exp/maps"

func Solve(grid string) Values {
	values := ParseGrid(grid)
	return search(values)
}

func search(values Values) Values {
	if values == nil {
		return nil
	}
	n, s := fewestPossibilites(values)
	if n == rank+1 {
		return values // already solved
	}
	for _, d := range values[s] {
		newValues := make(Values)
		maps.Copy(newValues, values)
		if val := search(assign(newValues, s, string(d))); val != nil {
			return val
		}
	}
	return nil
}

func fewestPossibilites(values Values) (int, string) {
	min, cell := rank+1, ""
	for _, s := range squares {
		if len(values[s]) > 1 && len(values[s]) < min {
			min = len(values[s])
			cell = s
		}
	}
	return min, cell
}

func solved(values Values) bool {
	for _, s := range squares {
		if len(values[s]) != 1 {
			return false
		}
	}
	return true
}
