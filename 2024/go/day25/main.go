package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
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
	locks, keys := parseInput(contents)

	uniquePairs := 0
	for _, key := range keys {
		// Try all locks;
		for _, lock := range locks {
			// Does the ey fit in the lock?
			if fitInLock(key, lock) {
				uniquePairs++
			}
		}

	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   25,
		Part:  1,
		Value: uniquePairs,
	})
}

func fitInLock(key [5]int, lock [5]int) bool {
	for i := 0; i < 5; i++ {
		// Overlap if the sum is bigger than 5
		if key[i]+lock[i] > 5 {
			return false
		}
	}

	return true
}
func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   25,
		Part:  2,
		Value: "Merry Christmas",
	})
}

func parseInput(contents []string) ([][5]int, [][5]int) {
	locks := make([][5]int, 0)
	keys := make([][5]int, 0)
	for i := 0; i < len(contents); i += 8 {
		current := [5]int{}
		if contents[i] == "#####" {
			for j := 1; j < 8; j++ {
				for n := range contents[j] {
					if contents[i+j][n] == '.' && contents[i+j-1][n] == '#' {
						current[n] = j - 1
					}
				}
			}
			locks = append(locks, current)
		} else {
			for j := 6; j >= 0; j-- {
				for n := range contents[j] {
					if contents[i+j][n] == '#' && contents[i+j-1][n] == '.' {
						current[n] = 6 - j
					}
				}
			}
			keys = append(keys, current)
		}
	}

	return locks, keys
}
