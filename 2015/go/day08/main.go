package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
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
	result := 0
	for _, line := range contents {
		result += len(line) // Add the memory size

		// Unquote the string to handle escape characters
		unquotedString, err := strconv.Unquote(line)
		if err != nil {
			panic("Error :c")
		}

		result -= len(unquotedString) // Subtract the character count after converting to utf8
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   8,
		Part:  1,
		Value: result,
	})
}

func partTwo(contents []string) {
	result := 0
	for _, line := range contents {
		quotedString := strconv.Quote(line)
		result += len(quotedString) // Add the count of the encoded-encoded string

		result -= len(line) // Subtract the character count after converting to utf8
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   8,
		Part:  2,
		Value: result,
	})
}
