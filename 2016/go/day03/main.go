package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"regexp"
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

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	possibleCount := 0
	for _, line := range contents {
		var sides [3]int

		re := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)
		matches := re.FindStringSubmatch(line)

		for i, match := range matches {
			if i == 0 {
				// i = 0 returns the entire match, so skip this one
				continue
			}
			intVal, _ := strconv.Atoi(match)
			sides[i-1] = intVal
		}

		// We could make generic combinations, but as it's only 3 sided shape, let's just manually handle each
		if sides[0]+sides[1] > sides[2] && sides[0]+sides[2] > sides[1] && sides[1]+sides[2] > sides[0] {
			possibleCount++
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   3,
		Part:  1,
		Value: possibleCount,
	})
}

func partTwo(contents []string) {
	possibleCount := 0
	threeTriangles := make([][3]int, 3) // We'll build three triangles at a time, so store that in this variable
	currIndex := 0                      // Track the current index of ALL of the triangles we are building
	for _, line := range contents {
		var sides [3]int

		re := regexp.MustCompile(`(\d+)\s+(\d+)\s+(\d+)`)
		matches := re.FindStringSubmatch(line)

		for i, match := range matches {
			if i == 0 {
				// i = 0 returns the entire match, so skip this one
				continue
			}
			intVal, _ := strconv.Atoi(match)
			sides[i-1] = intVal
		}

		threeTriangles[0][currIndex] = sides[0]
		threeTriangles[1][currIndex] = sides[1]
		threeTriangles[2][currIndex] = sides[2]

		if currIndex == 2 {
			// Completed triangles ready to be compared when index is 2!
			for _, sides := range threeTriangles {
				if sides[0]+sides[1] > sides[2] && sides[0]+sides[2] > sides[1] && sides[1]+sides[2] > sides[0] {
					possibleCount++
				}
			}

			// Then reset
			threeTriangles = make([][3]int, 3)
		}
		currIndex = (currIndex + 1) % 3
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   3,
		Part:  2,
		Value: possibleCount,
	})
}
