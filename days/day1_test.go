package days

import (
	"testing"
)

func TestTask1(t *testing.T) {
	day1 := new(Day1)
	var tests = []struct {
		fileName string
		want     int
	}{
		{"../inputs/day1_task1_test1.txt", 24000},
		{"../inputs/day1_task1_test2.txt", 500000},
		{"../inputs/day1_task1_test3.txt", 500000},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			ans := day1.Task1(tt.fileName)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestTask2(t *testing.T) {
	day1 := new(Day1)
	var tests = []struct {
		fileName string
		want     int
	}{
		{"../inputs/day1_task2_test1.txt", 45000},
		{"../inputs/day1_task2_test2.txt", 577000},
		{"../inputs/day1_task2_test3.txt", 577000},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			ans := day1.Task2(tt.fileName)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
