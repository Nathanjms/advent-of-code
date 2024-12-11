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
	rocks := parseInput(contents)

	numBlinks := 25

	for i := 0; i < numBlinks; i++ {
		newLayout := make([]int64, 0)

		for _, rock := range rocks {
			if rock == 0 {
				newLayout = append(newLayout, int64(1))
			} else {
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

					newLayout = append(newLayout, newRockOne)
					newLayout = append(newLayout, newRockTwo)
				} else {
					newLayout = append(newLayout, rock*2024)
				}
			}
		}

		rocks = make([]int64, len(newLayout))
		copy(rocks, newLayout)
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  1,
		Value: len(rocks),
	})
}

func partTwo(contents string) {
	rocks := parseInput(contents)

	numBlinks := 75

	// Store the output from input single digit, what it becomes
	cacheMap := make(map[int64][]int64, 400000)

	// Recursion makes sense?
	numRocks := 0
	for _, rock := range rocks {
		numRocks += getNumberOfRocks(rock, &cacheMap, 75)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  2,
		Value: len(rocks),
	})
}

func getNumberOfRocks(rock int64, cacheMap *map[int64][]int64, blinks int) int64 {
	val, ok := (*cacheMap)[rock]
	if ok {
		for _, newRock := range val {
			numRocks += getNumberOfRocks(newRock, cacheMap, numRocks)
		}

		return numRocks
	}
	// Otherwise, do it then store in cache
	if rock == 0 {
		newRock = 1
		(*cacheMap)[rock] = 1
	} else {
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

			// Add to cache;
			cacheMap[rock] = [2]int64{newRockOne, newRockTwo}
		} else {
			value := rock * 2024
			cacheMap[rock] = [2]int64{value, -1}
		}
	}
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
