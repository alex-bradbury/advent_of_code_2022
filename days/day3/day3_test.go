package day3

import (
	"testing"
)

func TestDay3Task1(t *testing.T) {
	day := new(Day3)
	var tests = []struct {
		fileName string
		want     int
	}{
		{"../../inputs/day3/day3_test1.txt", 157},
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

func TestDay3Task2(t *testing.T) {
	day := new(Day3)
	var tests = []struct {
		fileName string
		want     int
	}{
		{"../../inputs/day3/day3_test1.txt", 70},
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
