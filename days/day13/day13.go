package day13

import (
	"advent_of_code_2022/utils"
	"fmt"
	"regexp"
	"strconv"
)

type Day struct {
	pairs []pair
}

type pair struct {
	left, right [][]any
}

func (day *Day) Init(filePath string) {
	println("running Init")
	println("---------------")

	lines := utils.ReadTxtFile(filePath)
	var currentPair pair
	prevLine := "empty"
	for _, line := range lines {
		if len(line) > 0 {
			switch prevLine {
			case "empty":
				currentPair = pair{}
				currentPair.left = parseList(line)
			case "[":
				currentPair.right = parseList(line)
				day.pairs = append(day.pairs, currentPair)
			}
			prevLine = "["
		} else {
			prevLine = "empty"
		}
	}
	fmt.Println(day.pairs)
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	total := 0
	for i, pair := range day.pairs {
		if correctOrder(pair) {
			total += i
		}
	}

	return fmt.Sprint(total)
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	return fmt.Sprint(0)
}

func parseList(line string) [][]any {
	regex, _ := regexp.Compile("[0-9]+")
	fullList := make([][]any, 0)
	currentListStack := utils.ArrayStack{}
	currentListStack.Push(make([]any, 0))
	inList := false
	for i, letter := range line {
		if i == 0 {
			continue
		}
		if i == len(line)-1 {
			continue
		}
		letterStr := string(letter)
		if regex.Match([]byte(letterStr)) {
			num, _ := strconv.Atoi(letterStr)
			currentList, _ := currentListStack.ReadTop()
			*currentList = append(*currentList, num)
			if !inList {
				fullList = append(fullList, *currentList)
				currentListStack.Pop()
			}
		} else if letter == '[' {
			inList = true
		} else if letter == ']' {
			inList = false
			currentList, _ := currentListStack.Pop()
			if currentListStack.IsEmpty() {
				fullList = append(fullList, currentList)
			} else {
				nextTop, _ := currentListStack.ReadTop()
				*nextTop = append(*nextTop, currentList)
			}
			currentList = make([]any, 0)
		}
	}

	return fullList
}

func correctOrder(pairToCompare pair) bool {
	correctOrder := true
	shorter := pairToCompare.left
	if len(pairToCompare.left) > len(pairToCompare.right) {
		shorter = pairToCompare.right
	}

	// for i := range shorter {

	// }

	return correctOrder
}
