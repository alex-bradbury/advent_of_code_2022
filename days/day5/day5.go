package day5

import (
	"advent_of_code_2022/utils"
	"fmt"
	"regexp"
	"strconv"
	"unicode"
)

type Day5 struct {
	crates       []Stack
	instructions []string
}

func (day *Day5) Init(filePath string) {
	lines := utils.ReadTxtFile(filePath)
	var firstPart []string
	var secondPart []string

	for i, line := range lines {
		if len(line) == 0 {
			firstPart = lines[:i]
			secondPart = lines[i+1:]
		}
	}

	var crates []Stack
	for i := len(firstPart) - 1; i >= 0; i-- {
		if i == len(firstPart)-1 {
			for j := 0; j < (len(firstPart[i])+1)/4; j++ {
				crates = append(crates, Stack{})
			}
		} else {
			count := 0
			for j := 0; j < len(firstPart[i])+1; j += 4 {
				if unicode.IsLetter(rune(firstPart[i][j+1])) {
					crates[count].Push(string(firstPart[i][j+1]))
				}
				count++
			}
		}
	}

	day.crates = make([]Stack, len(crates))
	copy(day.crates, crates)
	day.instructions = secondPart
}

func (day *Day5) Task1() string {
	println("running task 1!")
	println("---------------")

	fmt.Println(day.crates)
	workingCrates := make([]Stack, len(day.crates))
	copy(workingCrates, day.crates)
	startingCrates := make([]Stack, len(day.crates))
	copy(startingCrates, day.crates)
	//fmt.Println(startingCrates)

	for _, instruction := range day.instructions {
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(instruction, -1)
		var actions = [3]int{}
		for i, number := range numbers {
			j, err := strconv.Atoi(number)
			if err != nil {
				panic(err)
			}
			actions[i] = j
		}

		for i := 0; i < actions[0]; i++ {
			popped, _ := workingCrates[actions[1]-1].Pop()
			workingCrates[actions[2]-1].Push(popped)
		}
	}

	result := ""
	for _, crate := range workingCrates {
		result = result + crate[len(crate)-1]
	}
	// day.crates = startingCrates
	// fmt.Println(startingCrates)
	fmt.Println(day.crates)

	return result
}

func (day *Day5) Task2() string {
	println("running task 2!")
	println("---------------")

	fmt.Println(day.crates)
	workingCrates := make([]Stack, len(day.crates))
	copy(workingCrates, day.crates)

	for _, instruction := range day.instructions {
		re := regexp.MustCompile("[0-9]+")
		actoinsStrings := re.FindAllString(instruction, -1)
		var actions = []int{}
		for _, i := range actoinsStrings {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			actions = append(actions, j)
		}

		amount := actions[0]
		from := actions[1] - 1
		to := actions[2] - 1
		index := len(workingCrates[from]) - amount
		cratesToMove := workingCrates[from][index:]
		workingCrates[from] = workingCrates[from][:index]
		workingCrates[to] = append(workingCrates[to], cratesToMove...)
	}

	result := ""
	for _, crate := range workingCrates {
		result = result + crate[len(crate)-1]
	}

	return result
}

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}
