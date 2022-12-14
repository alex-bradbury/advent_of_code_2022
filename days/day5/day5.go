package day5

import (
	"advent_of_code_2022/utils"
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

	day.crates = crates
	day.instructions = secondPart
}

func (day *Day5) Task1() string {
	println("running task 1!")
	println("---------------")

	for _, instruction := range day.instructions {
		re := regexp.MustCompile("[0-9]+")
		numbers := re.FindAllString(instruction, -1)
		var t2 = []int{}
		for _, i := range numbers {
			j, err := strconv.Atoi(i)
			if err != nil {
				panic(err)
			}
			t2 = append(t2, j)
		}

		for i := 0; i < t2[0]; i++ {
			popped, _ := day.crates[t2[1]-1].Pop()
			day.crates[t2[2]-1].Push(popped)
		}
	}

	result := ""
	for _, crate := range day.crates {
		result = result + crate[len(crate)-1]
	}

	return result
}

func (day *Day5) Task2() string {
	println("running task 2!")
	println("---------------")

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
		cratesToMove := []string{}
		for i := 0; i < amount; i++ {
			cratesToMove = append(cratesToMove, "")
		}
		for i := 0; i < amount; i++ {
			popped, _ := day.crates[actions[1]-1].Pop()
			cratesToMove[amount-1] = popped
			amount--
		}
		for _, crate := range cratesToMove {
			day.crates[actions[2]-1].Push(crate)
		}
	}

	result := ""
	for _, crate := range day.crates {
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
