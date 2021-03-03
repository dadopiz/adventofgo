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
		g = unique(g)
		res += len(g)
	}

	fmt.Println("result:", res)
}

func getGroups(lines []string) []string {
	groups := []string{}
	group := ""

	for _, line := range lines {
		if line == "" {
			groups = append(groups, group)
			group = line
		} else {
			group += line
		}
	}

	if group != "" {
		groups = append(groups, group)
	}

	return groups
}

func unique(group string) string {
	keys := map[rune]bool{}

	res := ""
	for _, r := range group {
		if _, value := keys[r]; !value {
			keys[r] = true
			res += string(r)
		}
	}

	return res
}
