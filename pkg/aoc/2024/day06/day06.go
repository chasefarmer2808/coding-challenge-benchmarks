package day06

import (
	"coding-challenge-runner/pkg/input"
	"fmt"
	"os"
	"strings"
)

type coord struct {
	row int
	col int
}

type direction uint8

const (
	left direction = iota
	up
	right
	down
)

func Part1(f *os.File) int64 {
	var labMap []string
	var guardCoord coord
	guardMap := make(map[coord]int)
	currDir := up

	// read map and get guard start coord
	i := 0
	for l := range input.Lines(f) {
		labMap = append(labMap, l)
		if j := strings.Index(l, "^"); j != -1 {
			guardCoord = coord{i, j}
		}
		i++
	}

	// Set first distinct guard pos
	guardMap[guardCoord] = 1

	// move guard one step at a time until out of bounds
	for {
		if guardCoord.row == len(labMap) || guardCoord.col == len(labMap[0]) {
			break
		}
		if facingObstruction(guardCoord, currDir, labMap) {
			currDir = nextDir(currDir)
			continue
		}
		// add guard current coord to coord map.
		switch currDir {
		case left:
			guardCoord.col--
		case up:
			guardCoord.row--
		case right:
			guardCoord.col++
		case down:
			guardCoord.row++
		}
		n, ok := guardMap[guardCoord]
		if !ok {
			guardMap[guardCoord] = 1
		} else {
			guardMap[guardCoord] = n + 1
		}
	}

	return int64(len(guardMap) - 1)
}

func Part2(f *os.File) int64 {
	var possibleObsticles int64 = 0
	var labMap []string
	guardMap := make(map[coord]direction)
	var guardCoord coord
	var startCoord coord
	currDir := up

	// repeat part 1; TODO extract common bits
	// read map and get guard start coord
	i := 0
	for l := range input.Lines(f) {
		labMap = append(labMap, l)
		if j := strings.Index(l, "^"); j != -1 {
			startCoord = coord{i, j}
		}
		i++
	}

	// Set first distinct guard pos
	guardCoord = startCoord
	guardMap[guardCoord] = up

	// move guard one step at a time until out of bounds
	for {
		if guardCoord.row == len(labMap) || guardCoord.col == len(labMap[0]) {
			break
		}
		if facingObstruction(guardCoord, currDir, labMap) {
			currDir = nextDir(currDir)
			continue
		}
		// add guard current coord to coord map.
		switch currDir {
		case left:
			guardCoord.col--
		case up:
			guardCoord.row--
		case right:
			guardCoord.col++
		case down:
			guardCoord.row++
		}
		guardMap[guardCoord] = currDir
	}

	// put an obsticle at every position where the guard goes and check for infinit cycle
	for c, _ := range guardMap {
		// reset the guard
		guardCoord = startCoord
		// try obsticle
		newLabMap := copyWithObsticle(labMap, c)

		if hasInfiniteCycle(guardCoord, up, newLabMap) {
			possibleObsticles++
		}
	}

	return possibleObsticles
}

func facingObstruction(currCoord coord, dir direction, mat []string) bool {
	switch dir {
	case left:
		currCoord.col--
	case up:
		currCoord.row--
	case right:
		currCoord.col++
	case down:
		currCoord.row++
	}

	if currCoord.row < 0 || currCoord.row == len(mat) || currCoord.col < 0 || currCoord.col == len(mat[0]) {
		return false
	}

	return mat[currCoord.row][currCoord.col] == '#'
}

func nextDir(dir direction) direction {
	switch dir {
	case left:
		return up
	case up:
		return right
	case right:
		return down
	case down:
		return left
	}

	panic(fmt.Errorf("invalid direction %d", dir))
}

func copyWithObsticle(orig []string, obPos coord) []string {
	var new []string

	for i, s := range orig {
		var line string
		for j, c := range s {
			if i == obPos.row && j == obPos.col {
				line += "#"
			} else {
				line += string(c)
			}
		}
		new = append(new, line)
	}

	return new
}

func hasInfiniteCycle(start coord, dir direction, mat []string) bool {
	guardMap := make(map[coord]direction)
	guardMap[start] = dir
	guardCoord := start
	currDir := dir

	for {
		if guardCoord.row < 0 || guardCoord.row == len(mat) || guardCoord.col < 0 || guardCoord.col == len(mat[0]) {
			return false
		}
		if facingObstruction(guardCoord, currDir, mat) {
			currDir = nextDir(currDir)
			continue
		}
		// add guard current coord to coord map.
		switch currDir {
		case left:
			guardCoord.col--
		case up:
			guardCoord.row--
		case right:
			guardCoord.col++
		case down:
			guardCoord.row++
		}

		prevDir, ok := guardMap[guardCoord]
		if ok && prevDir == currDir {
			return true
		}
		guardMap[guardCoord] = currDir
	}
}
