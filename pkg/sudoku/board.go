package sudoku

import (
	"fmt"
	"strconv"
	"strings"
)

type Options = []int

func allOptions() Options {
	var opts Options
	for i := 1; i <= rank; i++ {
		opts = append(opts, i)
	}
	return opts
}

func oneOption(val int) Options {
	return Options{val}
}

func Remove(opts Options, val int) Options {
	idx := -1
	for i := range opts {
		if opts[i] == val {
			idx = i
		}
	}
	if idx != -1 {
		opts = append(opts[0:idx], opts[idx+1:]...)
		return opts
	}
	return opts
}

const rank = 9
const subrank = 3

type Board struct {
	options [][]Options
}

func NewBoard(b string) (*Board, error) {
	if len(b) < rank*rank {
		return nil, fmt.Errorf("board should have at least %d digits", rank*rank)
	}
	options := make([][]Options, rank)
	for i := range options {
		options[i] = make([]Options, rank)
		for j := range options[i] {
			str_pos := i*rank + j
			val, err := strconv.ParseInt(string(b[str_pos]), 10, 32)
			if err != nil {
				return nil, err
			}
			if val != 0 {
				options[i][j] = oneOption(int(val))
			} else {
				options[i][j] = allOptions()
			}
		}
	}
	return &Board{
		options: options,
	}, nil
}

func MustNewBoard(b string)*Board {
	board, err := NewBoard(b)
	if err != nil {
		return nil
	}
	return board
}

func (b *Board) Valid() bool {
	for i := range b.options {
		for j := range b.options[i] {
			if len(b.options[i][j]) == 0 {
				return false
			}
		}
	}
	return true
}

func (b *Board) Solved() bool {
	for i := range b.options {
		for j := range b.options[i] {
			if len(b.options[i][j]) != 1 {
				return false
			}
		}
	}
	return true
}

func (b *Board) Copy() *Board {
	options := make([][]Options, rank)
	for i := range options {
		options[i] = make([]Options, rank)
		for j := range options[i] {
			for _, val := range b.options[i][j] {
				options[i][j] = append(options[i][j], val)
			}
		}
	}
	return &Board{options: options}
}

func (b *Board) String() string {
	var sb strings.Builder
	for _, row := range b.options {
		for _, cell := range row {
			if len(cell) == 1 {
				fmt.Fprintf(&sb, "%d", cell[0])
			} else {
				fmt.Fprintf(&sb, "0")
			}
		}
	}
	return sb.String()
}

func (b *Board) PrettyString() string {
	var sb strings.Builder
	for i, row := range b.options {
		for j, cell := range row {
			if len(cell) == 1 {
				fmt.Fprintf(&sb, "%d ", cell[0])
			} else {
				fmt.Fprintf(&sb, "0 ")
			}
			if (j+1)%subrank == 0 && j+1 < rank {
				fmt.Fprintf(&sb, "| ")
			}
		}
		fmt.Fprintln(&sb)
		if (i+1)%subrank == 0 && i+1 < rank {
			fmt.Fprintln(&sb, strings.Repeat("-", (rank + 2) * 2))
		}
	}
	return sb.String()
}

func (b*Board) Match(solution string) bool {
	if !b.Solved() || !b.Valid() {
		return false;
	}
	return b.String() == solution
}
