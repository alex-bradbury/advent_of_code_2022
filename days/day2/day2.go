package day2

import (
	"advent_of_code_2022/utils"
	"strings"
)

type Day2 struct {
	pairs []resultPair
}

type resultPair struct {
	opponent string
	you      string
}

func (day *Day2) Init(filePath string) {
	lines := utils.ReadTxtFile(filePath)

	var pairs []resultPair
	for _, line := range lines {
		parts := strings.Fields(line)
		pairs = append(pairs, resultPair{parts[0], parts[1]})
	}

	day.pairs = pairs
}

func (day *Day2) Task1() int {
	println("running task 1!")
	println("---------------")

	points := 0
	for _, pair := range day.pairs {
		switch pair.you {
		case "X":
			points += 1
		case "Y":
			points += 2
		case "Z":
			points += 3
		}

		points += pair.getPoints()
	}

	return points
}

func (day *Day2) Task2() int {
	println("running task 2!")
	println("---------------")

	points := 0
	for _, pair := range day.pairs {
		switch pair.you {
		case "X":
			points += 0
		case "Y":
			points += 3
		case "Z":
			points += 6
		}

		points += pair.whatToPlay()
	}

	return points
}

func (pair resultPair) getPoints() int {
	switch pair.opponent {
	case "A":
		switch pair.you {
		case "X":
			return 3
		case "Y":
			return 6
		case "Z":
			return 0
		}
		return 0
	case "B":
		switch pair.you {
		case "X":
			return 0
		case "Y":
			return 3
		case "Z":
			return 6
		}
		return 0
	case "C":
		switch pair.you {
		case "X":
			return 6
		case "Y":
			return 0
		case "Z":
			return 3
		}
		return 0
	default:
		return 0
	}
}

func (pair resultPair) whatToPlay() int {
	switch pair.opponent {
	case "A":
		switch pair.you {
		case "X":
			return 3
		case "Y":
			return 1
		case "Z":
			return 2
		}
		return 0
	case "B":
		switch pair.you {
		case "X":
			return 1
		case "Y":
			return 2
		case "Z":
			return 3
		}
		return 0
	case "C":
		switch pair.you {
		case "X":
			return 2
		case "Y":
			return 3
		case "Z":
			return 1
		}
		return 0
	default:
		return 0
	}
}
