package main

import (
	"adventofgo/utils"
	"fmt"
	"strings"
)

func main() {
	lines, err := utils.ReadLines("input.txt")
	utils.Assert(err)

	datas := []string{}
	data := ""
	for _, line := range lines {
		if line == "" {
			datas = append(datas, data)
			data = ""
		} else {
			data = data + line
		}
	}
	datas = append(datas, data)

	count := 0
	for _, data := range datas {
		if requiredFields(data) {
			count++
		}
	}

	fmt.Println("count: ", count)
}

func requiredFields(data string) bool {
	fields := []string{
		"byr",
		"iyr",
		"eyr",
		"hgt",
		"hcl",
		"ecl",
		"pid",
		//"cid",
	}

	for _, field := range fields {
		if !strings.Contains(data, field) {
			return false
		}
	}

	return true
}
