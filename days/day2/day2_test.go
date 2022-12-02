package day2

import (
	"testing"
)

func TestDay2Task1(t *testing.T) {
	day1 := new(Day2)
	var tests = []struct {
		fileName string
		want     int
	}{
		{"../../inputs/day2/day2_test1.txt", 15},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day1.Init(tt.fileName)
			ans := day1.Task1()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

func TestDay2Task2(t *testing.T) {
	day1 := new(Day2)
	var tests = []struct {
		fileName string
		want     int
	}{
		{"../../inputs/day2/day2_test1.txt", 12},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day1.Init(tt.fileName)
			ans := day1.Task2()
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
