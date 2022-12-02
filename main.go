package main

import (
	"advent_of_code_2022/days/day1"
	"advent_of_code_2022/days/day2"
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
	switch dayToRun {
	case 1:
		day = new(day1.Day1)
		day.Init("inputs/day1/day1.txt")
	case 2:
		day = new(day2.Day2)
		day.Init("inputs/day2/day2.txt")
	default:
		println("Day has not been written yet")
		return
	}

	task1Result := day.Task1()
	task2Result := day.Task2()

	fmt.Println("task1: ", task1Result)
	fmt.Println("task2: ", task2Result)
}

type day interface {
	Init(filename string)
	Task1() int
	Task2() int
}
