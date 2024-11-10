package day01

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Aoc2023Day01 struct {
}

func (d1 *Aoc2023Day01) Part1(f *os.File) int {
	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		curr := ""
		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= 48 && c <= 57 {
				curr += string(c)
				break
			}
		}

		for i := len(line) - 1; i > -1; i-- {
			c := line[i]
			if c >= 48 && c <= 57 {
				curr += string(c)
				break
			}
		}

		i, _ := strconv.Atoi(curr)
		sum += i
	}

	return sum
}

/*
one, two, three, four, five, six, seven, eight, nine
one -> 1
two -> 2
three -> 3
...

two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen

for each char
	scan line and append to string
	if current one exists in map, add it to the decoded string
	if current one is a digit, add it to the decoded string

take first and last of decoded string
*/

func (d1 *Aoc2023Day01) Part2(f *os.File) int {
	numMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		decoded := ""

		for i := 0; i < len(line); i++ {
			c := line[i]
			if c >= 48 && c <= 57 {
				decoded += string(c)
				continue
			}

			for k, v := range numMap {
				if strings.HasPrefix(line[i:], k) {
					decoded += v
				}
			}
		}
		numStr := fmt.Sprintf("%s%s", string(decoded[0]), string(decoded[len(decoded)-1]))
		num, _ := strconv.Atoi(string(numStr))
		sum += num
	}

	return sum
}

func (d1 *Aoc2023Day01) Run() error {
	input, err := os.Open("input.txt")
	if err != nil {
		return err
	}
	defer input.Close()

	d1.Part1(input)
	d1.Part2(input)

	return nil
}

func (d1 *Aoc2023Day01) Name() string {
	return "AOC 2023 Day 01"
}
