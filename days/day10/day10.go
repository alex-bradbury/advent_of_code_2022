package day10

import (
	"advent_of_code_2022/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day struct {
	xValues []int
}

func (day *Day) Init(filePath string) {
	println("running Init")
	println("---------------")

	lines := utils.ReadTxtFile(filePath)
	day.xValues = []int{1}
	for _, line := range lines {
		currentValue := day.xValues[len(day.xValues)-1]
		if line == "noop" {
			day.xValues = append(day.xValues, currentValue)
		} else {
			parts := strings.Split(line, " ")
			valueToAdd, _ := strconv.Atoi(parts[1])
			day.xValues = append(day.xValues, currentValue)
			day.xValues = append(day.xValues, currentValue+valueToAdd)
		}
	}
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	total := 0
	for i := 19; i < 220; i += 40 {
		total += day.xValues[i] * (i + 1)
	}

	return fmt.Sprint(total)
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	crtDrawing := make([]string, 240)
	for i := 0; i < 240; i++ {
		if spriteVisible(i, day.xValues) {
			crtDrawing[i] = "#"
		} else {
			crtDrawing[i] = "."
		}
	}

	fmt.Println(crtDrawing[:40])
	fmt.Println(crtDrawing[40:80])
	fmt.Println(crtDrawing[80:120])
	fmt.Println(crtDrawing[120:160])
	fmt.Println(crtDrawing[160:200])
	fmt.Println(crtDrawing[200:240])
	return fmt.Sprint(0)
}

func spriteVisible(index int, xValues []int) bool {
	modI := index % 40
	xVal := xValues[index]
	if xVal == modI || xVal == modI-1 || xVal == modI+1 {
		return true
	}
	return false
}
