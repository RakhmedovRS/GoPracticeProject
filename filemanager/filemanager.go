package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to open file %s\n", fm.InputFilePath, err.Error()))
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Unable to read data from file %s\n", fm.InputFilePath, err.Error()))
	}

	return lines, nil
}

func (fm FileManager) WriteDataAsJson(data any) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		return errors.New(fmt.Sprintf("Unable to open file %s\n", fm.OutputFilePath))
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.Encode(data)
	return nil
}

func New(inputFileName, outputFileName string) FileManager {
	return FileManager{
		InputFilePath:  inputFileName,
		OutputFilePath: outputFileName,
	}
}
