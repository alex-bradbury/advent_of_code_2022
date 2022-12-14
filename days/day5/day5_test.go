package day5

import (
	"testing"
)

func TestDay5Task1(t *testing.T) {
	day := new(Day5)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day5/day5_test2.txt", "CMZ"},
		{"../../inputs/day5/day5_test1.txt", "GPZ"},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day.Init(tt.fileName)
			ans := day.Task1()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestDay5Task2(t *testing.T) {
	day := new(Day5)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day5/day5_test1.txt", "PGL"},
		{"../../inputs/day5/day5_test2.txt", "MCD"},
		{"../../inputs/day5/day5_test3.txt", "PCQCVPPVZ"},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day.Init(tt.fileName)
			ans := day.Task2()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
