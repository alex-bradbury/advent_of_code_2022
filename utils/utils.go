package utils

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCsvFile(day string) [][]string {
	f, err := os.Open("./inputs/day" + day + ".csv")
	if err != nil {
		log.Fatal("Unable to read input file day"+day, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for day"+day, err)
	}

	return records
}

func ReadTxtFile(filePath string) []string {
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

type StringStack []string

// IsEmpty: check if stack is empty
func (s *StringStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *StringStack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *StringStack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

type IntStack []int

// IsEmpty: check if stack is empty
func (s *IntStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *IntStack) Push(num int) {
	*s = append(*s, num) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *IntStack) Pop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// read top element of stack. Return false if stack is empty.
func (s *IntStack) ReadTop() (int, bool) {
	if s.IsEmpty() {
		return 0, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

type ArrayStack [][]any

// IsEmpty: check if stack is empty
func (s *ArrayStack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *ArrayStack) Push(array []any) {
	*s = append(*s, array) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *ArrayStack) Pop() ([]any, bool) {
	if s.IsEmpty() {
		return []any{}, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

// read top element of stack. Return false if stack is empty.
func (s *ArrayStack) ReadTop() (*[]any, bool) {
	if s.IsEmpty() {
		return &[]any{}, false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return &element, true
	}
}
