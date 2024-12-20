package aoc

import (
	"coding-challenge-runner/pkg/aoc/2024/day01"
	"coding-challenge-runner/pkg/aoc/2024/day02"
	"coding-challenge-runner/pkg/aoc/2024/day03"
	"coding-challenge-runner/pkg/aoc/2024/day04"
	"coding-challenge-runner/pkg/aoc/2024/day05"
	"coding-challenge-runner/pkg/aoc/2024/day06"
	"coding-challenge-runner/pkg/aoc/2024/day07"
	"coding-challenge-runner/pkg/aoc/2024/day08"
	"coding-challenge-runner/pkg/aoc/2024/day09"
	"coding-challenge-runner/pkg/aoc/2024/day10"
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
				expectedVal:     5091,
				expectedTestVal: 143,
			},
			{
				partFunc:        day05.Part2,
				expectedVal:     4681,
				expectedTestVal: 123,
			},
		},
	},
	{
		inputFile:     "./2024/day06/input.txt",
		testInputFile: "./2024/day06/test_input.txt",
		parts: []Part{
			{
				partFunc:        day06.Part1,
				expectedVal:     4711,
				expectedTestVal: 41,
			},
			{
				partFunc:        day06.Part2,
				expectedVal:     1562,
				expectedTestVal: 6,
			},
		},
	},
	{
		inputFile:     "./2024/day07/input.txt",
		testInputFile: "./2024/day07/test_input.txt",
		parts: []Part{
			{
				partFunc:        day07.Part1,
				expectedVal:     7710205485870,
				expectedTestVal: 3749,
			},
			{
				partFunc:        day07.Part2,
				expectedVal:     20928985450275,
				expectedTestVal: 11387,
			},
		},
	},
	{
		inputFile:     "./2024/day08/input.txt",
		testInputFile: "./2024/day08/test_input.txt",
		parts: []Part{
			{
				partFunc:        day08.Part1,
				expectedVal:     426,
				expectedTestVal: 14,
			},
			{
				partFunc:        day08.Part2,
				expectedVal:     0,
				expectedTestVal: 34,
			},
		},
	},
	{
		inputFile:     "./2024/day09/input.txt",
		testInputFile: "./2024/day09/test_input.txt",
		parts: []Part{
			{
				partFunc:        day09.Part1,
				expectedVal:     6382875730645,
				expectedTestVal: 1928,
			},
			{
				partFunc:        day09.Part2,
				expectedVal:     6420913943576,
				expectedTestVal: 2858,
			},
		},
	},
	{
		inputFile:     "./2024/day10/input.txt",
		testInputFile: "./2024/day10/test_input.txt",
		parts: []Part{
			{
				partFunc:        day10.Part1,
				expectedVal:     778,
				expectedTestVal: 36,
			},
			{
				partFunc:        day10.Part2,
				expectedVal:     1925,
				expectedTestVal: 81,
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

func Test2024Day06(t *testing.T) {
	RunDay(days2024[5], 6, false, t)
}

func Test2024Day07(t *testing.T) {
	RunDay(days2024[6], 7, false, t)
}

func Test2024Day08(t *testing.T) {
	RunDay(days2024[7], 8, false, t)
}

func Test2024Day09(t *testing.T) {
	RunDay(days2024[8], 9, false, t)
}

func Test2024Day10(t *testing.T) {
	RunDay(days2024[9], 10, false, t)
}
