package day6

import (
	"advent_of_code_2022/utils"
	"fmt"
	"strings"
)

type Day struct {
	input string
}

func (day *Day) Init(filePath string) {
	lines := utils.ReadTxtFile(filePath)
	day.input = lines[0]
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	return fmt.Sprint(findPattern(day.input, 4))
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	return fmt.Sprint(findPattern(day.input, 14))
}

func findPattern(input string, patternLength int) int {
	result := 0
	inputResetCount := 0
	workingInput := input
main:
	for {
		set := make(map[rune]bool)
		for _, letter := range workingInput {
			if set[letter] {
				inputResetCount += strings.Index(workingInput, string(letter)) + 1
				workingInput = workingInput[strings.Index(workingInput, string(letter))+1:]
				break
			}
			set[letter] = true
			if len(set) == patternLength {
				result = patternLength + inputResetCount
				break main
			}
		}
	}
	return result
}
