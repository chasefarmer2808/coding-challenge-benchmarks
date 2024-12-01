package aoc

import (
	"coding-challenge-runner/pkg/aoc/2024/day01"
	"testing"
)

var days2024 = []Day{
	{
		inputFile:     "./2024/day01/input.txt",
		testInputFile: "./2024/day01/test_input.txt",
		parts: []Part{
			{
				partFunc:        day01.Part1,
				expectedVal:     2430334,
				expectedTestVal: 11,
			},
			{
				partFunc:        day01.Part2,
				expectedVal:     28786472,
				expectedTestVal: 31,
			},
		},
	},
}

func Test2024(t *testing.T) {
	for i, d := range days2024 {
		RunDay(d, i+1, false, t)
	}
}

func Test2024Day01(t *testing.T) {
	RunDay(days2024[0], 1, false, t)
}
