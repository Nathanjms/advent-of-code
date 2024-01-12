package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
var isUsingExample = true

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
		isUsingExample = false
	}

	var _, contents = sharedcode.ParseFile(inputPath)
	var _, contentsPt2 = sharedcode.ParseFile(inputPath + "2")

	partOne(contents)
	partTwo(contentsPt2)
}

func partOne(contents []string) {
	mapping, input := parseInput(contents)

	// Make this a map so we handle duplicates 'for free'
	possibilitiesByKey := make(map[string]bool)

	// We only do one step, so simply handle it out of any recursive functions etc below
	for _, mapArr := range mapping {
		orig := mapArr[0]
		new := mapArr[1]
		indexes := findAllOccurrences(input, orig)
		fmt.Println(orig, indexes)
		// For each found index, we need to string replace the indexes shown with the new values, then add this to the possibliltiesByKey
		for _, idxs := range indexes {
			start := idxs[0]
			end := idxs[1]
			output := input[0:start] + new + input[end:]
			fmt.Println(output)
			possibilitiesByKey[output] = true
		}
		// Do we get any matches in the remaining input for the orig?
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   19,
		Part:  1,
		Value: len(possibilitiesByKey),
	})
}

func partTwo(contents []string) {

	// Suspecting recursion may be needed:
	// var getPossibilities func(buildingString string, remainingInput string) int

	// // Only doing one step, but pt2 maybe will need recursion, so build it in that format
	// getPossibilities = func(buildingString string, remainingInput string) int {
	// 	// Go through the map, and for any that find a match, handle recursion
	// 	for orig, new := range mapping {
	// 		indexes := findAllOccurrences(remainingInput, orig)
	// 		// For each found index, we
	// 		//
	// 		fmt.Println(indexes)
	// 		fmt.Println(new)
	// 		// Do we get any matches in the remaining input for the orig?

	// 	}
	// 	return 0
	// }
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   19,
		Part:  2,
		Value: "TODO",
	})
}

func parseInput(contents []string) ([][]string, string) {
	// The last value will be input, the second last will be blank.

	input := contents[len(contents)-1]
	contentsTmp := contents[0 : len(contents)-2]
	var mapping [][]string

	for _, line := range contentsTmp {
		arr := strings.Split(line, " => ")
		mapping = append(mapping, arr)
	}

	return mapping, input
}

func findAllOccurrences(s, substring string) [][]int {
	r := regexp.MustCompile(substring)
	return r.FindAllStringIndex(s, -1)
}
