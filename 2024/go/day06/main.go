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
	startPosition := determineStartIndex(contents)
	visitedPositions := map[[2]int]bool{
		startPosition: true,
	}

	currDirectionIndex := 0
	directionsInOrder := [4][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	currPosition := startPosition

	for {
		// Take step in current direction;
		newPosition := [2]int{
			currPosition[0] + directionsInOrder[currDirectionIndex][0],
			currPosition[1] + directionsInOrder[currDirectionIndex][1],
		}

		// Have we finished?
		if newPosition[0] < 0 || newPosition[0] > len(contents)-1 || newPosition[1] < 0 || newPosition[1] > len(contents[0])-1 {
			break
		}

		// Have we encountered a '#'
		if contents[newPosition[0]][newPosition[1]] == '#' {
			// If we have, change direction and go again!
			currDirectionIndex = (currDirectionIndex + 1) % 4
			continue
		}

		// Otherwise, we're free to take a step!
		visitedPositions[newPosition] = true
		currPosition = newPosition

	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   6,
		Part:  1,
		Value: len(visitedPositions),
	})
}

func partTwo(contents []string) {
	obstructionPositions := make(map[[2]int]bool, 0)

	for i := 0; i < len(contents); i++ {
		for j := 0; j < len(contents[i]); j++ {
			if contents[i][j] == '#' {
				continue // Already has one so cant be here
			}

			if contents[i][j] == '^' {
				continue
			}

			testContents := make([]string, len(contents))

			copy(testContents, contents)
			testContents[i] = replaceAtIndex(testContents[i], '#', j)

			if testIfRepeat(testContents) {
				obstructionPositions[[2]int{i, j}] = true
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   6,
		Part:  2,
		Value: len(obstructionPositions),
	})

}

func testIfRepeat(contents []string) bool {
	startPosition := determineStartIndex(contents)

	currDirectionIndex := 0
	directionsInOrder := [4][2]int{
		{-1, 0}, // up
		{0, 1},  // right
		{1, 0},  // down
		{0, -1}, // left
	}

	visitedPositionAndDirection := map[[3]int]bool{
		{startPosition[0], startPosition[1], currDirectionIndex}: true,
	}

	currPosition := startPosition
	steps := 0
	for {
		// Take step in current direction;
		newPosition := [2]int{
			currPosition[0] + directionsInOrder[currDirectionIndex][0],
			currPosition[1] + directionsInOrder[currDirectionIndex][1],
		}

		// Have we finished?
		if newPosition[0] < 0 || newPosition[0] > len(contents)-1 || newPosition[1] < 0 || newPosition[1] > len(contents[0])-1 {
			return false
		}

		// Have we encountered a '#'
		if contents[newPosition[0]][newPosition[1]] == '#' {
			// If we have, change direction and go again!
			currDirectionIndex = (currDirectionIndex + 1) % 4
			continue
		}

		// Otherwise, we're free to take a step!
		currPosition = newPosition
		steps++

		// have we been here before?
		_, ok := visitedPositionAndDirection[[3]int{currPosition[0], currPosition[1], currDirectionIndex}]
		// If the key exists
		if ok {
			return true
		}

		visitedPositionAndDirection[[3]int{currPosition[0], currPosition[1], currDirectionIndex}] = true
	}
}

func determineStartIndex(contents []string) [2]int {
	for i, line := range contents {
		for j := 0; j < len(line); j++ {
			if contents[i][j] == '^' {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{-1, -1}
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
