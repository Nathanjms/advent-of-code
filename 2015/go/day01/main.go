package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
)

func getCurrentDirectory() string {
	_, filename, _, _ := runtime.Caller(0)
	dirname := filepath.Dir(filename)
	return dirname
}

// Default Input path is current directory + example-input
var inputPath = filepath.Join(getCurrentDirectory(), "example-input")

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents string) {
	floorNumber := 0

	// Go from both left and right to half the time (I think this is faster than just going from left)
	leftIndex := 0
	rightIndex := len(contents) - 1

	for leftIndex < len(contents) {
		// If the left is larger than the right, we're past the middle so stop
		if leftIndex > rightIndex {
			break
		}

		if contents[leftIndex] == '(' {
			floorNumber++
		} else if contents[leftIndex] == ')' {
			floorNumber--
		}

		// If the left and right are the same, we're at the middle, so we don't need to repeat the right side
		if rightIndex == leftIndex {
			break
		}

		if contents[rightIndex] == '(' {
			floorNumber++
		}
		if contents[rightIndex] == ')' {
			floorNumber--
		}

		leftIndex++
		rightIndex--
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   1,
		Part:  1,
		Value: floorNumber,
	})
}

func partTwo(contents string) {
	floorNumber := 0

	// Go from just the left this time

	indexOfBasement := 0

	for i := 0; i < len(contents); i++ {
		if contents[i] == '(' {
			floorNumber++
		} else if contents[i] == ')' {
			floorNumber--
		}
		if floorNumber == -1 {
			indexOfBasement = i
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   1,
		Part:  1,
		Value: indexOfBasement + 1,
	})
}
