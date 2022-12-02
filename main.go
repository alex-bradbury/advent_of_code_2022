package main

import (
	"advent_of_code_2022/days"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		println("Please enter a day's tasks to run")
		return
	}

	dayToRun, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println("You must enter just a number for the day")
		return
	}

	var day day
	var filename string
	switch dayToRun {
	case 1:
		day = new(days.Day1)
		filename = "inputs/day1/day1"
	case 2:
		day = new(days.Day2)
		filename = "inputs/day2/day2"
	default:
		println("Day has not been written yet")
		return
	}

	task1Result := day.Task1(filename + ".txt")
	task2Result := day.Task2(filename + ".txt")

	fmt.Println("task1: ", task1Result)
	fmt.Println("task2: ", task2Result)
}

type day interface {
	Task1(filename string) int
	Task2(filename string) int
}
