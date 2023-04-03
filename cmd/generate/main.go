package main

import (
	"errors"
	"flag"
	"fmt"
	"math"
	"path"
	"time"

	"github.com/m-nny/sudoku-solver/pkg/dataset"
	"github.com/m-nny/sudoku-solver/pkg/parallel"
	"github.com/m-nny/sudoku-solver/pkg/sudoku"
)

var NPtr = flag.Int("n", 100, "# of puzzles to generate")
var minHintsPtr = flag.Int("min-hints", 30, "min # of hints")
var maxHintsPtr = flag.Int("max-hints", 30, "max # of hints")
var workersPtr = flag.Int("workers", 8, "# of parallel workers")
var outputDirPtr = flag.String("output", "generated", "output directory")
var noGeneratePtr = flag.Bool("no-generate", false, "do not generate dataset")
var solvePtr = flag.Bool("solve", false, "solve dataset")

func main() {
	flag.Parse()
	for hints := *minHintsPtr; hints <= *maxHintsPtr; hints++ {
		filename := path.Join(*outputDirPtr, fmt.Sprintf("n%d-hints%d.csv", *NPtr, hints))
		if !*noGeneratePtr {
			if err := generateAndSave(filename, *NPtr, hints, *workersPtr); err != nil {
				fmt.Printf("Could not generate puzzles: %v\n", err)
			}
		}
		if *solvePtr {
			if err := loadAndSolve(filename, *workersPtr); err != nil {
				fmt.Printf("Could not solve puzzles: %v\n", err)
			}
		}
	}
}

func generateAndSave(datasetFile string, n, hints, workers int) error {
	results := parallel.Generate(n, hints, workers)
	var puzzles []string
	for _, puzzle := range results {
		if puzzle.Err != nil {
			return puzzle.Err
		}
		puzzles = append(puzzles, puzzle.Puzzle)
	}
	if err := dataset.SaveUnsolved(datasetFile, puzzles); err != nil {
		fmt.Printf("Could not generate puzzles: %v\n", err)
	}
	return nil
}

func loadAndSolve(datasetFile string, workers int) error {
	puzzles, err := dataset.ReadUnsolved(datasetFile)
	if err != nil {
		return nil
	}
	var sudokus []*parallel.SolveResult
	for id, puzzle := range puzzles {
		sudokus = append(sudokus, &parallel.SolveResult{
			Id:     id,
			Puzzle: puzzle,
		})
	}
	sudokus = parallel.Solve(sudokus, workers)
	fmt.Println(datasetFile)
	stats := map[string]stat{
		"no-solution":        {math.MaxInt, math.MinInt, 0, 0},
		"multiple-solutions": {math.MaxInt, math.MinInt, 0, 0},
		"ok":                 {math.MaxInt, math.MinInt, 0, 0},
	}
	for _, puzzle := range sudokus {
		tag := "ok"
		if errors.Is(puzzle.Err, sudoku.NoSolutionErr) {
			tag = "no-solution"
		} else if errors.Is(puzzle.Err, sudoku.MultipleSolutionsErr) {
			tag = "multiple-solutions"
		} else if puzzle.Err != nil {
			return puzzle.Err
		}
		stats[tag] = stats[tag].upd(puzzle.Took)
	}

	fmt.Println()
	fmt.Printf("%19s %6s %10s %10s %10s\n", "tag", "cnt", "min", "max", "avg")
	for _, tag := range []string{"ok", "no-solution", "multiple-solutions"} {
		s := stats[tag]
		var avg time.Duration
		if s.cnt != 0 {
			avg = time.Duration(int(s.total) / s.cnt)
		}
		fmt.Printf("%19s %6d %10s %10s %10s\n", tag, s.cnt, s.min.Truncate(time.Microsecond), s.max.Truncate(time.Microsecond), avg.Truncate(time.Microsecond))
	}
	fmt.Println()
	return nil
}

type stat struct {
	min   time.Duration
	max   time.Duration
	total time.Duration
	cnt   int
}

func (s stat) upd(val time.Duration) stat {
	if val < s.min {
		s.min = val
	}
	if s.max < val {
		s.max = val
	}
	s.total += val
	s.cnt++
	return s
}
