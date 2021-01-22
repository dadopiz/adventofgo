package main

import (
	"adventofgo/utils"
	"fmt"
)

const (
	tree = '#'
)

type point struct {
	x int
	y int
}

func (p *point) nextStep(slope point, size point) {
	p.x += slope.x
	p.y += slope.y
	if p.x >= size.x {
		p.x = p.x - size.x
	}
}

func sizeOfTreeMap(treeMap []string) point {
	return point{
		x: len(treeMap[0]),
		y: len(treeMap),
	}
}

func foundTree(treeMap []string, p point) bool {
	return []rune(treeMap[p.y])[p.x] == tree
}

func main() {
	treeMap, err := utils.ReadLines("input.txt")
	utils.Assert(err)

	slopes := []point{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}
	size := sizeOfTreeMap(treeMap)
	result := 1

	for _, slope := range slopes {
		treeCount := 0
		position := point{}

		for position.y < size.y {
			if foundTree(treeMap, position) {
				treeCount++
			}
			position.nextStep(slope, size)
		}
		fmt.Println("trees: ", treeCount)

		result = result * treeCount
	}

	fmt.Println("result: ", result)
}
