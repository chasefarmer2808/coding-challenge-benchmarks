package aoc

import (
	"coding-challenge-runner/pkg/aoc/2023/day01"
	"coding-challenge-runner/pkg/aoc/2023/day02"
	"coding-challenge-runner/pkg/aoc/2023/day03"
	"io"
	"os"
	"testing"
)

type Day struct {
	inputFile     string
	testInputFile string
	parts         []Part
}

type Part struct {
	partFunc        PartFunc
	expectedVal     int
	expectedTestVal int
}

var days = []Day{
	{
		inputFile:     "./2023/day01/input.txt",
		testInputFile: "./2023/day01/input_test.txt",
		parts: []Part{
			{
				partFunc:        day01.Part1,
				expectedVal:     55816,
				expectedTestVal: 0,
			},
			{
				partFunc:        day01.Part2,
				expectedVal:     54980,
				expectedTestVal: 0,
			},
		},
	},
	{
		inputFile:     "./2023/day02/input.txt",
		testInputFile: "./2023/day02/input_test.txt",
		parts: []Part{
			{
				partFunc:        day02.Part1,
				expectedVal:     2207,
				expectedTestVal: 0,
			},
			{
				partFunc:        day02.Part2,
				expectedVal:     62241,
				expectedTestVal: 0,
			},
		},
	},
	{
		inputFile:     "./2023/day03/input.txt",
		testInputFile: "./2023/day03/input_test.txt",
		parts: []Part{
			{
				partFunc:        day03.Part1,
				expectedVal:     539590,
				expectedTestVal: 0,
			},
			{
				partFunc:        day03.Part2,
				expectedVal:     80703636,
				expectedTestVal: 0,
			},
		},
	},
}

func Test2023(t *testing.T) {
	for i, d := range days {
		runDay(d, i+1, t)
	}
}

func Test2023Day01(t *testing.T) {
	runDay(days[0], 1, t)
}

func Test2023Day02(t *testing.T) {
	runDay(days[1], 2, t)
}

func Test2023Day03(t *testing.T) {
	runDay(days[2], 3, t)
}

func runDay(d Day, i int, t *testing.T) {
	f, err := os.Open(d.inputFile)
	if err != nil {
		t.Errorf("cannot open input file %s; %s", d.inputFile, err.Error())
	}
	defer f.Close()

	for j, p := range d.parts {
		soln := p.partFunc(f)
		if soln != p.expectedVal {
			t.Errorf("failed for day %d part %d; expected %d but got %d", i, j+1, p.expectedVal, soln)
		}

		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			t.Errorf("failed to rewind file; %s", err.Error())
		}
	}
}
