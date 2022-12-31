package main

import (
	"advent_of_code_2022/days/day1"
	"advent_of_code_2022/days/day2"
	"advent_of_code_2022/days/day3"
	"advent_of_code_2022/days/day4"
	"advent_of_code_2022/days/day5"
	"advent_of_code_2022/days/day6"
	"advent_of_code_2022/days/day7"
	"advent_of_code_2022/days/day8"
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
	case 3:
		day = new(day3.Day3)
		day.Init("inputs/day3/day3.txt")
	case 4:
		day = new(day4.Day4)
		day.Init("inputs/day4/day4.txt")
	case 5:
		day = new(day5.Day5)
		day.Init("inputs/day5/day5.txt")
	case 6:
		day = new(day6.Day)
		day.Init("inputs/day6/day6.txt")
	case 7:
		day = new(day7.Day)
		day.Init("inputs/day7/day7.txt")
	case 8:
		day = new(day8.Day)
		day.Init("inputs/day8/day8.txt")
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
	Task1() string
	Task2() string
}
