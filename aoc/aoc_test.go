package aoc

import (
	"coding-challenge-runner/aoc/2023/day01"
	"io"
	"os"
	"testing"
)

type DayCase struct {
	inputFile string
	day       Runnable
	p1Val     int
	p2Val     int
}

func Test2023(t *testing.T) {
	cases := []DayCase{
		{
			inputFile: "./2023/day01/input.txt",
			day:       &day01.Aoc2023Day01{},
			p1Val:     55816,
			p2Val:     54980,
		},
	}

	for _, c := range cases {
		f, err := os.Open(c.inputFile)
		if err != nil {
			t.Errorf("cannot open input file %s; %s", c.inputFile, err.Error())
		}
		defer f.Close()

		p1 := c.day.Part1(f)
		if p1 != c.p1Val {
			t.Errorf("part 1 failed. expected %v but got %v", c.p1Val, p1)
		}

		_, err = f.Seek(0, io.SeekStart)
		if err != nil {
			t.Errorf("failed to rewind file; %s", err.Error())
		}

		p2 := c.day.Part2(f)
		if p2 != c.p2Val {
			t.Errorf("part 2 failed. expected %v but got %v", c.p2Val, p2)
		}
	}
}
