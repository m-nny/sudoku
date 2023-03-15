package sudoku

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestInit(t *testing.T) {
	wantSubrank := 3
	if SUBRANK != wantSubrank {
		t.Errorf("subrank = %v; want = %v", SUBRANK, wantSubrank)
	}
	wantRank := wantSubrank * wantSubrank
	if RANK != wantRank {
		t.Errorf("rank = %v; want = %v", RANK, wantRank)
	}
	if len(squares) != RANK*RANK {
		t.Errorf("len(squares) != 81")
	}
	for _, s := range squares {
		if len(squareUnits[s]) != 3 {
			t.Errorf("len(squareUnits[%v]). got=%v, want=%v", s, len(squareUnits[s]), 3)
		}
	}
	for _, s := range squares {
		if len(peers[s]) != 20 {
			t.Errorf("len(peers[%v]). got=%v, want=%v", s, len(peers[s]), 20)
		}
	}
	if diff := cmp.Diff(peers[PosFrom(0, 0)], []Pos{
		PosFrom(0, 1), PosFrom(0, 2), PosFrom(0, 3), PosFrom(0, 4), PosFrom(0, 5), PosFrom(0, 6), PosFrom(0, 7), PosFrom(0, 8),
		PosFrom(1, 0), PosFrom(1, 1), PosFrom(1, 2),
		PosFrom(2, 0), PosFrom(2, 1), PosFrom(2, 2),
		PosFrom(3, 0), PosFrom(4, 0), PosFrom(5, 0), PosFrom(6, 0), PosFrom(7, 0), PosFrom(8, 0),
	}); diff != "" {
		t.Errorf("units[00] returned diff (-want +got):\n%s", diff)
	}
}
