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

func New(inputPath, outputPath string) *FileManager {
	return &FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	if err != nil {
		return nil, errors.New("Error opening file: " + err.Error())
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println("Error closing file:", err)
		}
	}()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, errors.New("Error scanning file: " + err.Error())
	}

	return lines, nil
}

func (fm FileManager) WriteJSON(data interface{}) (err error) {
	file, err := os.Create(fm.OutputFilePath + ".json")
	if err != nil {
		return errors.New("Error creating file: " + err.Error())
	}
	defer func() {
		if cerr := file.Close(); cerr != nil {
			err = errors.New("Error closing file: " + err.Error())
		}
	}()

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		return errors.New("Error encoding JSON to file: " + err.Error())
	}

	return // err was initially nil, so return nil here
}
