package day3

import (
	"advent_of_code_2022/utils"
	"strings"
	"unicode"
)

type Day3 struct {
	lines []string
}

func (day *Day3) Init(filePath string) {
	day.lines = utils.ReadTxtFile(filePath)
}

func (day *Day3) Task1() int {
	println("running task 1!")
	println("---------------")

	var commonChars []rune
	for _, line := range day.lines {
		firstHalf := line[:len(line)/2]
		secondHalf := line[len(line)/2:]

		for _, char := range secondHalf {
			if strings.ContainsRune(firstHalf, char) {
				commonChars = append(commonChars, char)
				break
			}
		}
	}

	return getPriorityTotal(commonChars)
}

func (day *Day3) Task2() int {
	println("running task 2!")
	println("---------------")

	var badges []rune
	for i := 0; i < len(day.lines)-2; i += 3 {
		workingSet := map[rune]bool{}
		for _, char := range day.lines[i+1] {
			if strings.ContainsRune(day.lines[i], char) {
				workingSet[char] = true
			}
		}
		for _, char := range day.lines[i+2] {
			if _, ok := workingSet[char]; ok {
				badges = append(badges, char)
				break
			}
		}
	}

	return getPriorityTotal(badges)
}

func getPriorityTotal(chars []rune) int {
	total := 0
	for _, char := range chars {
		if unicode.IsLower(char) {
			total += int(char) - 96
		} else {
			total += int(char) - 38
		}
	}
	return total
}
