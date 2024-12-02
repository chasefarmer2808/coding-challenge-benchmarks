package day02

import (
	"coding-challenge-runner/pkg/input"
	"os"
	"strconv"
	"strings"
)

func Part1(f *os.File) int64 {
	var safeCount int64 = 0

outer:
	for l := range input.Lines(f) {
		tokens := strings.Split(l, " ")
		prev := 0
		dir := false // true = asc, false = desc

		for i, s := range tokens {
			n, _ := strconv.Atoi(s)

			if i == 0 {
				prev = n
				continue
			}

			if i == 1 {
				dir = n > prev
			}

			diff := n - prev
			if diff == 0 || diff > 3 || diff < -3 {
				continue outer
			}

			if n > prev && !dir {
				continue outer
			}

			if n < prev && dir {
				continue outer
			}

			prev = n
		}

		safeCount++
	}

	return safeCount
}

func Part2(f *os.File) int64 {
	return 0
}
