package days

import (
	"advent_of_code_2022/utils"
	"sort"
	"strconv"
)

type Day1 struct{}

func (day Day1) Task1(filePath string) int {
	println("running task 1!")
	println("---------------")

	lines := utils.ReadTxtFile(filePath)

	max := 0
	runningTotal := 0
	for index, line := range lines {
		if len(line) != 0 {
			lineNum, err := strconv.Atoi(line)

			if err != nil {
				println("Input line is not a number")
				return 0
			}

			runningTotal += lineNum

			if index == len(lines)-1 {
				if runningTotal > max {
					max = runningTotal
				}
			}
		} else {
			if runningTotal > max {
				max = runningTotal
			}
			runningTotal = 0
		}
	}

	return max
}

func (day Day1) Task2(filePath string) int {
	println("running task 2!")
	println("---------------")

	lines := utils.ReadTxtFile(filePath)

	var totals []int
	runningTotal := 0
	for index, line := range lines {
		if len(line) != 0 {
			lineNum, err := strconv.Atoi(line)

			if err != nil {
				println("Input line is not a number")
				return 0
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

	sort.Ints(totals)
	totalsLen := len(totals)

	return totals[totalsLen-1] + totals[totalsLen-2] + totals[totalsLen-3]
}
