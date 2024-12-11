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

type cacheKey struct {
	rock            int64
	remainingBlinks int
}

func partOne(contents string) {
	rocks := parseInput(contents)

	cacheMap := make(map[cacheKey]int64, 0)

	numRocks := int64(0)
	for _, rock := range rocks {
		numRocks += getNumberOfRocks(rock, 25, &cacheMap)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  1,
		Value: numRocks,
	})
}

func partTwo(contents string) {
	rocks := parseInput(contents)

	cacheMap := make(map[cacheKey]int64, 0)

	numRocks := int64(0)
	for _, rock := range rocks {
		numRocks += getNumberOfRocks(rock, 75, &cacheMap)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  2,
		Value: numRocks,
	})
}

func getNumberOfRocks(rock int64, blinksRemaining int, cacheMap *map[cacheKey]int64) int64 {
	val, ok := (*cacheMap)[cacheKey{rock, blinksRemaining}]
	if ok {
		return val
	}

	if blinksRemaining == 0 {
		return 1 // Just this rock
	}

	if rock == 0 {
		value := getNumberOfRocks(int64(1), blinksRemaining-1, cacheMap)
		(*cacheMap)[cacheKey{rock, blinksRemaining}] = value
		return value
	}

	stringVal := strconv.FormatInt(rock, 10)
	stringValSize := len(stringVal)
	if stringValSize%2 == 0 {
		// Split down the middle;
		var newRockOneStr, newRockTwoStr string
		var newRockOne, newRockTwo int64
		// Special case if size is 2:
		if stringValSize == 2 {
			newRockOneStr = string(stringVal[0])
			newRockTwoStr = string(stringVal[1])
		} else {
			newRockOneStr = stringVal[0:(stringValSize / 2)]
			newRockTwoStr = stringVal[(stringValSize / 2):stringValSize]
		}

		newRockOne, _ = strconv.ParseInt(newRockOneStr, 10, 64)
		newRockTwo, _ = strconv.ParseInt(newRockTwoStr, 10, 64)

		value := getNumberOfRocks(newRockOne, blinksRemaining-1, cacheMap) + getNumberOfRocks(newRockTwo, blinksRemaining-1, cacheMap)

		(*cacheMap)[cacheKey{rock, blinksRemaining}] = value

		return value
	}

	value := getNumberOfRocks(rock*2024, blinksRemaining-1, cacheMap)
	(*cacheMap)[cacheKey{rock, blinksRemaining}] = value
	return value

}

func parseInput(contents string) []int64 {
	strValues := strings.Fields(contents)

	intVals := make([]int64, 0)
	for _, val := range strValues {
		intVal, _ := strconv.ParseInt(val, 10, 64)
		intVals = append(intVals, intVal)
	}
	return intVals
}
