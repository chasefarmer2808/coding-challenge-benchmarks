package aoc

import (
	"coding-challenge-runner/pkg/aoc/2023/day01"
	"coding-challenge-runner/pkg/aoc/2023/day02"
	"coding-challenge-runner/pkg/aoc/2023/day03"
	"io"
	"os"
	"testing"
)

type DayCase struct {
	inputFile string
	part1     PartFunc
	part2     PartFunc
	p1Val     int
	p2Val     int
}

func Test2023(t *testing.T) {
	cases := []DayCase{
		{
			inputFile: "./2023/day01/input.txt",
			part1:     day01.Part1,
			part2:     day01.Part2,
			p1Val:     55816,
			p2Val:     54980,
		},
		{
			inputFile: "./2023/day02/input.txt",
			part1:     day02.Part1,
			part2:     day02.Part2,
			p1Val:     2207,
			p2Val:     62241,
		},
	}

	for _, c := range cases {
		f, err := os.Open(c.inputFile)
		if err != nil {
			t.Errorf("cannot open input file %s; %s", c.inputFile, err.Error())
		}
		defer f.Close()

		p1 := c.part1(f)
		if p1 != c.p1Val {
			t.Errorf("part 1 failed. expected %v but got %v", c.p1Val, p1)
		}

		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			t.Errorf("failed to rewind file; %s", err.Error())
		}

		p2 := c.part2(f)
		if p2 != c.p2Val {
			t.Errorf("part 2 failed. expected %v but got %v", c.p2Val, p2)
		}
	}
}

func Test2023Day03(t *testing.T) {
	inputFile := "./2023/day03/input.txt"
	expectedP1 := 539590
	expectedP2 := 80703636
	f, err := os.Open(inputFile)
	if err != nil {
		t.Errorf("cannot open input file %s; %s", inputFile, err.Error())
	}
	defer f.Close()

	p1 := day03.Part1(f)
	if p1 != expectedP1 {
		t.Errorf("part 1 failed. expected %v but got %v", expectedP1, p1)
	}

	_, err = f.Seek(0, io.SeekStart)
	if err != nil {
		t.Errorf("failed to rewind file; %s", err.Error())
	}

	p2 := day03.Part2(f)
	if p2 != expectedP2 {
		t.Errorf("part 2 failed. expected %v but got %v", expectedP2, p2)
	}
}
