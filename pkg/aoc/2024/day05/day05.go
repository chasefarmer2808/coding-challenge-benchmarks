package day05

import (
	"coding-challenge-runner/pkg/input"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1(f *os.File) int64 {
	var midSum int64
	var rules []string
	var updates [][]int

	for l := range input.Lines(f) {
		if strings.Contains(l, "|") {
			rules = append(rules, l)
			continue
		}

		if strings.Contains(l, ",") {
			var currUpdates []int
			tokens := strings.Split(l, ",")
			for _, t := range tokens {
				n, _ := strconv.Atoi(t)
				currUpdates = append(currUpdates, n)
			}
			updates = append(updates, currUpdates)
		}
	}

	ruleMap := buildRuleMap(rules)
	fmt.Printf("rule map %+v\n", ruleMap)
	for _, u := range updates {
		fmt.Printf("checking %+v\n", u)
		if slices.IsSortedFunc(u, func(a, b int) int {
			return ruleCompare(a, b, ruleMap)
		}) {
			fmt.Println("ordered!")
			midSum += int64(getMid(u))
		}
	}

	return midSum
}

func Part2(f *os.File) int64 {
	var midSum int64
	var rules []string
	var updates [][]int

	for l := range input.Lines(f) {
		if strings.Contains(l, "|") {
			rules = append(rules, l)
			continue
		}

		if strings.Contains(l, ",") {
			var currUpdates []int
			tokens := strings.Split(l, ",")
			for _, t := range tokens {
				n, _ := strconv.Atoi(t)
				currUpdates = append(currUpdates, n)
			}
			updates = append(updates, currUpdates)
		}
	}

	ruleMap := buildRuleMap(rules)

	for _, u := range updates {
		if !slices.IsSortedFunc(u, func(a, b int) int {
			return ruleCompare(a, b, ruleMap)
		}) {
			slices.SortFunc(u, func(a, b int) int {
				return ruleCompare(a, b, ruleMap)
			})

			midSum += int64(getMid(u))
		}
	}

	return midSum
}

func buildRuleMap(rules []string) map[int][]int {
	rm := make(map[int][]int)

	for _, r := range rules {
		tokens := strings.Split(r, "|")
		n1, _ := strconv.Atoi(tokens[0])
		n2, _ := strconv.Atoi(tokens[1])

		_, ok := rm[n1]
		if !ok {
			rm[n1] = []int{n2}
		} else {
			rm[n1] = append(rm[n1], n2)
		}

		_, ok = rm[n2]
		if !ok {
			rm[n2] = []int{}
		}
	}

	return rm
}

func ruleCompare(a, b int, ruleMap map[int][]int) int {
	return func(a, b int) int {
		if a == b {
			return 0
		}
		if slices.Contains(ruleMap[a], b) {
			return -1
		}
		return 1
	}(a, b)
}

func getMid(nums []int) int {
	idx := len(nums) / 2
	return nums[idx]
}
