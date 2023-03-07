package sudoku

import (
	"testing"
)

func TestSolve(t *testing.T) {
	puz := "004300209005009001070060043006002087190007400050083000600000105003508690042910300"
	solution := "864371259325849761971265843436192587198657432257483916689734125713528694542916378"
	got, gotErr := Solve(MustNewBoard(puz))
	if gotErr != nil {
		t.Errorf("Solve() got error %v, but we don't want", gotErr)
	}
	if got.String() != solution {
		t.Errorf("Solve() = %v; want = %v", got.String(), solution)
	}
}

func BenchmarkSolve(b *testing.B) {
	puz := "004300209005009001070060043006002087190007400050083000600000105003508690042910300"
	solution := "864371259325849761971265843436192587198657432257483916689734125713528694542916378"
	for n := 0; n <= b.N; n++ {
		got, gotErr := Solve(MustNewBoard(puz))
		if gotErr != nil {
			b.Errorf("Solve() got error %v, but we don't want", gotErr)
		}
		if got.String() != solution {
			b.Errorf("Solve() = %v; want = %v", got.String(), solution)
		}
	}
}
