package day04

import (
	"coding-challenge-runner/pkg/input"
	"fmt"
	"os"
	"unicode/utf8"
)

type node struct {
	c         rune
	neighbors []node
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
			if search("XMAS", c, i, j, unknown, cw) {
				found++
				fmt.Printf("found %d!\n", found)
			}
		}
	}

	return found
}

func Part2(f *os.File) int64 {
	return 0
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

func search(target string, c rune, row, col int, dir direction, crossword []string) bool {
	fmt.Printf("searching for %s; curr letter %s at %d,%d; dir %s\n", target, string(c), row, col, dir)

	if target[0] != byte(c) {
		return false
	}

	if len(target) == 1 && target[0] == byte(c) {
		return true
	}

	neighbors := neighbors(row, col, crossword)

	if dir == unknown {
		// search all neighbors and set a direction to follow if next is found
		for _, n := range neighbors {
			res := search(target[1:], n.c, n.row, n.col, n.dir, crossword)
			if res {
				return res
			}
		}
	}

	// Keep searchning in the same direction
	nextNeighbor := neighborByDir(neighbors, dir)
	if nextNeighbor != nil {
		return search(target[1:], nextNeighbor.c, nextNeighbor.row, nextNeighbor.col, nextNeighbor.dir, crossword)
	}

	// Couldn't find anymore neighbors to search
	return false
}

// func next(c rune) rune {
// 	switch c {
// 	case 'X':
// 		return 'M'
// 	case 'M':
// 		return 'A'
// 	case 'A':
// 		return 'S'
// 	default:
// 		return ' '
// 	}
// }

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
