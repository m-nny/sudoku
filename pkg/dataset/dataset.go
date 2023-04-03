package dataset

import (
	"encoding/csv"
	"os"
	"path"
)

func ReadUnsolved(filepath string) ([]string, error) {
	records, err := ReadSolved(filepath)
	if err != nil {
		return nil, err
	}
	puzzles := make([]string, len(records))
	for i, record := range records {
		puzzles[i] = record[0]
	}

	return puzzles, nil
}
func ReadSolved(filepath string) ([][]string, error) {
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	if _, err = csvReader.Read(); err != nil {
		return nil, err
	}
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}

func SaveUnsolved(filepath string, puzzles []string) error {
	puzzlesTable := make([][]string, len(puzzles))
	for i, puzzle := range puzzles {
		puzzlesTable[i] = []string{puzzle}
	}
	return SaveSolved(filepath, puzzlesTable)
}

func SaveSolved(filepath string, puzzles [][]string) error {
	if err := os.MkdirAll(path.Dir(filepath), 0700); err != nil {
		return err
	}
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	csvWriter := csv.NewWriter(f)
	csvWriter.Write([]string{"puzzle", "solution"})
	return csvWriter.WriteAll(puzzles)
}
