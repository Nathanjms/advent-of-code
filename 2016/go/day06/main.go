package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
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
	frequencyMapPerRow := make([]map[byte]int, len(contents[0]))
	for i := range frequencyMapPerRow {
		frequencyMapPerRow[i] = make(map[byte]int)
	}
	for i := 0; i < len(contents[0]); i++ {
		mapping := frequencyMapPerRow[i]
		for _, char := range contents {
			if _, ok := mapping[char[i]]; !ok {
				mapping[char[i]] = 1
			} else {
				mapping[char[i]]++
			}
		}
	}

	name := ""
	// Now Loop to get the most common char
	for _, mapping := range frequencyMapPerRow {
		var largestKey byte
		currVal := 0
		for key, val := range mapping {
			if val > currVal {
				largestKey = key
				currVal = val
			}
		}
		name += string(largestKey)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   6,
		Part:  1,
		Value: name,
	})
}

func partTwo(contents []string) {
	frequencyMapPerRow := make([]map[byte]int, len(contents[0]))
	for i := range frequencyMapPerRow {
		frequencyMapPerRow[i] = make(map[byte]int)
	}
	for i := 0; i < len(contents[0]); i++ {
		mapping := frequencyMapPerRow[i]
		for _, char := range contents {
			if _, ok := mapping[char[i]]; !ok {
				mapping[char[i]] = 1
			} else {
				mapping[char[i]]++
			}
		}
	}

	name := ""
	// Now Loop to get the most common char
	for _, mapping := range frequencyMapPerRow {
		var smallestChar byte
		currVal := math.MaxInt
		for key, val := range mapping {
			if val < currVal {
				smallestChar = key
				currVal = val
			}
		}
		name += string(smallestChar)
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   6,
		Part:  2,
		Value: name,
	})
}
