package sudoku

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRemoveDuplicates(t *testing.T) {
	want := []Pos{0, 1, 2, 9}
	got := removeDuplicates([]Pos{9, 1, 2, 9, 0})
	if diff := cmp.Diff(got, want); diff != "" {
		t.Errorf("removeUnits(%v) returned: %v; want:%v\n%v", []Pos{9, 1, 2, 9, 0}, got, want, diff)
	}
}
