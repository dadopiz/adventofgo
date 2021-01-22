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

type policy struct {
	min    int
	max    int
	letter rune
}

type policyAndPassword struct {
	policy   policy
	password string
}

func createPolicyAndPassword(line string) (*policyAndPassword, error) {
	re := regexp.MustCompile("(\\d+)-(\\d+)\\s(\\w):\\s(\\w+)")
	parts := re.FindStringSubmatch(line)

	min, err := strconv.Atoi(parts[1])
	if err != nil {
		return nil, err
	}

	max, err := strconv.Atoi(parts[2])
	if err != nil {
		return nil, err
	}

	res := new(policyAndPassword)
	res.policy.min = min
	res.policy.max = max
	res.policy.letter = []rune(parts[3])[0]
	res.password = parts[4]
	return res, nil
}

func (ptr *policyAndPassword) isValid() bool {
	count := 0
	for _, l := range ptr.password {
		if l == ptr.policy.letter {
			count++
		}
	}

	return count >= ptr.policy.min && count <= ptr.policy.max
}
