package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	lines, err := readLines("input.txt")
	assert(err)

	count := 0
	for _, line := range lines {
		res, err := createPolicyAndPassword(line)
		assert(err)

		if res.isValid() {
			count++
		}
	}

	fmt.Println("Valid passwords: ", count)
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

type policy struct {
	firstPos  int
	secondPos int
	letter    rune
}

type policyAndPassword struct {
	policy   policy
	password string
}

func createPolicyAndPassword(line string) (*policyAndPassword, error) {
	re := regexp.MustCompile("(\\d+)-(\\d+)\\s(\\w):\\s(\\w+)")
	parts := re.FindStringSubmatch(line)

	firstPos, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	secondPos, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	res := new(policyAndPassword)
	res.policy.firstPos = firstPos - 1
	res.policy.secondPos = secondPos - 1
	res.policy.letter = []rune(parts[3])[0]
	res.password = parts[4]
	return res, nil
}

func (ptr *policyAndPassword) isValid() bool {
	runes := []rune(ptr.password)
	letterInFirstPos := runes[ptr.policy.firstPos] == ptr.policy.letter
	letterInSecondPos := runes[ptr.policy.secondPos] == ptr.policy.letter
	return letterInFirstPos != letterInSecondPos
}
