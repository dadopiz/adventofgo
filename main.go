package main

import (
	"adventofgo/utils"
	"fmt"
)

const (
	rows    int  = 128
	columns int  = 8
	front   rune = 'F'
	back    rune = 'B'
	right   rune = 'R'
	left    rune = 'L'
)

func main() {
	lines, err := utils.ReadLines("input.txt")
	utils.Assert(err)

	maxSeatID := 0
	for _, line := range lines {
		row, column := seat(line)
		seatID := seatID(row, column)
		if seatID > maxSeatID {
			maxSeatID = seatID
		}
	}

	fmt.Println("seatID=", maxSeatID)
}

func upperHalf(min int, max int) int {
	return (min + max) / 2
}

func lowerHalf(min int, max int) int {
	return (min + 1 + max) / 2
}

func seatID(row int, column int) int {
	return row*columns + column
}

func seat(line string) (int, int) {
	row := [2]int{0, rows - 1}
	column := [2]int{0, columns - 1}
	for _, c := range line {
		if c == front {
			row[1] = upperHalf(row[0], row[1])
		} else if c == back {
			row[0] = lowerHalf(row[0], row[1])
		} else if c == right {
			column[0] = lowerHalf(column[0], column[1])
		} else if c == left {
			column[1] = upperHalf(column[0], column[1])
		}
	}

	return row[0], column[0]
}
