package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"

	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

	if err != nil {
		return nil, errors.New("Could not open file: " + err.Error())
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		return nil, errors.New("Reading the file content failed: " + err.Error())
	}


	file.Close()
	return lines, nil
}

func WriteJSON(path string, data interface{}) error {
	file, err :=os.Create(path) 
	if err != nil {
		return errors.New("Could not create file: " + err.Error())
	}	
	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		file.Close()
		return errors.New("Could not encode data to JSON: " + err.Error())
	}
	file.Close()
	return nil	
}