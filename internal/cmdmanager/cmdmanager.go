package cmdmanager

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type CMDManager struct {
}

func New() *CMDManager {
	return &CMDManager{}
}

func (cmdm CMDManager) ReadLines() ([]string, error) {

	fmt.Println("Please enter prices (one per line). Press Ctrl+D (Unix) or Ctrl+Z (Windows) to end input:")

	var lines []string
	for {
		var line string
		_, err := fmt.Scanln(&line)
		if err != nil {
			break
		}

		lines = append(lines, line)
	}

	return lines, nil
}

func (cmdm CMDManager) WriteJSON(data interface{}) (err error) {
	err = json.NewEncoder(os.Stdout).Encode(data)
	if err != nil {
		return errors.New("Error encoding JSON to file: " + err.Error())
	}

	return nil
}
