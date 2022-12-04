package day4

import (
	"testing"
)

func TestDay4Task1(t *testing.T) {
	day := new(Day4)
	var tests = []struct {
		fileName string
		want     int
	}{
		{"../../inputs/day4/day4_test1.txt", 2},
		{"../../inputs/day4/day4_test2.txt", 3},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day.Init(tt.fileName)
			ans := day.Task1()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestDay4Task2(t *testing.T) {
	day := new(Day4)
	var tests = []struct {
		fileName string
		want     int
	}{
		{"../../inputs/day4/day4_test1.txt", 4},
		{"../../inputs/day4/day4_test2.txt", 6},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day.Init(tt.fileName)
			ans := day.Task2()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
