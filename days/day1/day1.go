package day1

import (
	"advent_of_code_2022/utils"
	"fmt"
	"sort"
	"strconv"
)

type Day1 struct {
	Totals []int
}

func (day *Day1) Init(filePath string) {
	lines := utils.ReadTxtFile(filePath)

	var totals []int
	runningTotal := 0
	for index, line := range lines {
		if len(line) != 0 {
			lineNum, err := strconv.Atoi(line)

			if err != nil {
				println("Input line is not a number")
			}

			runningTotal += lineNum

			if index == len(lines)-1 {
				totals = append(totals, runningTotal)
			}
		} else {
			totals = append(totals, runningTotal)
			runningTotal = 0
		}
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	day.Totals = totals
}

func (day *Day1) Task1() string {
	println("running task 1!")
	println("---------------")

	return fmt.Sprint(day.Totals[0])
}

func (day *Day1) Task2() string {
	println("running task 2!")
	println("---------------")

	return fmt.Sprint(day.Totals[0] + day.Totals[1] + day.Totals[2])
}
