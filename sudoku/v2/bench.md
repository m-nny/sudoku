```bash
$ go run cmd/dataset/main.go --dataset testdata/sudoku_100k.csv  --workers 16
 100% |█████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████████| (100000/100000, 15513 it/s)
solveSudokusParallel() finished in 6.473806968s
workerNum 16
All 100000 puzzles solved!
correct: 100000 incorrent: 0 errors: 0

```
