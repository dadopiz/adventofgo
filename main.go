package main

import (
	"adventofgo/utils"
	"fmt"
)

const (
	tree  = '#'
	right = 3
	down  = 1
)

type point struct {
	x int
	y int
}

func (p *point) nextStep(size point) {
	p.x += right
	p.y += down
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

	size := sizeOfTreeMap(treeMap)
	position := point{}
	treeCount := 0
	for position.y < size.y {
		if foundTree(treeMap, position) {
			treeCount++
		}
		position.nextStep(size)
	}

	fmt.Println("trees: ", treeCount)
}
