package day09

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1(f *os.File) int64 {
	buf := make([]byte, 100000)
	_, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	diskMap := bytes.Trim(buf, "\x00")
	blocks := getBlocks(string(diskMap))
	printDisk(blocks)

	packed := packBlocks(blocks)
	// printDisk(packed)

	return calcChecksum(packed)
}

func Part2(f *os.File) int64 {
	buf := make([]byte, 100000)
	_, err := f.Read(buf)
	if err != nil {
		panic(err)
	}

	diskMap := bytes.Trim(buf, "\x00")
	groups := getBlockGroups(string(diskMap))
	printDiskGroups(groups)

	packed := packFiles(groups)
	printDiskGroups(packed)

	return calcGroupChecksum(packed)
}

type block struct {
	f *file
}

func (b block) isEmpty() bool {
	return b.f == nil
}

type file struct {
	id   int
	size int
}

type blockGroup struct {
	id   int
	size int
}

func (bg blockGroup) isEmpty() bool {
	return bg.id == -1
}

func (bg blockGroup) blocks() []block {
	var blocks []block

	for i := 0; i < bg.size; i++ {
		if bg.isEmpty() {
			blocks = append(blocks, block{nil})
		} else {
			blocks = append(blocks, block{&file{bg.id, bg.size}})
		}
	}

	return blocks
}

// maybe have this take []int instead of string
func getBlocks(diskMap string) []block {
	var blocks []block
	currFileId := 0

	for i, f := range diskMap {
		size := int(f - '0')
		if i%2 == 0 {
			// Expanding a file
			for i := 0; i < size; i++ {
				blocks = append(blocks, block{
					&file{currFileId, size},
				})
			}
			currFileId++
			continue
		}
		// Expanding free space
		for i := 0; i < size; i++ {
			blocks = append(blocks, block{nil})
		}
	}

	return blocks
}

func getBlockGroups(diskMap string) []blockGroup {
	var groups []blockGroup
	currFileId := 0

	for i, f := range diskMap {
		size := int(f - '0')
		var g blockGroup
		if i%2 == 0 {
			g = blockGroup{currFileId, size}
			currFileId++
		} else {
			g = blockGroup{-1, size}
		}
		groups = append(groups, g)
	}

	return groups
}

func packBlocks(blocks []block) []block {
	packed := make([]block, len(blocks))
	i := 0
	j := len(blocks) - 1

	for {
		if i > j {
			break
		}
		b := blocks[i]

		if blocks[j].isEmpty() {
			j--
			continue
		}

		if b.isEmpty() {
			packed[i] = blocks[j]
			j--
		} else {
			packed[i] = b
		}
		i++
	}

	return packed
}

func packFiles(groups []blockGroup) []blockGroup {
	// visited := make(map[int]int)
	currFile := groups[len(groups)-1]
	for j := currFile.id; j > -1; j-- {
		var g blockGroup
		var k int
		for k = len(groups) - 1; k > 0; k-- {
			g = groups[k]
			if g.isEmpty() || g.id != j {
				continue
			} else {
				break
			}
		}
		// g := groups[j]
		// if g.isEmpty() {
		// 	continue
		// }
		// _, ok := visited[g.id]
		// if ok {
		// 	continue
		// }
		// visited[g.id] = 1
		// look for empty space starting from left
		for i, g2 := range groups {
			if i > k {
				break
			}
			if g2.isEmpty() && g2.size >= g.size {
				// reduce space at i
				groups[i].size -= g.size
				// remove g at j
				groups[k] = blockGroup{-1, g.size}
				// insert g at i
				groups = slices.Insert(groups, i, g)
				break
			}
		}
		groups = consolidateEmptyGroups(groups)
		// printDiskGroups(groups)
	}

	return groups
}

func consolidateEmptyGroups(groups []blockGroup) []blockGroup {
	var consolidated []blockGroup
	emptyGroup := blockGroup{-1, 0}

	for _, g := range groups {
		if !g.isEmpty() {
			if emptyGroup.size > 0 {
				consolidated = append(consolidated, emptyGroup)
				emptyGroup = blockGroup{-1, 0}
			}
			consolidated = append(consolidated, g)
			continue
		}
		emptyGroup.size += g.size
	}

	return consolidated
}

func printDisk(blocks []block) {
	for _, b := range blocks {
		if b.f != nil {
			idStr := strconv.Itoa(b.f.id)
			fmt.Print(idStr)
			continue
		}

		fmt.Print(".")
	}
	fmt.Println()
}

func printDiskGroups(groups []blockGroup) {
	for _, g := range groups {
		if g.isEmpty() {
			fmt.Print(strings.Repeat(".", g.size))
		} else {
			idStr := strconv.Itoa(g.id)
			fmt.Print(strings.Repeat(idStr, g.size))
		}
	}
	fmt.Println()
}

func calcChecksum(blocks []block) int64 {
	var checkSum int64 = 0

	for i, b := range blocks {
		if !b.isEmpty() {
			checkSum += int64(i * b.f.id)
		}
	}

	return checkSum
}

func calcGroupChecksum(groups []blockGroup) int64 {
	var blocks []block

	for _, g := range groups {
		blocks = append(blocks, g.blocks()...)
	}

	return calcChecksum(blocks)
}
