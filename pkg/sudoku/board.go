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

type Values = map[string]string

func ParseGrid(grid string) Values {
	values := make(Values)
	for _, s := range squares {
		values[s] = digits
	}
	for s, d := range gridValues(grid) {
		if !strings.Contains(digits, d) {
			continue
		}
		if val := assign(values, s, d); val == nil {
			return nil
		}
	}
	return values
}

func gridValues(grid string) map[string]string {
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
	values := make(map[string]string)
	for i, s := range squares {
		values[s] = string(g[i])
	}
	return values
}

func assign(values Values, s, d string) Values {
	otherValues := strings.ReplaceAll(values[s], d, "")
	for _, d2 := range otherValues {
		if val := eliminate(values, s, string(d2)); val == nil {
			return nil
		}
	}
	return values
}

func eliminate(values Values, s string, d string) Values {
	if !strings.Contains(values[s], d) {
		return values // already eliminated
	}
	values[s] = strings.ReplaceAll(values[s], string(d), "")
	// if square is reduced to one value d, then eliminate d from the peers
	if len(values[s]) == 0 {
		return nil
	} else if len(values[s]) == 1 {
		d2 := values[s]
		for _, s2 := range peers[s] {
			if val := eliminate(values, s2, d2); val == nil {
				return nil
			}
		}
	}
	// if a unit u is reduced to only one place for a value d, then put it there
	for _, u := range units[s] {
		var dplaces []string
		for _, s := range u {
			if strings.Contains(values[s], d) {
				dplaces = append(dplaces, s)
			}
		}
		if len(dplaces) == 0 {
			return nil
		} else if len(dplaces) == 1 {
			if val := assign(values, dplaces[0], d); val == nil {
				return nil
			}
		}
	}
	return values
}

func PrettyString(values Values) string {
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

func CompactString(values Values) string {
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
