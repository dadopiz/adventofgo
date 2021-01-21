package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	lines, err := readLines("input.txt")
	assert(err)

	values, err := convertoToIntArray(lines)
	assert(err)

	for i, first := range values {
		for j, second := range values[i+1:] {
			for _, third := range values[j+1:] {
				if first+second+third == 2020 {
					fmt.Println("answer: ", first*second*third)
					return
				}
			}
		}
	}
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(fileName string) ([]string, error) {
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

func convertoToIntArray(lines []string) ([]int, error) {
	values := []int{}

	for _, line := range lines {
		v, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		values = append(values, v)
	}

	return values, nil
}
