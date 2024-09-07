package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
)

// ReadCSV reads a CSV file and returns the data as a slice of maps.
// Columns is an array of column names to be used as keys in the map.
func ReadCSV(filePath string, columns []string) ([]map[string]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV: %v", err)
	}

	if len(rows) == 0 {
		return nil, fmt.Errorf("no data found in CSV file")
	}

	header := rows[0]

	columnIndex := make(map[string]int)
	for i, col := range header {
		columnIndex[col] = i
	}

	for _, col := range columns {
		if _, exists := columnIndex[col]; !exists {
			return nil, fmt.Errorf("column %s not found in CSV file", col)
		}
	}

	var result []map[string]string
	for _, row := range rows[1:] { // Skip header row
		data := make(map[string]string)
		for _, col := range columns {
			index := columnIndex[col]
			data[col] = row[index]
		}
		result = append(result, data)
	}

	return result, nil
}

func GetFileNames(folderPath string) ([]string, error) {
	var fileNames []string

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, filepath.Join(folderPath, file.Name()))
		}
	}

	return fileNames, nil
}
