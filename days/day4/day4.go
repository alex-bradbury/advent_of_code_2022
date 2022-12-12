package day4

import (
	"advent_of_code_2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day4 struct {
	elfPairs []elfPair
}

type elfPair struct {
	first, second []int
}

func (day *Day4) Init(filePath string) {
	lines := utils.ReadTxtFile(filePath)

	var elfPairs []elfPair
	for _, line := range lines {
		pair := strings.Split(line, ",")
		var elfPair elfPair

		min1, err := strconv.Atoi((strings.Split(pair[0], "-")[0]))
		if err != nil {
			println("Error converting to int")
		}
		max1, err := strconv.Atoi((strings.Split(pair[0], "-")[1]))
		if err != nil {
			println("Error converting to int")
		}
		elfPair.first = makeRange(min1, max1)

		min2, err := strconv.Atoi((strings.Split(pair[1], "-")[0]))
		if err != nil {
			println("Error converting to int")
		}
		max2, err := strconv.Atoi((strings.Split(pair[1], "-")[1]))
		if err != nil {
			println("Error converting to int")
		}
		elfPair.second = makeRange(min2, max2)

		elfPairs = append(elfPairs, elfPair)
	}
	day.elfPairs = elfPairs
}

func (day *Day4) Task1() string {
	println("running task 1!")
	println("---------------")

	containedPairs := 0
	for _, pair := range day.elfPairs {
		contained := true
		var smaller, bigger []int
		if len(pair.first) > len(pair.second) {
			bigger = pair.first
			smaller = pair.second
		} else {
			bigger = pair.second
			smaller = pair.first
		}

		for _, num := range smaller {
			if !intInSlice(num, bigger) {
				contained = false
				break
			}
		}
		if contained {
			containedPairs += 1
		}
	}

	return fmt.Sprint(containedPairs)
}

func (day *Day4) Task2() string {
	println("running task 2!")
	println("---------------")

	overlappingPairs := 0
	for _, pair := range day.elfPairs {
		overlapped := false
		var smaller, bigger []int
		if len(pair.first) > len(pair.second) {
			bigger = pair.first
			smaller = pair.second
		} else {
			bigger = pair.second
			smaller = pair.first
		}

		for _, num := range smaller {
			if intInSlice(num, bigger) {
				overlapped = true
				break
			}
		}
		if overlapped {
			overlappingPairs += 1
		}
	}

	return fmt.Sprint(overlappingPairs)
}

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func intInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
