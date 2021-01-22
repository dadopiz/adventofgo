package utils

import (
	"bufio"
	"os"
)

// Assert check error and exit
func Assert(err error) {
	if err != nil {
		panic(err)
	}
}

// ReadLines parses a file and put each line into an array
func ReadLines(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	lines := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	err = scanner.Err()

	return lines, err
}
