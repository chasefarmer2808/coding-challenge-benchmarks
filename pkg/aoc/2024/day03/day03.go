package day03

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func Part1(f *os.File) int64 {
	var sum int64 = 0
	mulCmd := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	buf := make([]byte, 100000)
	_, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	cmdMatches := mulCmd.FindAllString(string(buf), -1)

	for _, m := range cmdMatches {
		fmt.Println(m)
		n1, n2, err := extractOperands(m)
		if err != nil {
			panic(err)
		}
		sum += int64(n1) * int64(n2)
	}

	return sum
}

func Part2(f *os.File) int64 {
	var sum int64 = 0
	mulCmd := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(don't\(\))|(do\(\))`)
	enabled := true
	buf := make([]byte, 100000)
	_, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	cmdMatches := mulCmd.FindAllString(string(buf), -1)

	for _, m := range cmdMatches {
		fmt.Println(m)

		switch m {
		case "do()":
			enabled = true
		case "don't()":
			enabled = false
		default:
			if enabled {
				n1, n2, err := extractOperands(m)
				if err != nil {
					panic(err)
				}
				sum += int64(n1) * int64(n2)
			}
		}
	}

	return sum
}

func extractOperands(s string) (int, int, error) {
	digits := regexp.MustCompile(`\d{1,3}`)
	nums := digits.FindAllString(s, -1)

	if len(nums) != 2 {
		return 0, 0, errors.New("failed to extract operands: " + s)
	}

	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	return n1, n2, nil
}
