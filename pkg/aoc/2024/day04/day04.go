package day04

import (
	"coding-challenge-runner/pkg/input"
	"fmt"
	"os"
	"unicode/utf8"
)

type crossword []string

func (cw crossword) at(row, col int, dir direction) byte {
	r := row + dir.Offset()[0]
	c := col + dir.Offset()[1]
	return cw[r][c]
}

func Part1(f *os.File) int64 {
	var found int64 = 0
	cw := parseCrossword(f)

	// for each char in crossword
	// if curr char is 'X'
	// look for an 'M' and store the direction
	// recurse until we get to the end of the "XMAS" sequence, or the curr char is not in the sequence.

	for i, l := range cw {
		for j, c := range l {
			if res := search("XMAS", c, i, j, unknown, cw); res > 0 {
				found += int64(res)
				fmt.Printf("found %d!\n", found)
			}
		}
	}

	return found
}

func Part2(f *os.File) int64 {
	var found int64 = 0
	cw := parseCrossword(f)

	for i, l := range cw {
		for j, c := range l {
			if c == 'A' {
				if isCrossMas(i, j, cw) {
					found++
				}
			}
		}
	}

	return found
}

func parseCrossword(f *os.File) []string {
	var lines []string

	for l := range input.Lines(f) {
		lines = append(lines, l)
	}

	return lines
}

type direction uint8

const (
	unknown direction = iota
	left
	leftup
	up
	rightup
	right
	rightdown
	down
	leftdown
)

func (d direction) String() string {
	switch d {
	case unknown:
		return "unknown"
	case left:
		return "left"
	case leftup:
		return "leftup"
	case up:
		return "up"
	case rightup:
		return "rightup"
	case right:
		return "right"
	case rightdown:
		return "rightdown"
	case down:
		return "down"
	case leftdown:
		return "leftdown"
	default:
		return ""
	}
}

func (d direction) Offset() []int {
	switch d {
	case left:
		return []int{0, -1}
	case leftup:
		return []int{-1, -1}
	case up:
		return []int{-1, 0}
	case rightup:
		return []int{-1, 1}
	case right:
		return []int{0, 1}
	case rightdown:
		return []int{1, 1}
	case down:
		return []int{1, 0}
	case leftdown:
		return []int{1, -1}
	default:
		return []int{0, 0}
	}
}

func search(target string, c rune, row, col int, dir direction, crossword []string) int {
	fmt.Printf("searching for %s; curr letter %s at %d,%d; dir %s\n", target, string(c), row, col, dir)
	found := 0

	if target[0] != byte(c) {
		return 0
	}

	if len(target) == 1 && target[0] == byte(c) {
		return 1
	}

	neighbors := neighbors(row, col, crossword)

	if dir == unknown {
		// search all neighbors and set a direction to follow if next is found
		for _, n := range neighbors {
			found += search(target[1:], n.c, n.row, n.col, n.dir, crossword)
		}
	}

	// Keep searchning in the same direction
	nextNeighbor := neighborByDir(neighbors, dir)
	if nextNeighbor != nil {
		return search(target[1:], nextNeighbor.c, nextNeighbor.row, nextNeighbor.col, nextNeighbor.dir, crossword)
	}

	// Couldn't find anymore neighbors to search
	return found
}

type neighbor struct {
	c   rune
	row int
	col int
	dir direction
}

func neighbors(row, col int, mat []string) []neighbor {
	possibleDirs := []direction{
		left,
		leftup,
		up,
		rightup,
		right,
		rightdown,
		down,
		leftdown,
	}
	var neighbors []neighbor

	for _, pd := range possibleDirs {
		r := row + pd.Offset()[0]
		c := col + pd.Offset()[1]

		if r > -1 && r < len(mat) && c > -1 && c < len(mat[0]) {
			decoded, _ := utf8.DecodeRune([]byte{mat[r][c]})
			neighbors = append(neighbors, neighbor{
				c:   decoded,
				row: r,
				col: c,
				dir: pd,
			})
		}
	}

	return neighbors
}

func neighborByDir(neighbors []neighbor, dir direction) *neighbor {
	for _, n := range neighbors {
		if n.dir == dir {
			return &n
		}
	}

	return nil
}

func isCrossMas(row, col int, mat crossword) bool {
	// Eliminate borders
	if row == 0 || row == len(mat)-1 || col == 0 || col == len(mat[0])-1 {
		return false
	}
	fmt.Printf("searching at %d,%d\n", row, col)

	hasBackCross := (mat.at(row, col, leftup) == 'M' && mat.at(row, col, rightdown) == 'S') || (mat.at(row, col, leftup) == 'S' && mat.at(row, col, rightdown) == 'M')
	hasForwardCross := (mat.at(row, col, rightup) == 'M' && mat.at(row, col, leftdown) == 'S') || (mat.at(row, col, rightup) == 'S' && mat.at(row, col, leftdown) == 'M')

	return hasBackCross && hasForwardCross
}
