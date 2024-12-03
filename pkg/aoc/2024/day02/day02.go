package day02

import (
	"coding-challenge-runner/pkg/input"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1(f *os.File) int64 {
	var safeCount int64 = 0

	for l := range input.Lines(f) {
		tokens := strings.Split(l, " ")
		report := make([]int, len(tokens))

		for i, s := range tokens {
			n, _ := strconv.Atoi(s)
			report[i] = n
		}

		bl := getBadLevel(report)
		if bl == nil {
			safeCount++
		}
	}

	return safeCount
}

func Part2(f *os.File) int64 {
	var safeCount int64 = 0
	j := -1

	for l := range input.Lines(f) {
		j++
		fmt.Printf("checking line %d", j)
		fmt.Println()
		tokens := strings.Split(l, " ")
		report := make([]int, len(tokens))

		for i, s := range tokens {
			n, _ := strconv.Atoi(s)
			report[i] = n
		}

		bl := getBadLevel(report)
		if bl == nil {
			safeCount++
			fmt.Println("safe")
			continue
		}

		for i := range report {
			sub := slices.Delete(slices.Clone(report), i, i+1)
			bl := getBadLevel(sub)
			if bl == nil {
				safeCount++
				fmt.Println("safe")
				break
			}
		}
	}

	return safeCount
}

type badLevel struct {
	val int
	idx int
}

func getBadLevel(report []int) *badLevel {
	prev := 0
	prevDiff := 0

	for i, n := range report {
		if i == 0 {
			prev = n
			continue
		}

		if isBadLevel(n, prev, prevDiff) {
			return &badLevel{
				idx: i,
				val: n,
			}
		} else {
			prevDiff = n - prev
			prev = n
		}
	}
	return nil
}

func isBadLevel(lev, prev, prevDiff int) bool {
	diff := lev - prev
	if diff == 0 {
		fmt.Printf("found zero jump %d", lev)
		fmt.Println()
		return true
	}

	if math.Abs(float64(diff)) > 3 {
		fmt.Printf("unsafe jump from %d to %d", prev, lev)
		fmt.Println()
		return true
	}

	if prevDiff > 0 && diff < 0 {
		fmt.Printf("unsafe change in dir from %d to %d", prev, lev)
		fmt.Println()
		return true
	}

	if prevDiff < 0 && diff > 0 {
		fmt.Printf("unsafe change in dir from %d to %d", prev, lev)
		fmt.Println()
		return true
	}

	return false
}
