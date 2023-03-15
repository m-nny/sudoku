package sudoku

import (
	"fmt"
	"strings"

	"golang.org/x/exp/slices"
)

const rank = 9
const subrank = 3

const digits = "123456789"
const rows = "ABCDEFGHI"
const cols = digits

var squares = cross(rows, cols)
var unitList [][]string
var units = make(map[string][][]string)
var peers = make(map[string][]string)

func init() {
	// unitList
	for _, c := range cols {
		unitList = append(unitList, cross(rows, string(c)))
	}
	for _, r := range rows {
		unitList = append(unitList, cross(string(r), cols))
	}
	for _, rs := range []string{"ABC", "DEF", "GHI"} {
		for _, cs := range []string{"123", "456", "789"} {
			unitList = append(unitList, cross(rs, cs))
		}
	}

	// units
	for _, s := range squares {
		var sUnits [][]string
		for _, u := range unitList {
			if slices.Contains(u, s) {
				sUnits = append(sUnits, u)
			}
		}
		units[s] = sUnits
	}

	// peers
	for _, s := range squares {
		peers[s] = uniqueWithout(units[s], s)
	}
}

type Grid = map[string]string

func gridValues(grid string) Grid {
	var sb strings.Builder
	for _, c := range grid {
		if strings.ContainsRune(digits, c) || strings.ContainsRune(".0", c) {
			fmt.Fprintf(&sb, "%v", string(c))
		}
	}
	g := sb.String()
	if len(g) != 81 {
		panic("Board should have 81 cells")
	}
	values := make(Grid)
	for i, s := range squares {
		values[s] = string(g[i])
	}
	return values
}

func ParseGrid(sGrid string) Grid {
	grid := make(Grid)
	for _, pos := range squares {
		grid[pos] = digits
	}
	for pos, digit := range gridValues(sGrid) {
		if digit == "0" {
			continue
		}
		if val := assign(grid, pos, rune(digit[0])); val == nil {
			return nil
		}
	}
	return grid
}

func PrettyString(values Grid) string {
	var sb strings.Builder
	width := 0
	for _, s := range squares {
		if len(values[s]) > width {
			width = len(values[s]) + 1
		}
	}
	width++
	line := strings.Join(repeatString(strings.Repeat("-", width*3), 3), "+")
	for i, r := range rows {
		for j, c := range cols {
			fmt.Fprintf(&sb, "%*s", width, values[string(r)+string(c)])
			if (j+1)%subrank == 0 && j+1 < rank {
				fmt.Fprint(&sb, "|")
			}
		}
		fmt.Fprintln(&sb)
		if (i+1)%subrank == 0 && (i+1) < rank {
			fmt.Fprintln(&sb, line)
		}
	}
	return sb.String()
}

func CompactString(values Grid) string {
	var sb strings.Builder
	for _, s := range squares {
		if len(values[s]) == 1 {
			fmt.Fprintf(&sb, "%v", values[s])
		} else {
			fmt.Fprintf(&sb, "0")
		}
	}
	return sb.String()
}
