package sudoku

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRank(t *testing.T) {
	wantSubrank := 3
	if subrank != wantSubrank {
		t.Errorf("subrank = %v; want = %v", subrank, wantSubrank)
	}
	wantRank := wantSubrank * wantSubrank
	if rank != wantRank {
		t.Errorf("rank = %v; want = %v", rank, wantRank)
	}
}

func TestInit(t *testing.T) {
	if len(squares) != 81 {
		t.Errorf("len(squares) != 81")
	}
	if len(unitList) != 27 {
		t.Errorf("len(squares) != 27")
	}
	for _, s := range squares {
		if len(units[s]) != 3 {
			t.Errorf("len(units[%v]). got=%v, want=%v", s, len(units[s]), 3)
		}
	}
	for _, s := range squares {
		if len(peers[s]) != 20 {
			t.Errorf("len(peers[%v]). got=%v, want=%v", s, len(peers[s]), 20)
		}
	}
	if diff := cmp.Diff(units["C2"], [][]string{
		{"A2", "B2", "C2", "D2", "E2", "F2", "G2", "H2", "I2"},
		{"C1", "C2", "C3", "C4", "C5", "C6", "C7", "C8", "C9"},
		{"A1", "A2", "A3", "B1", "B2", "B3", "C1", "C2", "C3"},
	}); diff != "" {
		t.Errorf("units[C2] returned diff (-want +got):\n%s", diff)
	}
}
