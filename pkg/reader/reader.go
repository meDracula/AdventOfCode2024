package reader

import (
	"bufio"
	"os"
)

func FileReadlines(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}, nil
	}

	defer file.Close()

	filescanner := bufio.NewScanner(file)
	filescanner.Split(bufio.ScanLines)
	var content []string

	for filescanner.Scan() {
		content = append(content, filescanner.Text())
	}

	return content, nil
}

func FileScanner(filename string) (*bufio.Scanner, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return bufio.NewScanner(file), nil
}
