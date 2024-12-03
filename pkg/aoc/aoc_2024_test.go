package aoc

import (
	"coding-challenge-runner/pkg/aoc/2024/day01"
	"coding-challenge-runner/pkg/aoc/2024/day02"
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
	{
		inputFile:     "./2024/day02/input.txt",
		testInputFile: "./2024/day02/test_input.txt",
		parts: []Part{
			{
				partFunc:        day02.Part1,
				expectedVal:     432,
				expectedTestVal: 2,
			},
			{
				partFunc:        day02.Part2,
				expectedVal:     0,
				expectedTestVal: 4,
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

func Test2024Day02(t *testing.T) {
	RunDay(days2024[1], 2, false, t)
}
