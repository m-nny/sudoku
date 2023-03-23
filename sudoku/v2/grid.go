package sudoku

import (
	"fmt"
	"strings"

	"github.com/kelindar/bitmap"
	"golang.org/x/exp/slices"
)

const RANK = 9
const SUBRANK = 3

type Pos int

func PosFrom(i, j int) Pos {
	return Pos(i*RANK + j)
}
func (pos Pos) IJ() (int, int) {
	i, j := int(pos)/RANK, int(pos)%RANK
	return i, j
}
func (pos Pos) String() string {
	i, j := pos.IJ()
	return fmt.Sprintf("%dx%d", i, j)
}

type Unit []Pos

var squares []Pos
var squareUnits = make([][]Unit, RANK*RANK)
var peers = make([][]Pos, RANK*RANK)

var onesBitmap bitmap.Bitmap

func init() {
	for i := 0; i < RANK; i++ {
		for j := 0; j < RANK; j++ {
			squares = append(squares, PosFrom(i, j))
		}
	}
	var allUnits []Unit
	// rowUnits
	for i := 0; i < RANK; i++ {
		var rowUnit Unit
		for j := 0; j < RANK; j++ {
			rowUnit = append(rowUnit, PosFrom(i, j))
		}
		allUnits = append(allUnits, rowUnit)
	}
	// colUnits
	for j := 0; j < RANK; j++ {
		var colUnit Unit
		for i := 0; i < RANK; i++ {
			colUnit = append(colUnit, PosFrom(i, j))
		}
		allUnits = append(allUnits, colUnit)
	}
	// blockUnits
	for i0 := 0; i0 < RANK; i0 += SUBRANK {
		for j0 := 0; j0 < RANK; j0 += SUBRANK {
			var blockUnit Unit
			for i := i0; i < i0+SUBRANK; i++ {
				for j := j0; j < j0+SUBRANK; j++ {
					blockUnit = append(blockUnit, PosFrom(i, j))
				}
			}
			allUnits = append(allUnits, blockUnit)
		}
	}
	// peers && squareUnits
	for _, pos := range squares {
		for _, unit := range allUnits {
			if !slices.Contains(unit, pos) {
				continue
			}
			squareUnits[pos] = append(squareUnits[pos], unit)
			for _, peer := range unit {
				if peer != pos {
					peers[pos] = append(peers[pos], peer)
				}
			}
		}
	}
	for _, pos := range squares {
		peers[pos] = removeDuplicates(peers[pos])
	}

	for i := uint32(1); i <= RANK; i++ {
		onesBitmap.Set(i)
	}
}

type Grid []bitmap.Bitmap // Slice of dimention [RANK*RANK]

func (grid Grid) Clone() Grid {
	newGrid := make(Grid, RANK*RANK)
	for pos, b := range grid {
		newGrid[pos] = b.Clone(nil)
	}
	return newGrid
}

func (grid Grid) Options(pos Pos) []uint32 {
	options := make([]uint32, 0, RANK)
	grid[pos].Range(func(x uint32) { options = append(options, x) })
	return options
}

func gridValues(sGrid string) []uint32 {
	g := make([]uint32, 0, RANK*RANK)
	for _, c := range sGrid {
		if '0' <= c && c <= '9' {
			g = append(g, uint32(c-'0'))
		} else if c == '.' {
			g = append(g, 0)
		}
	}
	if len(g) != RANK*RANK {
		panic("Board should have 81 cells")
	}
	return g
}

func ParseGrid(sGrid string) (Grid, error) {
	grid := make(Grid, RANK*RANK)
	for _, pos := range squares {
		grid[pos] = onesBitmap.Clone(nil)
	}
	for pos, digit := range gridValues(sGrid) {
		if digit == 0 {
			continue
		}
		if err := assign(grid, Pos(pos), digit); err != nil {
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
	for i := 0; i < RANK; i++ {
		for j := 0; j < RANK; j++ {
			fmt.Fprintf(&sb, "%*s", width, bitmapString(values[PosFrom(i, j)]))
			if (j+1)%SUBRANK == 0 && j+1 < RANK {
				fmt.Fprint(&sb, "|")
			}
		}
		fmt.Fprintln(&sb)
		if (i+1)%SUBRANK == 0 && (i+1) < RANK {
			fmt.Fprintln(&sb, line)
		}
	}
	return sb.String()
}

func PrettySudoku(sGrid string) string {
	var sb strings.Builder
	width := 2
	line := strings.Join(repeatString(strings.Repeat("-", width*3), 3), "+")
	for i := 0; i < RANK; i++ {
		for j := 0; j < RANK; j++ {
			fmt.Fprintf(&sb, "%*s", width, string(sGrid[PosFrom(i, j)]))
			if (j+1)%SUBRANK == 0 && j+1 < RANK {
				fmt.Fprint(&sb, "|")
			}
		}
		fmt.Fprintln(&sb)
		if (i+1)%SUBRANK == 0 && (i+1) < RANK {
			fmt.Fprintln(&sb, line)
		}
	}
	return sb.String()
}

func CompactString(values Grid) string {
	if values == nil {
		return "<nil>"
	}
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
