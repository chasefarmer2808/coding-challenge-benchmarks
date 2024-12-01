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
	expectedVal     int64
	expectedTestVal int64
}

var days2023 = []Day{
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
	for i, d := range days2023 {
		RunDay(d, i+1, false, t)
	}
}

func Test2023Day01(t *testing.T) {
	RunDay(days2023[0], 1, false, t)
}

func Test2023Day02(t *testing.T) {
	RunDay(days2023[1], 2, false, t)
}

func Test2023Day03(t *testing.T) {
	RunDay(days2023[2], 3, false, t)
}

func RunDay(d Day, i int, useTest bool, t *testing.T) {
	inputFile := d.inputFile
	if useTest {
		inputFile = d.testInputFile
	}

	f, err := os.Open(inputFile)
	if err != nil {
		t.Errorf("cannot open input file %s; %s", inputFile, err.Error())
	}
	defer f.Close()

	for j, p := range d.parts {
		expected := p.expectedVal
		if useTest {
			expected = p.expectedTestVal
		}

		soln := p.partFunc(f)
		if soln != expected {
			t.Errorf("failed for day %d part %d; expected %d but got %d", i, j+1, expected, soln)
		}

		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			t.Errorf("failed to rewind file; %s", err.Error())
		}
	}
}
