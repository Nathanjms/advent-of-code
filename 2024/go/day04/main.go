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
			if contents[i][j] == 'X' {
				qty := checkForXmas(i, j, &contents)
				xmasCount += qty
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
	crossMasCount := 0
	// Go through the grid, and for every X we encounter, check for the word Xmas
	for i, line := range contents {
		for j := range line {
			if contents[i][j] == 'A' && checkForCrossMas(i, j, &contents) {
				crossMasCount++
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   4,
		Part:  2,
		Value: crossMasCount,
	})
}

func checkForXmas(i int, j int, contents *[]string) int {
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

	matches := 0

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
			matches++
		}
	}

	return matches

}

func checkForCrossMas(i int, j int, contents *[]string) bool {
	// Manually check this time as fewer cases

	// First check boundaries are not exceeded
	if i-1 < 0 || j-1 < 0 || j+1 > len(*contents)-1 || i+1 > len(*contents)-1 {
		// Out of range this way
		return false
	}

	// first check left diag - \
	if ((*contents)[i-1][j-1] == 'M' && (*contents)[i+1][j+1] == 'S') || ((*contents)[i-1][j-1] == 'S' && (*contents)[i+1][j+1] == 'M') {
		//safe to continue if here - mainly using the else
	} else {
		return false
	}

	// next check right diag - /
	if ((*contents)[i-1][j+1] == 'M' && (*contents)[i+1][j-1] == 'S') || ((*contents)[i-1][j+1] == 'S' && (*contents)[i+1][j-1] == 'M') {
		return true
	}

	return false

}
