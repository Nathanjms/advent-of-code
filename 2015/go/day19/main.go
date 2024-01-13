package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"sort"
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
	if isUsingExample {
		inputPath += "2"
	}
	var _, contentsPt2 = sharedcode.ParseFile(inputPath)

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
		// For each found index, we need to string replace the indexes shown with the new values, then add this to the possibliltiesByKey
		for _, idxs := range indexes {
			start := idxs[0]
			end := idxs[1]
			output := input[0:start] + new + input[end:]
			possibilitiesByKey[output] = true
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   19,
		Part:  1,
		Value: len(possibilitiesByKey),
	})
}

func partTwo(contents []string) {

	mapping, input := parseInput(contents)

	// // Sort the mapping by the second element in each inner slice
	sort.Slice(mapping, func(i, j int) bool {
		return len(mapping[i][1]) > len(mapping[j][1])
	})

	// Randomness to the rescue? Super fast but should it work this well?

	validResults := make(map[int]int)
	validResultCount := 0
	inputOriginal := input
	var steps int
	var repeats int

	for {
		input = inputOriginal
		steps = 0
		repeats = 0
		for {
			// Choose a random array element to use for this loop
			mapArr := mapping[rand.Intn(len(mapping))]
			orig := mapArr[0]
			new := mapArr[1]
			if orig == "e" {
				if input == new {
					input = "e"
					steps++
					break
				}
			} else {
				inputOld := input

				// Replace any occurrences of the new variable with the old
				input = strings.Replace(input, new, orig, 1)
				if inputOld != input {
					steps++
				}
				if input == "e" {
					break
				}
				// 'Big enough' to mean we wont ever get a good result
				if repeats > 1000 {
					break
				}
				repeats++
			}
		}
		if input == "e" {
			if _, ok := validResults[steps]; !ok {
				validResults[steps] = 1
			} else {
				validResults[steps]++
			}
			validResultCount++
		}
		if validResultCount > 100 {
			break
		}
	}

	// Choose the one with the largest occurrence. In reality, it looks like it only ever returns the correct result.
	result := math.MaxInt
	for key := range validResults {
		result = min(key, result)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   19,
		Part:  2,
		Value: result,
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
