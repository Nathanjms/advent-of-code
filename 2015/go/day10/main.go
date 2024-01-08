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
	result := contents
	for i := 0; i < 40; i++ {
		result = lookAndSay(result)
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   10,
		Part:  1,
		Value: len(result),
	})
}

func partTwo(contents string) {
	result := contents
	for i := 0; i < 50; i++ {
		result = lookAndSay(result)
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   10,
		Part:  2,
		Value: len(result),
	})
}

// Because strings are immutable in Go, the performance of appending to a string is really slow. Use builder instead.
func lookAndSay(input string) string {
	var builder strings.Builder

	currentLetter := input[0]
	currentLetterCount := 1

	for i := 1; i < len(input); i++ {
		if input[i] == currentLetter {
			currentLetterCount++
		} else {
			builder.WriteString(strconv.Itoa(currentLetterCount))
			builder.WriteByte(currentLetter)

			currentLetter = input[i]
			currentLetterCount = 1
		}
	}

	builder.WriteString(strconv.Itoa(currentLetterCount))
	builder.WriteByte(currentLetter)

	return builder.String()
}

// --------------------------------------

/* Initial implementation, using strings and strconv - this was really slow for n = 50 and so a faster solution was sought. */
// func lookAndSay(input string) string {
// 	// Go through each letter, and count the number of times it appears in a row. Append this to the output string.
// 	outputString := ""
// 	currentLetter := input[0]
// 	currentLetterCount := 1
// 	for i := 1; i < len(input); i++ {
// 		if input[i] == currentLetter {
// 			currentLetterCount++
// 		} else {
// 			outputString += strconv.Itoa(currentLetterCount) + string(currentLetter)
// 			currentLetter = input[i]
// 			currentLetterCount = 1
// 		}
// 	}
// 	outputString += strconv.Itoa(currentLetterCount) + string(currentLetter)
// 	return outputString
// }
// --------------------------------------

/* Another implementation, using byte slices. This was faster than the above, but not as good practice as using a builder (I think). */
// func lookAndSay(input string) string {
// 	var result []byte
// 	currentLetter := input[0]
// 	currentLetterCount := 1

// 	for i := 1; i < len(input); i++ {
// 		if input[i] == currentLetter {
// 			currentLetterCount++
// 		} else {
// 			result = appendCountAndLetter(result, currentLetterCount, currentLetter)

// 			currentLetter = input[i]
// 			currentLetterCount = 1
// 		}
// 	}

// 	result = appendCountAndLetter(result, currentLetterCount, currentLetter)

// 	return string(result)
// }
