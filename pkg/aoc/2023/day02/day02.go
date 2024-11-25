package day02

import (
	"coding-challenge-runner/pkg/input"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const MAX_RED = 12
const MAX_GREEN = 13
const MAX_BLUE = 14

func Part1(f *os.File) int {
	sum, i := 0, 0
	var colorGroup = regexp.MustCompile(`\d+\s+(red|green|blue)`)

outer:
	for l := range input.Lines(f) {
		i++
		// Parse tokens
		colors := colorGroup.FindAllString(l, -1)
		for _, c := range colors {
			r, g, b := 0, 0, 0
			tokens := strings.Split(c, " ")
			value, err := strconv.Atoi(tokens[0])
			if err != nil {
				panic(err)
			}
			color := tokens[1]
			switch color {
			case "red":
				r += value
			case "green":
				g += value
			case "blue":
				b += value
			}

			if r > MAX_RED || g > MAX_GREEN || b > MAX_BLUE {
				continue outer
			}
		}

		// Else, got a good game
		sum += i
	}
	return sum
}

func Part2(f *os.File) int {
	return 0
}
