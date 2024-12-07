package aoc

import (
	"coding-challenge-runner/pkg/aoc/2024/day01"
	"coding-challenge-runner/pkg/aoc/2024/day02"
	"coding-challenge-runner/pkg/aoc/2024/day03"
	"coding-challenge-runner/pkg/aoc/2024/day04"
	"coding-challenge-runner/pkg/aoc/2024/day05"
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
				expectedVal:     488,
				expectedTestVal: 4,
			},
		},
	},
	{
		inputFile:     "./2024/day03/input.txt",
		testInputFile: "./2024/day03/test_input.txt",
		parts: []Part{
			{
				partFunc:        day03.Part1,
				expectedVal:     166630675,
				expectedTestVal: 161,
			},
			{
				partFunc:        day03.Part2,
				expectedVal:     93465710,
				expectedTestVal: 48,
			},
		},
	},
	{
		inputFile:     "./2024/day04/input.txt",
		testInputFile: "./2024/day04/test_input.txt",
		parts: []Part{
			{
				partFunc:        day04.Part1,
				expectedVal:     2718,
				expectedTestVal: 18,
			},
			{
				partFunc:        day04.Part2,
				expectedVal:     2046,
				expectedTestVal: 9,
			},
		},
	},
	{
		inputFile:     "./2024/day05/input.txt",
		testInputFile: "./2024/day05/test_input.txt",
		parts: []Part{
			{
				partFunc:        day05.Part1,
				expectedVal:     0,
				expectedTestVal: 0,
			},
			{
				partFunc:        day05.Part2,
				expectedVal:     0,
				expectedTestVal: 0,
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

func Test2024Day03(t *testing.T) {
	RunDay(days2024[2], 3, false, t)
}

func Test2024Day04(t *testing.T) {
	RunDay(days2024[3], 4, false, t)
}

func Test2024Day05(t *testing.T) {
	RunDay(days2024[4], 5, false, t)
}
