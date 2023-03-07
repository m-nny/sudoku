package sudoku

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestRank(t *testing.T) {
	want := 9
	if rank != want {
		t.Errorf("rank = %v; want = %v", rank, want)
	}
}

func TestOneOption(t *testing.T) {
	want := Options{1}
	got := OneOption(1)
	if !cmp.Equal(got, want) {
		t.Errorf("OneOption() = %v; want = %v", got, want)
	}
}
func TestAllOptions(t *testing.T) {
	want := Options{1, 2, 3, 4, 5, 6, 7, 8, 9}
	got := AllOptions()
	if !cmp.Equal(got, want) {
		t.Errorf("AllOptions() = %v; want = %v", got, want)
	}
}

func TestNewBoard(t *testing.T) {
	testCases := []struct {
		name      string
		b         string
		want      *Board
		wantError bool
	}{
		{
			name: "Half empty board",
			// 004|300|209
			// 005|009|001
			// 070|060|043
			// ---+---+---
			// 006|002|087
			// 190|007|400
			// 050|083|000
			// ---+---+---
			// 600|000|105
			// 003|508|690
			// 042|910|300
			b: "004300209005009001070060043006002087190007400050083000600000105003508690042910300",
			want: &Board{
				options: [][]Options{
					{
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{4},
						{3},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{2},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{9},
					},
					{
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{5},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1},
					},
					{
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{7},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{6},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{4},
						{3},
					},
					{
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{6},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{2},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{8},
						{7},
					},
					{
						{1},
						{9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{7},
						{4},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
					},
					{
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{5},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{8},
						{3},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
					},
					{
						{6},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{5},
					},
					{
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{3},
						{5},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{8},
						{6},
						{9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
					},
					{
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{4},
						{2},
						{9},
						{1},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{3},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
						{1, 2, 3, 4, 5, 6, 7, 8, 9},
					},
				},
			},
			wantError: false,
		},
		{
			name:      "Too small board",
			b:         "123",
			want:      nil,
			wantError: true,
		},
	}
	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			got, gotErr := NewBoard(test.b)
			if !test.wantError && gotErr != nil {
				t.Errorf("NewBoard() got error %v, but we don't want", gotErr)
			}
			if diff := cmp.Diff(got, test.want, cmp.AllowUnexported(Board{})); diff != "" {
				t.Errorf("newBoard() returned diff (-want +got):\n%s", diff)
			}
		})
	}
}
