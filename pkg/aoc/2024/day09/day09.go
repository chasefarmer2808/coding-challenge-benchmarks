package day09

import (
	"bytes"
	"fmt"
	"os"
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

	//packed := packBlocks(blocks)
	//fmt.Println(packed)

	return 0
}

func Part2(f *os.File) int64 {
	return 0
}

type block struct {
	f *file
}

type file struct {
	id   int
	size int
}

// maybe have this take []int instead of string
// TODO: store a list of blocks.  Each block is a free space OR part of a file.
func getBlocks(diskMap string) []block {
	var blocks []block
	currFileId := 0

	for i, f := range diskMap {
		size := int(f - '0')
		if i%2 == 0 {
			// Expanding a file
			blocks = append(blocks, block{
				size,
				&file{currFileId},
			})
			currFileId++
			continue
		}
		// Expanding free space
		blocks = append(blocks, block{size, nil})
	}

	return blocks
}

func packBlocks(blocks []block) []block {

}

func printDisk(blocks []block) {
	for _, b := range blocks {
		if b.f != nil {
			idStr := strconv.Itoa(b.f.id)
			fmt.Print(strings.Repeat(idStr, b.size))
			continue
		}

		fmt.Print(strings.Repeat(".", b.size))
	}
	fmt.Println()
}
