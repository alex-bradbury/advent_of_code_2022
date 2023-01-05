package day11

import (
	"advent_of_code_2022/utils"
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Day struct {
	monkeys []monkey
}

type monkey struct {
	items       []int
	operation   func(int) int
	test        int
	trueMonkey  int
	falseMonkey int
}

func (day *Day) Init(filePath string) {
	println("running Init")
	println("---------------")

	currentMonkey := -1
	lines := utils.ReadTxtFile(filePath)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		parts := strings.Split(line, " ")
		if len(parts) > 0 {
			switch parts[0] {
			case "Monkey":
				currentMonkey++
				day.monkeys = append(day.monkeys, monkey{})
			case "Starting":
				for i, part := range parts {
					if i < 2 {
						continue
					}
					if part[len(part)-1:] == "," {
						part = part[:len(part)-1]
					}
					num, _ := strconv.Atoi(part)
					day.monkeys[currentMonkey].items = append(day.monkeys[currentMonkey].items, num)
				}
			case "Operation:":
				var op func(int) int
				opNum := 0
				if match, _ := regexp.Match("[0-9]+", []byte(parts[5])); match {
					opNum, _ = strconv.Atoi(parts[5])
				}
				switch parts[4] {
				case "*":
					op = func(i int) int {
						if opNum != 0 {
							return i * opNum
						} else if parts[5] == "old" {
							return i * i
						}
						return 0
					}
				case "+":
					op = func(i int) int {
						if opNum != 0 {
							return i + opNum
						} else if parts[5] == "old" {
							return i + i
						}
						return 0
					}
				}
				day.monkeys[currentMonkey].operation = op
			case "Test:":
				num, _ := strconv.Atoi(parts[3])
				day.monkeys[currentMonkey].test = num
			case "If":
				switch parts[1] {
				case "true:":
					num, _ := strconv.Atoi(parts[5])
					day.monkeys[currentMonkey].trueMonkey = num
				case "false:":
					num, _ := strconv.Atoi(parts[5])
					day.monkeys[currentMonkey].falseMonkey = num
				}

			}
		}
	}
}

func (day *Day) Task1() string {
	println("running task 1!")
	println("---------------")

	const ROUNDS int = 20
	const WORRY_DIVIDER int = 3
	monkeys := make([]monkey, len(day.monkeys))
	copy(monkeys, day.monkeys)

	return fmt.Sprint(calculateInteractions(ROUNDS, WORRY_DIVIDER, monkeys))
}

func (day *Day) Task2() string {
	println("running task 2!")
	println("---------------")

	const ROUNDS int = 10000
	const WORRY_DIVIDER int = 1

	monkeys := make([]monkey, len(day.monkeys))
	copy(monkeys, day.monkeys)

	return fmt.Sprint(calculateInteractions(ROUNDS, WORRY_DIVIDER, monkeys))
}

func calculateInteractions(rounds, worryDivider int, monkeys []monkey) int {
	monkeyInspections := make([]int, len(monkeys))
	reduceBy := 1
	for _, monkey := range monkeys {
		reduceBy *= monkey.test
	}
	println(reduceBy)
	for i := 0; i < rounds; i++ {
		for j, monkey := range monkeys {
			for _, item := range monkey.items {
				monkeyInspections[j]++
				item = monkey.operation(item)
				item = item / worryDivider
				item = item % reduceBy
				if item%monkey.test == 0 {
					monkeys[monkey.trueMonkey].items = append(monkeys[monkey.trueMonkey].items, item)
				} else {
					monkeys[monkey.falseMonkey].items = append(monkeys[monkey.falseMonkey].items, item)
				}
			}
			monkeys[j].items = make([]int, 0)
		}
	}

	sort.Slice(monkeyInspections, func(i, j int) bool {
		return monkeyInspections[i] > monkeyInspections[j]
	})

	return monkeyInspections[0] * monkeyInspections[1]
}
