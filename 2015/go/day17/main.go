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
	/* We could either brute force this, or maybe write a recursive function. There are not that many, so brute force is probable viable.
	 *	As a learning exercise, I'll try a recursive function... what could possibly go wrong
	 */

	containers := make([]int, len(contents))
	for i, line := range contents {
		intVal, _ := strconv.Atoi(line)
		containers[i] = intVal
	}

	targetCapacity := 150
	if isUsingExample {
		targetCapacity = 25
	}

	var computeSize func(containers []int, remainingCapacity int) int

	computeSize = func(containers []int, remainingCapacity int) int {
		if remainingCapacity == 0 {
			// We've got a match!
			return 1
		} else if remainingCapacity < 0 || len(containers) == 0 {
			// We've gone below zero, or we've got no containers left. Either way, we've not got a match with this combo
			return 0
		}
		// Handle when first container is used, and when first container is not used
		return computeSize(containers[1:], remainingCapacity-containers[0]) + computeSize(containers[1:], remainingCapacity)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   17,
		Part:  1,
		Value: computeSize(containers, targetCapacity),
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   17,
		Part:  2,
		Value: "TODO",
	})
}
