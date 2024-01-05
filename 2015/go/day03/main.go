package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
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

var DIRECTION_MAPPINGS = map[string][2]int{
	"^": {0, 1},
	"v": {0, -1},
	">": {1, 0},
	"<": {-1, 0},
}

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
	currentPosition := [2]int{0, 0}
	visited := map[[2]int]int{
		currentPosition: 1,
	}

	visitCount := 1 // Start on 1 house

	for _, direction := range contents {
		currentPosition[0] += DIRECTION_MAPPINGS[string(direction)][0]
		currentPosition[1] += DIRECTION_MAPPINGS[string(direction)][1]

		// If we haven't visited this position before, increment the visit count and add it to the map:
		if _, ok := visited[currentPosition]; !ok {
			visitCount++
			visited[currentPosition] = 1
		}
	}

	fmt.Println(visitCount)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   3,
		Part:  1,
		Value: visitCount,
	})
}

func partTwo(contents string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   3,
		Part:  1,
		Value: "TODO",
	})
}
