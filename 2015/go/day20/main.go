package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents string) {
	targetValue, _ := strconv.Atoi(contents)
	lowestHouse := math.MaxInt

	// Let's try using brute force, but first make two changes:
	// 1. we can simply divide the score and awarded points by 10
	// 2. we can skip one and add 1 to the score

	targetValue = targetValue / 10

	houseValues := make([]int, targetValue)
	for i := 2; i <= targetValue; i++ {
		// Now go through and get all the houses that will be visited by this num, adding the score to their value
		for j := i; j <= targetValue; j += i {
			houseValues[j-1] += i
		}
	}

	for house, val := range houseValues {
		// Subtract one to account for skipping the +1 on all houses above
		if val >= targetValue-1 {
			lowestHouse = house + 1 // add one to account for 0th index
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   20,
		Part:  1,
		Value: lowestHouse,
	})
}

func partTwo(contents string) {
	targetValue, _ := strconv.Atoi(contents)
	lowestHouse := math.MaxInt

	// Essentially lose the assumptions this time, but we gain a bit of efficiency from knowing that we stop after 50 visits.
	// Lets try and start it by still only counting to the target/10, because it probably will be below this value. If it's wrong, we'll
	// accept defeat and do a super large array

	houseValues := make([]int, targetValue/10)
	for i := 1; i <= targetValue/10; i++ {
		visits := 0 // Track the number of visits for this elf
		// Now go through and get all the houses that will be visited by this num, adding the score to their value
		for j := i; j <= targetValue/10; j += i {
			if visits >= 50 {
				// Elves stop after their first 50 visits, so account for this
				break
			}
			houseValues[j-1] += i * 11
			visits++
		}

	}

	for house, val := range houseValues {
		// Subtract one to account for skipping the +1 on all houses above
		if val >= targetValue {
			lowestHouse = house + 1 // add one to account for 0th index
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   20,
		Part:  2,
		Value: lowestHouse,
	})
}
