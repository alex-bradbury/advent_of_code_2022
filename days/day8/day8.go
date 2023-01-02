package day8

import (
	"advent_of_code_2022/utils"
	"fmt"
	"strconv"
)

type Day struct {
	trees [][]int
	dimX  int
	dimY  int
}

func (day *Day) Init(filePath string) {
	println("running Init")
	println("---------------")

	lines := utils.ReadTxtFile(filePath)
	day.trees = make([][]int, len(lines))
	day.dimY = len(lines)
	day.dimX = len(lines[0])
	for i, line := range lines {
		day.trees[i] = make([]int, len(line))
		for j, numRune := range line {
			numInt, err := strconv.Atoi(string(numRune))
			if err != nil {
				panic(err)
			}
			day.trees[i][j] = int(numInt)
		}
	}
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	visibleTrees := (day.dimX * 2) + ((day.dimY - 2) * 2)
	for i, row := range day.trees {
		if i == 0 || i == day.dimY-1 {
			continue
		}
		for j, tree := range row {
			if j == 0 || j == day.dimX-1 {
				continue
			}
			if isVisible(j, i, tree, *day) {
				visibleTrees++
			}
		}
	}

	return fmt.Sprint(visibleTrees)
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	bestVisibility := 0
	for i, row := range day.trees {
		if i == 0 || i == day.dimY-1 {
			continue
		}
		for j, tree := range row {
			if j == 0 || j == day.dimX-1 {
				continue
			}
			visibility := scenicScore(j, i, tree, *day)
			if visibility > bestVisibility {
				bestVisibility = visibility
			}
		}
	}

	return fmt.Sprint(bestVisibility)
}

func isVisible(x int, y int, height int, day Day) bool {
	visTop, visBottom, visLeft, visRight := true, true, true, true

	// check left
	for i := 0; i < x; i++ {
		if day.trees[y][i] >= height {
			visLeft = false
			break
		}
	}

	// check right
	for i := x + 1; i < day.dimX; i++ {
		if day.trees[y][i] >= height {
			visRight = false
			break
		}
	}

	// check top
	for i := 0; i < y; i++ {
		if day.trees[i][x] >= height {
			visTop = false
			break
		}
	}

	// check bottom
	for i := y + 1; i < day.dimY; i++ {
		if day.trees[i][x] >= height {
			visBottom = false
			break
		}
	}

	return visTop || visBottom || visLeft || visRight
}

func scenicScore(x int, y int, height int, day Day) int {
	visTop, visBottom, visLeft, visRight := 0, 0, 0, 0

	// check left
	for i := x - 1; i > -1; i-- {
		visLeft++
		if day.trees[y][i] >= height {
			break
		}
	}

	// check right
	for i := x + 1; i < day.dimX; i++ {
		visRight++
		if day.trees[y][i] >= height {
			break
		}
	}

	// check top
	for i := y - 1; i > -1; i-- {
		visTop++
		if day.trees[i][x] >= height {
			break
		}
	}

	// check bottom
	for i := y + 1; i < day.dimY; i++ {
		visBottom++
		if day.trees[i][x] >= height {
			break
		}
	}

	return visTop * visBottom * visLeft * visRight
}
