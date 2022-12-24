package day6

import (
	"testing"
)

func TestDay6Task1(t *testing.T) {
	day := new(Day)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day6/day6_test1.txt", "5"},
		{"../../inputs/day6/day6_test2.txt", "6"},
		{"../../inputs/day6/day6_test3.txt", "10"},
		{"../../inputs/day6/day6_test4.txt", "11"},
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

func TestDay6Task2(t *testing.T) {
	day := new(Day)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day6/day6_test1.txt", "23"},
		{"../../inputs/day6/day6_test2.txt", "23"},
		{"../../inputs/day6/day6_test3.txt", "29"},
		{"../../inputs/day6/day6_test4.txt", "26"},
		{"../../inputs/day6/day6_test5.txt", "19"},
	}
	for _, tt := range tests {
		t.Run(tt.fileName, func(t *testing.T) {
			day.Init(tt.fileName)
			day.Task1()
			ans := day.Task2()
			if ans != tt.want {
				t.Errorf("got %s, want %s", ans, tt.want)
			}
		})
	}
}
