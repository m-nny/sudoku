package sudoku

import (
	"fmt"
	"strings"

	"github.com/kelindar/bitmap"
	"golang.org/x/exp/slices"
)

const rank = 9
const subrank = 3

var digits = []rune("123456789")
var rows = []rune("ABCDEFGHI")
var cols = digits

type Pos string

var squares = cross(rows, cols)
var unitList [][]Pos
var units = make(map[Pos][][]Pos)
var peers = make(map[Pos][]Pos)

var onesBitmap bitmap.Bitmap

func init() {
	// unitList
	for _, c := range cols {
		unitList = append(unitList, cross(rows, []rune{c}))
	}
	for _, r := range rows {
		unitList = append(unitList, cross([]rune{r}, cols))
	}
	for _, rs := range []string{"ABC", "DEF", "GHI"} {
		for _, cs := range []string{"123", "456", "789"} {
			unitList = append(unitList, cross([]rune(rs), []rune(cs)))
		}
	}

	// units
	for _, s := range squares {
		var sUnits [][]Pos
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

type Grid map[Pos]bitmap.Bitmap

func (grid Grid) Clone() Grid {
	newGrid := make(Grid)
	for pos, b := range grid {
		newGrid[pos] = b.Clone(nil)
	}
	return newGrid
}

func gridValues(sGrid string) map[Pos]uint32 {
	var g []uint32
	for _, c := range sGrid {
		if '0' <= c && c <= '9' {
			g = append(g, uint32(c-'0'))
		} else if c == '.' {
			g = append(g, 0)
		}
	}
	if len(g) != 81 {
		panic("Board should have 81 cells")
	}
	values := make(map[Pos]uint32)
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
			fmt.Fprintf(&sb, "%*s", width, bitmapString(values[Pos(string(r)+string(c))]))
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
