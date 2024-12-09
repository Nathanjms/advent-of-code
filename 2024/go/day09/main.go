package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

func getCurrentDirectory() string {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	return dirname
}

// Default Input path is current directory + example-input
var inputPath = filepath.Join(getCurrentDirectory(), "example-input")
var isUsingExample = true

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
		isUsingExample = false
	}

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

type FileBlock struct {
	id        int64
	size      int64
	freeSpace int64
}

func partOne(contents string) {
	fileBlocks := parseInput(contents)

	// Instead of building a very long string with dots in, we'll build an array which each contains the multiplier and ID
	// index 0 is ID index 1 is qty
	reorderedFiles := make([][2]int64, 0)
	for i := 0; i < len(fileBlocks); i++ {
		fileBlock := fileBlocks[i]
		// Go upwards and whenever we reach a free space, look for the LAST element
		reorderedFiles = append(reorderedFiles, [2]int64{fileBlock.id, fileBlock.size})

		if fileBlock.freeSpace > 0 {
			toAdd := takeFromEnd(fileBlock.id, fileBlock.freeSpace, &fileBlocks)
			reorderedFiles = append(reorderedFiles, toAdd...)
		}
	}

	sum := int64(0)
	currentPosition := int64(0)
	for i := 0; i < len(reorderedFiles); i++ {
		for j := 0; j < int(reorderedFiles[i][1]); j++ {
			sum += currentPosition * reorderedFiles[i][0]
			currentPosition++
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   9,
		Part:  1,
		Value: sum,
	})
}

func takeFromEnd(currentFileId int64, freeSpace int64, fileBlocks *[]FileBlock) [][2]int64 {
	returnVal := make([][2]int64, 0)
	if freeSpace == 0 {
		return returnVal
	}

	for j := len(*fileBlocks) - 1; j > 0; j-- {
		if freeSpace == 0 {
			break
		}
		// Start from the last and see if we can build this array, working downwards until we find one with size remaining
		fileBlock := (*fileBlocks)[j]

		if currentFileId == fileBlock.id {
			break
		}

		if fileBlock.size > 0 {
			amountUsed := int64(0)
			if fileBlock.size <= freeSpace {
				// All of it
				amountUsed = fileBlock.size
			} else {
				amountUsed = freeSpace
			}

			returnVal = append(returnVal, [2]int64{fileBlock.id, amountUsed})
			(*fileBlocks)[j].size = fileBlock.size - amountUsed
			freeSpace -= amountUsed
		}
	}

	return returnVal
}

func partTwo(contents string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   9,
		Part:  2,
		Value: "TODO",
	})
}

func parseInput(contents string) []FileBlock {
	fileBlocks := make([]FileBlock, 0)
	currentId := int64(0)
	for i := 0; i < len(contents); i = i + 2 {
		size, err := strconv.ParseInt(string(contents[i]), 10, 64)
		if err != nil {
			panic(err)
		}
		lastIndex := i + 1
		freeSpace := int64(0)
		if i+1 < len(contents) {
			freeSpace, err = strconv.ParseInt(string(contents[lastIndex]), 10, 64)
		}
		if err != nil {
			panic(err)
		}

		fileBlocks = append(fileBlocks, FileBlock{
			id:        currentId,
			size:      size,
			freeSpace: freeSpace,
		})
		currentId++
	}

	return fileBlocks
}
