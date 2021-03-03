package main

import (
	"adventofgo/utils"
	"fmt"
)

func main() {
	lines, err := utils.ReadLines("input.txt")
	utils.Assert(err)

	groups := getGroups(lines)

	res := 0
	for _, g := range groups {
		res += count(g)
	}

	fmt.Println("result:", res)
}

type group = []string

func getGroups(lines []string) []group {
	groups := []group{}
	g := group{}

	for _, line := range lines {
		if line == "" {
			groups = append(groups, g)
			g = nil
		} else {
			g = append(g, line)
		}
	}

	if g != nil {
		groups = append(groups, g)
	}

	return groups
}

func count(g group) int {
	keys := map[rune]int{}

	for _, p := range g {
		for _, r := range p {
			keys[r]++
		}
	}

	people := len(g)
	res := 0

	for _, v := range keys {
		if v == people {
			res++
		}
	}

	return res
}
