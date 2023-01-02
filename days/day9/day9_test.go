package day9

import (
	"testing"
)

func TestDay9Task1(t *testing.T) {
	day := new(Day)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day9/day9_test1.txt", "13"},
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

func TestDay9Task2(t *testing.T) {
	day := new(Day)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day9/day9_test1.txt", "1"},
		{"../../inputs/day9/day9_test2.txt", "36"},
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
