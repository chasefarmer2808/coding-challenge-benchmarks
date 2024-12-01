package day01

import (
	"coding-challenge-runner/pkg/input"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func Part1(f *os.File) int64 {
	var diff int64 = 0
	var left []int
	var right []int

	for l := range input.Lines(f) {
		tokens := strings.Fields(l)
		n1, _ := strconv.Atoi(tokens[0])
		n2, _ := strconv.Atoi(tokens[1])
		left = append(left, n1)
		right = append(right, n2)
	}

	sort.Ints(left)
	sort.Ints(right)

	for i, n := range left {
		diff += int64(math.Abs(float64(n - right[i])))
	}

	return diff
}

func Part2(f *os.File) int64 {
	score := 0
	countMap := make(map[int]int)
	var left []int

	for l := range input.Lines(f) {
		tokens := strings.Fields(l)
		n1, _ := strconv.Atoi(tokens[0])
		n2, _ := strconv.Atoi(tokens[1])
		left = append(left, n1)

		count, ok := countMap[n2]
		if !ok {
			countMap[n2] = 1
		} else {
			countMap[n2] = count + 1
		}
	}

	for _, n := range left {
		count, ok := countMap[n]
		if ok {
			score += n * count
		}
	}

	return int64(score)
}
