package sudoku

import (
	"fmt"
	"strings"

	"github.com/kelindar/bitmap"
	"golang.org/x/exp/slices"
)

func removeDuplicates(s []Pos) []Pos {
	slices.Sort(s)
	return slices.Compact(s)
}

func repeatString(val string, n int) []string {
	var res []string
	for i := 0; i < n; i++ {
		res = append(res, val)
	}
	return res
}

func bitmapString(b bitmap.Bitmap) string {
	var sb strings.Builder
	b.Range(func(digit uint32) {
		fmt.Fprintf(&sb, "%d", digit)
	})
	return sb.String()
}
