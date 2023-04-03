package sudoku

import "errors"

var InvalidBoardErr = errors.New("board is not valid")
var NoSolutionErr = errors.New("board has no solutions")
var UnknownErr = errors.New("unknown")
