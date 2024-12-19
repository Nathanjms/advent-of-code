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

type cacheKey struct {
	towel   string
	pattern string
}

func partOne(contents []string) {
	availableTowels, patterns := parseInput(contents)

	possibleDesigns := 0

	cacheMap := make(map[cacheKey]bool, 0)

	for _, pattern := range patterns {
		if canMakeDesign(availableTowels, pattern, &cacheMap) {
			possibleDesigns++
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   19,
		Part:  1,
		Value: possibleDesigns,
	})
}

func canMakeDesign(availableTowels map[string]bool, pattern string, cacheMap *map[cacheKey]bool) bool {
	if len(pattern) == 0 {
		return true
	}

	for towel := range availableTowels {
		if val, ok := (*cacheMap)[cacheKey{pattern: pattern, towel: towel}]; ok {
			if val {
				return true
			} else {
				return false
			}
		}

		if len(towel) > len(pattern) {
			continue // This towel is too big. Maybe we could drop from future recursions if slow performance
		}

		// Does this towel work as the next slot?
		if towel == pattern[:len(towel)] {
			// This is valid, so we can recursively determine if this towel is a valid one for this step
			newPattern := pattern[len(towel):]
			isValid := canMakeDesign(availableTowels, newPattern, cacheMap)
			(*cacheMap)[cacheKey{pattern: newPattern, towel: towel}] = isValid
			if isValid {
				return true
			}
		}
	}

	return false
}

func partTwo(contents []string) {
	availableTowels, patterns := parseInput(contents)

	numDesignsForPattern := make(map[string]int, 0)

	totalWaysToMakeDesign := 0
	for _, pattern := range patterns {
		totalWaysToMakeDesign += howManyWaysToMakeDesign(availableTowels, pattern, &numDesignsForPattern)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   19,
		Part:  2,
		Value: totalWaysToMakeDesign,
	})
}

func howManyWaysToMakeDesign(availableTowels map[string]bool, pattern string, numDesignsForPattern *map[string]int) int {
	if len(pattern) == 0 {
		return 1
	}

	if val, ok := (*numDesignsForPattern)[pattern]; ok {
		return val
	}
	totalNumWays := 0
	for towel := range availableTowels {
		if len(towel) > len(pattern) {
			continue // This towel is too big. Maybe we could drop from future recursions if slow performance
		}

		// Does this towel work as the next slot?
		if towel == pattern[:len(towel)] {
			// This is valid, so we can recursively determine if this towel is a valid one for this step
			newPattern := pattern[len(towel):]
			numWays := howManyWaysToMakeDesign(availableTowels, newPattern, numDesignsForPattern)
			(*numDesignsForPattern)[newPattern] = numWays
			totalNumWays += numWays
		}
	}

	return totalNumWays
}

func parseInput(contents []string) (map[string]bool, []string) {
	availableTowels := make(map[string]bool, 0)
	patterns := make([]string, 0)
	for i, line := range contents {
		if len(line) == 0 {
			continue
		}
		if i == 0 {
			values := strings.Fields(line)
			for j, value := range values {
				if j != len(values)-1 {
					availableTowels[value[:len(value)-1]] = true // Remove comma
				} else {
					availableTowels[value] = true
				}
			}

			continue
		}
		// Parse sets of patterns here
		patterns = append(patterns, line)
	}

	return availableTowels, patterns
}
