package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

func partOne(contents string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   9,
		Part:  1,
		Value: decompress(contents, false),
	})
}

func partTwo(contents string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   9,
		Part:  2,
		Value: decompress(contents, true),
	})
}

// Recursive function to handle getting the length - no need to build a string (and building one fot pt2 is too large)
func decompress(word string, isPartTwo bool) int {
	hasBracket := false
	bracketIndex := 0
	for i := 0; i < len(word); i++ {
		if word[i] == '(' {
			bracketIndex = i
			hasBracket = true
			break
		}
	}

	if !hasBracket {
		return len(word)
	}

	// If here, then there is a bracket, so we have to handle this!

	indexTmp := bracketIndex + 1
	for {
		indexTmp++
		if word[indexTmp] == ')' {
			break
		}
	}

	split := strings.Split(word[bracketIndex+1:indexTmp], "x")
	repeats, _ := strconv.Atoi(split[1])
	repeatLength, _ := strconv.Atoi(split[0])

	if !isPartTwo {
		// For part one, add the amount of the modifier then move on
		return len(word[:bracketIndex]) +
			repeatLength*repeats +
			decompress(word[indexTmp+1+repeatLength:], isPartTwo)
	}

	// For part two, we need to decompress the bits inside the modifier, by the repeats
	return len(word[:bracketIndex]) +
		decompress(word[indexTmp+1:indexTmp+1+repeatLength], isPartTwo)*repeats +
		decompress(word[indexTmp+1+repeatLength:], isPartTwo)

}
