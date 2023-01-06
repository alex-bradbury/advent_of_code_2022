package day12

import (
	"advent_of_code_2022/utils"
	"fmt"
	"sort"
)

type Day struct {
	landscape   [][]rune
	start       location
	destination location
	aLocations  []location
}

type location struct {
	x, y int
}

type travelWeight struct {
	location   location
	distanceTo int
}

func (day *Day) Init(filePath string) {
	println("running Init")
	println("---------------")

	lines := utils.ReadTxtFile(filePath)
	day.landscape = make([][]rune, len(lines))

	for y, line := range lines {
		day.landscape[y] = make([]rune, len(line))
		for x, letter := range line {
			if string(letter) == "S" {
				day.start.x = x
				day.start.y = y
				letter = 'a'
			}
			if string(letter) == "E" {
				day.destination.x = x
				day.destination.y = y
				letter = 'z'
			}
			day.landscape[y][x] = letter
			if letter == 'a' {
				day.aLocations = append(day.aLocations, location{x, y})
			}
		}
	}
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	distance, possible := getShortestPathToDesitnation(day.landscape, day.start, day.destination)
	if possible {
		return fmt.Sprint(distance)
	} else {
		return "Not possible to reach destination from starting location"
	}
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	distances := make([]int, 0)
	for _, location := range day.aLocations {
		distance, possible := getShortestPathToDesitnation(day.landscape, location, day.destination)
		if possible {
			distances = append(distances, distance)
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i] < distances[j]
	})

	return fmt.Sprint(distances[0])
}

func saveNewLocation(newLoc location, newWeight int, list *[]travelWeight, weightMap map[location]int) {
	*list = append(*list, travelWeight{newLoc, newWeight})
	weightMap[newLoc] = newWeight
}

func getShortestPathToDesitnation(landscape [][]rune, start, dest location) (int, bool) {
	var travelWeightsList []travelWeight
	visitedNodes := make(map[location]bool)
	travelWeightsMap := make(map[location]int)
	travelWeightsMap[start] = 0
	currentNode := start

	for currentNode != dest {
		currentWeight := travelWeightsMap[currentNode]
		currentRune := landscape[currentNode.y][currentNode.x]

		// check left
		if currentNode.x != 0 &&
			landscape[currentNode.y][currentNode.x-1] <= currentRune+1 &&
			!visitedNodes[location{currentNode.x - 1, currentNode.y}] {
			saveNewLocation(location{currentNode.x - 1, currentNode.y}, currentWeight+1, &travelWeightsList, travelWeightsMap)
		}

		// check right
		if currentNode.x != len(landscape[0])-1 &&
			landscape[currentNode.y][currentNode.x+1] <= currentRune+1 &&
			!visitedNodes[location{currentNode.x + 1, currentNode.y}] {
			saveNewLocation(location{currentNode.x + 1, currentNode.y}, currentWeight+1, &travelWeightsList, travelWeightsMap)
		}

		// check top
		if currentNode.y != 0 &&
			landscape[currentNode.y-1][currentNode.x] <= currentRune+1 &&
			!visitedNodes[location{currentNode.x, currentNode.y - 1}] {
			saveNewLocation(location{currentNode.x, currentNode.y - 1}, currentWeight+1, &travelWeightsList, travelWeightsMap)

		}

		// check bottom
		if currentNode.y != len(landscape)-1 &&
			landscape[currentNode.y+1][currentNode.x] <= currentRune+1 &&
			!visitedNodes[location{currentNode.x, currentNode.y + 1}] {
			saveNewLocation(location{currentNode.x, currentNode.y + 1}, currentWeight+1, &travelWeightsList, travelWeightsMap)
		}

		visitedNodes[currentNode] = true

		newTravelWeights := make([]travelWeight, 0)
		for _, travelWeight := range travelWeightsList {
			if !visitedNodes[travelWeight.location] {
				newTravelWeights = append(newTravelWeights, travelWeight)
			}
		}
		travelWeightsList = make([]travelWeight, len(newTravelWeights))
		copy(travelWeightsList, newTravelWeights)
		if len(travelWeightsList) == 0 {
			return 0, false
		}

		sort.Slice(travelWeightsList, func(i, j int) bool {
			return travelWeightsList[i].distanceTo < travelWeightsList[j].distanceTo
		})

		currentNode = travelWeightsList[0].location
	}

	return travelWeightsMap[dest], true
}
