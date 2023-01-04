package day11

import (
	"testing"
)

func TestDay11Task1(t *testing.T) {
	day := new(Day)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day11/day11_test1.txt", "10605"},
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

func TestDay10Task2(t *testing.T) {
	day := new(Day)
	var tests = []struct {
		fileName string
		want     string
	}{
		{"../../inputs/day11/day11_test1.txt", "0"},
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
