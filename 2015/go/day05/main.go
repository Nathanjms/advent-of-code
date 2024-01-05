package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
	"strings"
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

const VOWELS = "aeiou"

var BAD_STRINGS = []string{"ab", "cd", "pq", "xy"}

func partOne(contents []string) {
	niceStrings := 0

	for _, str := range contents {
		vowelCount := 0
		hasDoubleLetter := false
		hasBadString := false

		// Go through each character in the string
		for i := 0; i < len(str); i++ {
			// First check for vowels (if not already found)
			if vowelCount < 3 && strings.Contains(VOWELS, string(str[i])) {
				vowelCount++
			}

			// Next check if this letter is the same as the next one
			if (i < len(str)-1) && (str[i] == str[i+1]) {
				hasDoubleLetter = true
			}

			// Finally, check it has no bad strings
			for _, badString := range BAD_STRINGS {
				if strings.Contains(str, badString) {
					hasBadString = true
				}
			}
			if hasBadString {
				break // No need to continue checking if we know it's bad
			}

		}

		if vowelCount >= 3 && hasDoubleLetter && !hasBadString {
			niceStrings++
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   5,
		Part:  1,
		Value: niceStrings,
	})
}

func partTwo(contents []string) {
	niceStrings := 0

	for _, str := range contents {
		hasPair := false
		hasRepeat := false

		// Go through each character in the string
		for i := 0; i < len(str); i++ {
			// Easier one - check for repeats
			if !hasRepeat && i < len(str)-2 {
				if str[i] == str[i+2] {
					hasRepeat = true
				}
			}

			// Harder one - check for pairs
			if !hasPair && i < len(str)-1 {
				pair := string(str[i]) + string(str[i+1])
				if strings.Count(str, pair) > 1 {
					hasPair = true
				}
			}

			if hasPair && hasRepeat {
				break // No need to continue checking if we know it's nice
			}
		}

		if hasPair && hasRepeat {
			niceStrings++
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   5,
		Part:  2,
		Value: niceStrings,
	})
}
