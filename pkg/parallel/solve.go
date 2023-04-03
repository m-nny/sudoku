package parallel

import (
	"fmt"
	"sync"
	"time"

	"github.com/m-nny/sudoku-solver/pkg/sudoku"
	"github.com/schollz/progressbar/v3"
)

type SolveResult struct {
	Id              int
	Puzzle          string
	Solution        string
	CorrectSolution string
	Took            time.Duration
	Err             error
}

func Solve(sudokus []*SolveResult, numWorkers int) []*SolveResult {
	start := time.Now()
	defer func() {
		fmt.Printf("Solve(%d) finished in %v\n", numWorkers, time.Since(start).Truncate(time.Millisecond))
	}()

	var wg sync.WaitGroup
	jobs := make(chan *SolveResult, len(sudokus))
	jobResults := make(chan *SolveResult, len(sudokus))

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go solveSudokuWorker(jobs, jobResults, &wg)
	}

	for _, entry := range sudokus {
		jobs <- entry
	}
	close(jobs)
	go func() {
		wg.Wait()
		close(jobResults)
	}()

	bar := progressbar.Default(int64(len(sudokus)), "Solve")
	defer bar.Finish()
	var result []*SolveResult
	for jobResult := range jobResults {
		bar.Add(1)
		result = append(result, jobResult)
	}
	return result
}

func solveSudokuWorker(jobs <-chan *SolveResult, jobResults chan<- *SolveResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		start := time.Now()
		board, err := sudoku.Solve(job.Puzzle)
		if board != nil && err == nil {
			job.Solution = sudoku.CompactString(board)
		}
		job.Err = err
		job.Took = time.Since(start)
		jobResults <- job
	}
}
