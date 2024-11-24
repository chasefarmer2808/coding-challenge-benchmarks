package aoc2023

import (
	aoc2023day01 "coding-challenge-runner/pkg/aoc/2023/day01"
	"os"
)

type Day01 struct {
}

func (d1 *Day01) Run() error {
	input, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer input.Close()

	aoc2023day01.Part1(input)
	aoc2023day01.Part2(input)

	return nil
}

func (d1 *Day01) Name() string {
	return "AOC 2023 Day 01"
}
