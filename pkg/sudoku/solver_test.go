package sudoku

import (
	"testing"
)

func TestSolve(t *testing.T) {
	testCase := []struct {
		name     string
		puzzle   string
		solution string
		wantErr  bool
	}{
		{
			name:     "second example",
			puzzle:   "000230050673400800005007900310780000064009200082546009008000103000801700000600405",
			solution: "941238657673495821825167934319782546564319278782546319258974163436851792197623485",
			wantErr:  false,
		},
		{
			name:     "first example",
			puzzle:   "004300209005009001070060043006002087190007400050083000600000105003508690042910300",
			solution: "864371259325849761971265843436192587198657432257483916689734125713528694542916378",
			wantErr:  false,
		},
	}
	for _, test := range testCase {
		t.Run(test.name, func(t *testing.T) {
			// got, gotErr := Solve(MustNewBoard(test.puzzle))
			// if gotErr != nil {
			// 	t.Errorf("Solve() got error %v, but we don't want", gotErr)
			// }
			// if got.String() != test.solution {
			// 	t.Logf("got:\n%v\n", got.PrettyString())
			// 	t.Errorf("Solve() = %v; want = %v", got.String(), test.solution)
			// }
		})
	}
}
