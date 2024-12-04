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
var isUsingExample = true

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
		isUsingExample = false
	}

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	xmasCount := 0
	// Go through the grid, and for every X we encounter, check for the word Xmas
	for i, line := range contents {
		for j := range line {
			if contents[i][j] == 'X' && checkForXmas(i, j, &contents) {
				xmasCount++
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   4,
		Part:  1,
		Value: xmasCount,
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   4,
		Part:  2,
		Value: "TODO",
	})
}

func checkForXmas(i int, j int, contents *[]string) bool {
	// deifne the steps - probably a better way to do this just with [-1,0,1] but this isnt too long to write
	directions := [8][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
		{1, 1},
		{1, -1},
		{-1, 1},
		{-1, -1},
	}

	stepsToCharMap := map[int]byte{
		1: 'M',
		2: 'A',
		3: 'S',
	}

	for _, direction := range directions {
		// Check we've not exceeded the range for the current direction
		if direction[0]*3+i > len((*contents)[0])-1 ||
			direction[0]*3+i < 0 ||
			direction[1]*3+j > len((*contents))-1 ||
			direction[1]*3+j < 0 {
			continue
		}

		isAMatch := true
		// Otherwise, check for 'MAS' with 3 intervals:
		for n := 1; n <= 3; n++ {
			if (*contents)[i+(direction[0]*n)][j+(direction[1]*n)] != stepsToCharMap[n] {
				isAMatch = false
				break
			}
		}

		if isAMatch {
			return true
		}
	}

	return false

}
