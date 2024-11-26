package day03

import (
	"coding-challenge-runner/pkg/input"
	"os"
	"strconv"
	"unicode"
)

// How to store schematic?
// double char array or array of strings?

func Part1(f *os.File) int {
	var schematic []string
	sum := 0

	for l := range input.Lines(f) {
		schematic = append(schematic, l)
	}

	// Iterate through schematic and detect numbers
	for i := 0; i < len(schematic); i++ {
		s := schematic[i]
		numStr := ""

		for j := 0; j < len(s); j++ {
			c := rune(s[j])
			if unicode.IsNumber(c) {
				numStr += string(c)
				continue
			}

			if len(numStr) > 0 {
				// look around number
				isPartNum := touchesSymbol(i, j-len(numStr), len(numStr), schematic)
				//fmt.Printf("%s, %v", numStr, isPartNum)
				//fmt.Println()
				if isPartNum {
					n, _ := strconv.Atoi(numStr)
					sum += n
				}
				numStr = ""
			}
		}
		if len(numStr) > 0 {
			// look around number
			isPartNum := touchesSymbol(i, len(schematic[0])-len(numStr), len(numStr), schematic)
			//fmt.Printf("%s, %v", numStr, isPartNum)
			//fmt.Println()
			if isPartNum {
				n, _ := strconv.Atoi(numStr)
				sum += n
			}
			numStr = ""
		}
	}

	return sum
}

func Part2(f *os.File) int {
	return 0
}

type coord struct {
	i int
	j int
}

func touchesSymbol(row, col, size int, mat []string) bool {
	// Get coord around number
	var coords []coord

	if row > 0 {
		// top
		for i := 0; i < size; i++ {
			coords = append(coords, coord{row - 1, col + i})
		}
	}

	if row < len(mat)-1 {
		// bottom
		for i := 0; i < size; i++ {
			coords = append(coords, coord{row + 1, col + i})
		}
	}

	if col > 0 {
		// left
		coords = append(coords, coord{row, col - 1})
	}

	if col+size < len(mat[0]) {
		// right
		coords = append(coords, coord{row, col + size})
	}

	if row > 0 && col > 0 {
		// top left corner
		coords = append(coords, coord{row - 1, col - 1})
	}

	if row < len(mat)-1 && col > 0 {
		// bottom left corner
		coords = append(coords, coord{row + 1, col - 1})
	}

	if row > 0 && col+size < len(mat[0]) {
		// top right corner
		coords = append(coords, coord{row - 1, col + size})
	}

	if row < len(mat)-1 && col+size < len(mat[0]) {
		// bottom right corner
		coords = append(coords, coord{row + 1, col + size})
	}
	//fmt.Println(coords)
	for _, c := range coords {
		r := rune(mat[c.i][c.j])

		if !unicode.IsNumber(r) && r != '.' {
			return true
		}
	}

	return false
}
