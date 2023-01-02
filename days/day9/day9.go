package day9

import (
	"advent_of_code_2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day struct {
	instructions []instruction
}

type instruction struct {
	direction string
	amount    int
}

type position struct {
	x int
	y int
}

func (day *Day) Init(filePath string) {
	println("running Init")
	println("---------------")

	lines := utils.ReadTxtFile(filePath)
	day.instructions = make([]instruction, 0)
	for _, line := range lines {
		parts := strings.Split(line, " ")
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		day.instructions = append(day.instructions, instruction{parts[0], amount})
	}
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	visitedPositions := make(map[position]bool)
	headPosition, tailPosition := position{0, 0}, position{0, 0}

	for _, instruction := range day.instructions {
		for i := 0; i < instruction.amount; i++ {
			moveHead(&headPosition, instruction.direction)
			moveKnot(&tailPosition, headPosition)
			visitedPositions[tailPosition] = true
		}
	}

	return fmt.Sprint(len(visitedPositions))
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	const NUMBER_OF_KNOTS int = 10
	visitedPositions := make(map[position]bool)
	knotPositions := make([]position, NUMBER_OF_KNOTS)

	for _, instruction := range day.instructions {
		for i := 0; i < instruction.amount; i++ {
			moveHead(&knotPositions[0], instruction.direction)
			for j := 1; j < NUMBER_OF_KNOTS; j++ {
				moveKnot(&knotPositions[j], knotPositions[j-1])
			}
			visitedPositions[knotPositions[NUMBER_OF_KNOTS-1]] = true
		}
	}

	return fmt.Sprint(len(visitedPositions))
}

func moveHead(headPosition *position, direction string) {
	switch direction {
	case "L":
		headPosition.x--
	case "R":
		headPosition.x++
	case "D":
		headPosition.y--
	case "U":
		headPosition.y++
	}
}

func moveKnot(knotPosition *position, infrontPosition position) {
	hy, hx := infrontPosition.y, infrontPosition.x
	ty, tx := knotPosition.y, knotPosition.x
	if hx == tx+2 {
		if hy > ty {
			knotPosition.y++
		} else if hy < ty {
			knotPosition.y--
		}
		knotPosition.x++
	} else if hx == tx-2 {
		if hy > ty {
			knotPosition.y++
		} else if hy < ty {
			knotPosition.y--
		}
		knotPosition.x--
	} else if hy == ty+2 {
		if hx > tx {
			knotPosition.x++
		} else if hx < tx {
			knotPosition.x--
		}
		knotPosition.y++
	} else if hy == ty-2 {
		if hx > tx {
			knotPosition.x++
		} else if hx < tx {
			knotPosition.x--
		}
		knotPosition.y--
	}
}
