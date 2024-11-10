package aoc

import "os"

type Runnable interface {
	Part1(f *os.File) int
	Part2(f *os.File) int
}
