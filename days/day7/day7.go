package day7

import (
	"advent_of_code_2022/utils"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Day struct {
	directories []directory
	totalSize   int
}

type directory struct {
	name  string
	files []file
	size  int
}

type file struct {
	name string
	size string
}

func (day *Day) Init(filePath string) {
	println("running Init")
	println("---------------")
	lines := utils.ReadTxtFile(filePath)
	currentDir := ""
	currentIndex := -1
	dirStack := utils.StringStack{}
	currentDirSize := 0
	for _, line := range lines {
		partsOfLine := strings.Split(line, " ")
		if partsOfLine[0] == "$" {
			if partsOfLine[1] == "cd" {
				if currentIndex > -1 {
					day.directories[currentIndex].size = currentDirSize
				}
				currentIndex = -1
				currentDirSize = 0
				if partsOfLine[2] == ".." {
					dirStack.Pop()
					currentDir = ""
					continue
				}

				currentDir = partsOfLine[2]
				dirStack.Push(partsOfLine[2])
				currentIndex = len(day.directories)
				day.directories = append(day.directories, directory{getFullPath(dirStack), []file{}, 0})
				continue
			} else if partsOfLine[1] == "ls" {
				continue
			}
		}
		if len(currentDir) > 0 {
			fileName := partsOfLine[1]
			fileSize := partsOfLine[0]
			if fileSize != "dir" {
				size, err := strconv.Atoi(partsOfLine[0])
				if err != nil {
					panic(err)
				}
				currentDirSize += size
			} else {
				fileName = getFullPath(dirStack) + partsOfLine[1] + "/"
			}
			day.directories[currentIndex].files = append(day.directories[currentIndex].files, file{fileName, fileSize})
		}
	}
	if currentIndex > -1 {
		day.directories[currentIndex].size = currentDirSize
	}
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	totalSize := 0
	for _, directory := range day.directories {
		dirSize := calcSizeOfDir(directory, day.directories)
		if dirSize <= 100000 {
			totalSize += dirSize
		}
	}

	return fmt.Sprint(totalSize)
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	remainingSpace := 70000000 - calcSizeOfDir(day.directories[0], day.directories)
	lowestSize := 70000000
	spaceNeeded := 30000000 - remainingSpace
	for _, directory := range day.directories {
		dirSize := calcSizeOfDir(directory, day.directories)
		if dirSize >= spaceNeeded {
			if dirSize < lowestSize {
				lowestSize = dirSize
			}
		}
	}

	return fmt.Sprint(lowestSize)
}

func calcSizeOfDir(dirToCalc directory, allDirs []directory) int {
	dirSize := dirToCalc.size
	for _, file := range dirToCalc.files {
		if file.size == "dir" {
			index, err := findIndexOfDirectory(allDirs, file.name)
			if err != nil {
				panic(err)
			}
			dirSize += calcSizeOfDir(allDirs[index], allDirs)
		}
	}
	return dirSize
}

func findIndexOfDirectory(dir []directory, dirName string) (int, error) {
	for i, dir := range dir {
		if dir.name == dirName {
			return i, nil
		}
	}
	return -1, errors.New("can't find directory")
}

func getFullPath(dirStack utils.StringStack) string {
	path := ""
	for _, dir := range dirStack {
		path += dir + "/"
	}
	return path
}
