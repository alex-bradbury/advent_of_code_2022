package day11

import (
	"fmt"
)

type Day struct {
}

type monkey struct {
	items []int
}

func (day *Day) Init(filePath string) {
	println("running Init")
	println("---------------")

	//lines := utils.ReadTxtFile(filePath)
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	return fmt.Sprint(0)
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	return fmt.Sprint(0)
}
