package main

import (
	"adventofgo/utils"
	"fmt"
	"regexp"
	"strconv"
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
			data = data + " " + line
		}
	}
	datas = append(datas, data)

	passports := []string{}
	for _, data := range datas {
		if requiredFields(data) {
			passports = append(passports, data)
		}
	}

	count := 0
	for _, passport := range passports {
		if checkRules(passport) {
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

func checkRules(data string) bool {
	rules := []func(string) bool{
		byrRule,
		iyrRule,
		eyrRule,
		hgtRule,
		hclRule,
		eclRule,
		pidRule,
	}

	for _, rule := range rules {
		if !rule(data) {
			return false
		}
	}

	fmt.Println(data)
	return true
}

func eyrRule(data string) bool {
	return yearCheck("eyr", 2020, 2030, data)
}

func byrRule(data string) bool {
	return yearCheck("byr", 1920, 2002, data)
}

func iyrRule(data string) bool {
	return yearCheck("iyr", 2010, 2020, data)
}

func pidRule(data string) bool {
	re := regexp.MustCompile("pid:\\d{9}")
	return len(re.FindString(data)) > 0
}

func eclRule(data string) bool {
	re := regexp.MustCompile("ecl:(\\D{3})")
	values := re.FindStringSubmatch(data)
	if len(values) != 2 {
		return false
	}

	tokens := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
	for _, t := range tokens {
		if values[1] == t {
			return true
		}
	}

	return false
}

func hgtRule(data string) bool {
	re := regexp.MustCompile("hgt:(\\d{3})cm")
	values := re.FindStringSubmatch(data)
	if len(values) == 2 {
		value, _ := strconv.Atoi(values[1])
		return value >= 150 && value <= 193
	}

	re = regexp.MustCompile("hgt:(\\d{2})in")
	values = re.FindStringSubmatch(data)
	if len(values) == 2 {
		value, _ := strconv.Atoi(values[1])
		return value >= 59 && value <= 76
	}

	return false
}

func hclRule(data string) bool {
	re := regexp.MustCompile("hcl:(#[0-9 a-f]{6})")
	return len(re.FindString(data)) > 0
}

func yearCheck(field string, min int, max int, data string) bool {
	re := regexp.MustCompile(field + ":(\\d{4})")
	values := re.FindStringSubmatch(data)
	if len(values) != 2 {
		return false
	}

	year, _ := strconv.Atoi(values[1])
	return year >= min && year <= max
}
