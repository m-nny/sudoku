package sudoku

import (
	"fmt"
	"strconv"
)

type Options = []int

func AllOptions() Options {
	var opts Options
	for i := 1; i <= rank; i++ {
		opts = append(opts, i)
	}
	return opts
}

func OneOption(val int) Options {
	return Options{val}
}

const rank = 9

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
				options[i][j] = OneOption(int(val))
			} else {
				options[i][j] = AllOptions()
			}
		}
	}
	return &Board{
		options: options,
	}, nil
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
