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
outer:
	for _, u := range updates {
		fmt.Printf("checking %+v\n", u)
		for i, currUpdate := range u {
			if !isOrdered(currUpdate, u[i+1:], ruleMap) {
				continue outer
			}
		}
		fmt.Println("ordered!")
		midSum += int64(getMid(u))
	}

	return midSum
}

func Part2(f *os.File) int64 {
	return 0
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

func isOrdered(n int, updates []int, ruleMap map[int][]int) bool {
	rules := ruleMap[n]
	fmt.Printf("checking %d with updates %+v and rules %+v\n", n, updates, rules)

	if len(updates) == 0 {
		return true
	}

	for _, u := range updates {
		if !slices.Contains(rules, u) {
			return false
		}
	}

	return true
}

func getMid(nums []int) int {
	idx := len(nums) / 2
	return nums[idx]
}
