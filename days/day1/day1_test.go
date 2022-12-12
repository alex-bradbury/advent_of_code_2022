package day1

import (
	"testing"
)

func TestDay1Task1(t *testing.T) {
	day1 := new(Day1)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day1/day1_test1.txt", "24000"},
		{"../../inputs/day1/day1_test2.txt", "500000"},
		{"../../inputs/day1/day1_test3.txt", "500000"},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day1.Init(tt.fileName)
			ans := day1.Task1()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}

func TestDay1Task2(t *testing.T) {
	day1 := new(Day1)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day1/day1_test1.txt", "45000"},
		{"../../inputs/day1/day1_test2.txt", "577000"},
		{"../../inputs/day1/day1_test3.txt", "577000"},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day1.Init(tt.fileName)
			ans := day1.Task2()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
