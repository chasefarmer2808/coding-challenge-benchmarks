package day10

import (
	"coding-challenge-runner/pkg/input"
	"os"
)

type coord struct {
	row int
	col int
}

func (c coord) adjacentSteps(tm [][]int) []coord {
	height := tm[c.row][c.col]
	adjSteps := []coord{
		{c.row - 1, c.col},
		{c.row, c.col + 1},
		{c.row + 1, c.col},
		{c.row, c.col - 1},
	}
	var nextSteps []coord

	for _, step := range adjSteps {
		// fmt.Printf("step %+v\n", step)
		if step.row > -1 && step.row < len(tm) && step.col > -1 && step.col < len(tm[0]) && tm[step.row][step.col]-height == 1 {
			nextSteps = append(nextSteps, step)
		}
	}
	// fmt.Printf("next steps %+v\n", nextSteps)
	return nextSteps
}

func Part1(f *os.File) int64 {
	trailHeads, trailMap := parseInput(f)

	scoreSum := 0
	for _, th := range trailHeads {
		summitMap := make(map[coord]int)
		summitMap = getTrailScore(th, trailMap, summitMap)
		scoreSum += len(summitMap)
	}

	return int64(scoreSum)
}

func Part2(f *os.File) int64 {
	trailHeads, trailMap := parseInput(f)

	ratingSum := 0
	for _, th := range trailHeads {
		summitMap := make(map[coord]int)
		summitMap = getTrailScore(th, trailMap, summitMap)
		for _, s := range summitMap {
			ratingSum += s
		}
	}

	return int64(ratingSum)
}

func parseInput(f *os.File) ([]coord, [][]int) {
	var trailMap [][]int
	var trailHeads []coord

	i := 0
	for l := range input.Lines(f) {
		var tops []int
		for j, c := range l {
			n := int(c - '0')
			tops = append(tops, n)
			if n == 0 {
				trailHeads = append(trailHeads, coord{i, j})
			}
		}
		trailMap = append(trailMap, tops)
		i++
	}
	return trailHeads, trailMap
}

func getTrailScore(th coord, tm [][]int, summitMap map[coord]int) map[coord]int {
	if tm[th.row][th.col] == 9 {
		// fmt.Println("here")
		c, ok := summitMap[coord{th.row, th.col}]
		if !ok {
			summitMap[coord{th.row, th.col}] = 1
		} else {
			summitMap[coord{th.row, th.col}] = c + 1
		}
	}

	adj := th.adjacentSteps(tm)
	// fmt.Printf("%+v\n", th)
	for _, step := range adj {
		getTrailScore(step, tm, summitMap)
	}

	return summitMap
}
