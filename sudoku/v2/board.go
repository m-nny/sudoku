package sudoku

import (
	"fmt"
	"strings"

	"github.com/kelindar/bitmap"
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

var onesBitmap bitmap.Bitmap

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

	for i := uint32(1); i <= 9; i++ {
		onesBitmap.Set(i)
	}
}

type Grid map[string]bitmap.Bitmap

func (grid Grid) Clone() Grid {
	newGrid := make(Grid)
	for pos, b := range grid {
		newGrid[pos] = b.Clone(nil)
	}
	return newGrid
}

func gridValues(sGrid string) map[string]uint32 {
	var g []uint32
	for _, c := range sGrid {
		if strings.ContainsRune(digits, c) {
			g = append(g, uint32(c-'0'))
		}
		if strings.ContainsRune(".0", c) {
			g = append(g, 0)
		}
	}
	if len(g) != 81 {
		panic("Board should have 81 cells")
	}
	values := make(map[string]uint32)
	for i, s := range squares {
		values[s] = g[i]
	}
	return values
}

func ParseGrid(sGrid string) (Grid, error) {
	grid := make(Grid)
	for _, pos := range squares {
		grid[pos] = onesBitmap.Clone(nil)
	}
	for pos, digit := range gridValues(sGrid) {
		if digit == 0 {
			continue
		}
		if err := assign(grid, pos, digit); err != nil {
			return nil, err
		}
	}
	return grid, nil
}

func PrettyString(values Grid) string {
	var sb strings.Builder
	width := 0
	for _, s := range squares {
		if values[s].Count() > width {
			width = values[s].Count() + 1
		}
	}
	width++
	line := strings.Join(repeatString(strings.Repeat("-", width*3), 3), "+")
	for i, r := range rows {
		for j, c := range cols {
			fmt.Fprintf(&sb, "%*s", width, bitmapString(values[string(r)+string(c)]))
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
		if values[s].Count() == 1 {
			digit, _ := values[s].Min()
			fmt.Fprintf(&sb, "%v", digit)
		} else {
			fmt.Fprintf(&sb, "0")
		}
	}
	return sb.String()
}
