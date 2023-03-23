import csv

freq = [0 for i in range(81 + 1)]

with open('testdata/sudoku_1k.csv') as f:
  reader = csv.DictReader(f)
  for row in reader:
    puzzle = row['quizzes']
    zeros = list(puzzle).count("0")
    hints = len(puzzle) - zeros
    # print(puzzle, zeros, hints)
    freq[hints] += 1
for hints in range(len(freq)):
  if freq[hints] > 0:
    print(hints, freq[hints])
