package day08

import (
	"coding-challenge-runner/pkg/input"
	"fmt"
	"math"
	"os"
)

type location struct {
	coord
	freq      rune
	antiNodes []rune
}

type coord struct {
	row int
	col int
}

func Part1(f *os.File) int64 {
	var sigMap [][]location
	freqMap := make(map[rune][]coord)

	i := 0
	for l := range input.Lines(f) {
		var locs []location
		for j, c := range l {
			locs = append(locs, location{coord{i, j}, c, []rune{}})
			if c != '.' {
				coords, ok := freqMap[c]
				if !ok {
					freqMap[c] = []coord{{i, j}}
				} else {
					freqMap[c] = append(coords, coord{i, j})
				}
			}
		}
		sigMap = append(sigMap, locs)
		i++
	}
	fmt.Printf("freqs: %+v\n", freqMap)

	antiNodesMap := make(map[coord]int)

	for freq, coords := range freqMap {
		antiNodes := getAntiNodes(freq, coords, sigMap)
		fmt.Printf("freq %s has nodes %+v\n", string(freq), antiNodes)
		for _, node := range antiNodes {
			_, ok := antiNodesMap[node]
			if !ok {
				antiNodesMap[node] = 1
			} else {
				antiNodesMap[node]++
			}
		}
	}
	printMap(sigMap, antiNodesMap)
	return int64(len(antiNodesMap))
}

func Part2(f *os.File) int64 {
	return 0
}

func getAntiNodes(freq rune, freqCoords []coord, sigMap [][]location) []coord {
	mapSize := len(sigMap)
	var coords []coord

	if len(freqCoords) == 1 {
		// no pairs
		return coords
	}

	pairs := getPairs(freqCoords)
	fmt.Printf("pairs for %s: %+v\n", string(freq), pairs)
	for _, p := range pairs {
		p1 := p[0]
		p2 := p[1]

		dx := int(math.Abs(float64(p1.col - p2.col)))
		dy := int(math.Abs(float64(p1.row - p2.row)))

		currCoord := coord{}
		if p1.row < p2.row {
			if p1.col > p2.col {
				// up and right
				currCoord.row = p1.row - dy
				currCoord.col = p1.col + dx
			} else {
				// up and left
				currCoord.row = p1.row - dy
				currCoord.col = p1.col - dx
			}
		} else {
			if p1.col > p2.col {
				// down and right
				currCoord.row = p1.row + dy
				currCoord.col = p1.col + dx
			} else {
				// down and left
				currCoord.row = p1.row + dy
				currCoord.col = p1.col - dx
			}
		}

		if currCoord.row >= 0 && currCoord.row < mapSize && currCoord.col >= 0 && currCoord.col < mapSize {
			coords = append(coords, currCoord)
		}
	}

	return coords
}

func getPairs(coords []coord) [][]coord {
	var pairs [][]coord

	for i, c1 := range coords {
		for j, c2 := range coords {
			if i == j {
				continue
			}

			pairs = append(pairs, []coord{c1, c2})
		}
	}

	return pairs
}

func slope(p1, p2 coord) float64 {
	return float64((p2.row - p1.row) / (p2.col - p1.col))
}

func printMap(sigMap [][]location, antiNodesMap map[coord]int) {
	for _, rows := range sigMap {
		for _, loc := range rows {
			if loc.freq != '.' {
				fmt.Printf("%s", string(loc.freq))
				continue
			}
			if _, ok := antiNodesMap[loc.coord]; ok {
				fmt.Print("#")
				continue
			}
			fmt.Print(".")
		}
		fmt.Println()
	}
}
