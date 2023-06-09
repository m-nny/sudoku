package parallel

import (
	"fmt"
	"sync"

	"github.com/m-nny/sudoku-solver/pkg/sudoku"
	"github.com/schollz/progressbar/v3"
)

type GenerateResult struct {
	Puzzle string
	Err    error
}

func Generate(n, hints, workers int) []*GenerateResult {
	var wg sync.WaitGroup
	jobs := make(chan int, n)
	results := make(chan *GenerateResult, n)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go generateWorker(jobs, results, &wg)
	}
	for i := 0; i < n; i++ {
		jobs <- hints
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(results)
	}()

	bar := progressbar.Default(int64(n), fmt.Sprintf("Generate h:%d", hints))
	defer bar.Finish()

	var puzzles []*GenerateResult
	for puzzle := range results {
		puzzles = append(puzzles, puzzle)
		bar.Add(1)
	}
	return puzzles
}

func generateWorker(jobs <-chan int, results chan<- *GenerateResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for hints := range jobs {
		puzzle, err := sudoku.Generate(hints)
		results <- &GenerateResult{puzzle, err}
	}
}
