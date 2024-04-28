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
	discs := parseInput(contents)

	// Begin by just going up each second, and seeing if we get through every disc
	time := 0
	for {
		hasFailed := false
		for i, disc := range discs {
			// The position of the disc will be it's initial position, plus the number of 'steps' (seconds), aka time + initial position + discNumber (or i+1)
			// We can take modulo of the number of positions, since it resets to 0 at the end
			if (time+disc.initialPosition+i+1)%disc.numPositions != 0 {
				hasFailed = true
				break
			}
		}
		if hasFailed {
			time++
		} else {
			break
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   15,
		Part:  1,
		Value: time,
	})
}

func partTwo(contents []string) {
	discs := parseInput(contents)
	discs = append(discs, disc{len(discs) + 1, 11, 0})

	// Begin by just going up each second, and seeing if we get through every disc
	time := 0
	for {
		hasFailed := false
		for i, disc := range discs {
			if (time+disc.initialPosition+i+1)%disc.numPositions != 0 {
				hasFailed = true
				break
			}
		}
		if hasFailed {
			time++
		} else {
			break
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   15,
		Part:  2,
		Value: time,
	})
}

type disc struct {
	key             int
	numPositions    int
	initialPosition int
}

func parseInput(contents []string) []disc {
	discs := make([]disc, 0)
	for _, line := range contents {
		var key, numPositions, initialPosition int
		_, err := fmt.Sscanf(line, "Disc #%d has %d positions; at time=0, it is at position %d.", &key, &numPositions, &initialPosition)
		if err != nil {
			panic(err)
		}
		discs = append(discs, disc{key, numPositions, initialPosition})
	}
	return discs
}
