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
			v, color := parseColorToken(c)

			switch color {
			case "red":
				r += v
			case "green":
				g += v
			case "blue":
				b += v
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
	sum := 0
	var colorGroup = regexp.MustCompile(`\d+\s+(red|green|blue)`)

	for l := range input.Lines(f) {
		minR, minG, minB := 0, 0, 0
		colors := colorGroup.FindAllString(l, -1)

		for _, c := range colors {
			v, cStr := parseColorToken(c)

			switch cStr {
			case "red":
				if v > minR {
					minR = v
				}
			case "green":
				if v > minG {
					minG = v
				}
			case "blue":
				if v > minB {
					minB = v
				}
			}
		}

		power := minR * minG * minB
		sum += power
	}
	return sum
}

func parseColorToken(token string) (int, string) {
	tokens := strings.Split(token, " ")
	value, err := strconv.Atoi(tokens[0])
	if err != nil {
		panic(err)
	}
	color := tokens[1]
	return value, color
}
