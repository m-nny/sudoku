package sudoku

import (
	"fmt"
	"strings"

	"github.com/kelindar/bitmap"
)

func cross(a, b string) []string {
	var prod []string
	for _, i := range a {
		for _, j := range b {
			prod = append(prod, string(i)+string(j))
		}
	}
	return prod
}

func uniqueWithout(s [][]string, val string) []string {
	exists := make(map[string]bool)
	for _, slice := range s {
		for _, val := range slice {
			exists[val] = true
		}
	}
	exists[val] = false
	var allValues []string
	for val, ok := range exists {
		if ok {
			allValues = append(allValues, val)
		}
	}
	return allValues
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
