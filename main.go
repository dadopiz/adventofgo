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

	for i, lhs := range lines {
		for _, rhs := range lines[i+1:] {
			if lhs+rhs == 2020 {
				fmt.Println("answer: ", lhs*rhs)
				return
			}
		}
	}
}

func assert(err error) {
	if err != nil {
		panic(err)
	}
}

func readLines(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	lines := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		lines = append(lines, v)
	}
	err = scanner.Err()

	return lines, err
}
