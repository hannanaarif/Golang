package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManger struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManger) Readfile() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file!")
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("failed to reading the file")
	}

	file.Close()

	return lines, nil

}

func (fm FileManger) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		file.Close()
		return errors.New("failed to convert data into json")
	}
	file.Close()
	return nil
}

func New(inputPath, outputPath string) FileManger {
	return FileManger{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}
