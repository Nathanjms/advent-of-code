package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"regexp"
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
	// Build 1000x1000 grid of bools
	var grid [1000][1000]bool
	// Define type of 'turn on', 'turn off', 'toggle'
	for _, line := range contents {
		var startRowIndex, startColIndex, endRowIndex, endColIndex int
		var action string
		// Parse line from eg 'turn on 489,959 through 759,964' using regex
		re := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
		var matches = re.FindStringSubmatch(line)

		action = matches[1]
		startRowIndex, _ = strconv.Atoi(matches[2])
		startColIndex, _ = strconv.Atoi(matches[3])
		endRowIndex, _ = strconv.Atoi(matches[4])
		endColIndex, _ = strconv.Atoi(matches[5])

		if action == "toggle" {
			for i := startRowIndex; i <= endRowIndex; i++ {
				for j := startColIndex; j <= endColIndex; j++ {
					grid[i][j] = !grid[i][j]
				}
			}
		} else {
			for i := startRowIndex; i <= endRowIndex; i++ {
				for j := startColIndex; j <= endColIndex; j++ {
					grid[i][j] = action == "turn on"
				}
			}
		}
	}

	// Count number of lights that are on
	var count int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] {
				count++
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   6,
		Part:  1,
		Value: count,
	})
}

func partTwo(contents []string) {
	// Build 1000x1000 grid of bools
	var grid [1000][1000]int
	// Define type of 'turn on', 'turn off', 'toggle'
	for _, line := range contents {
		var startRowIndex, startColIndex, endRowIndex, endColIndex int
		var action string
		// Parse line from eg 'turn on 489,959 through 759,964' using regex
		re := regexp.MustCompile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
		var matches = re.FindStringSubmatch(line)

		action = matches[1]
		startRowIndex, _ = strconv.Atoi(matches[2])
		startColIndex, _ = strconv.Atoi(matches[3])
		endRowIndex, _ = strconv.Atoi(matches[4])
		endColIndex, _ = strconv.Atoi(matches[5])

		if action == "toggle" {
			for i := startRowIndex; i <= endRowIndex; i++ {
				for j := startColIndex; j <= endColIndex; j++ {
					grid[i][j] += 2
				}
			}
		} else if action == "turn on" {
			for i := startRowIndex; i <= endRowIndex; i++ {
				for j := startColIndex; j <= endColIndex; j++ {
					grid[i][j]++
				}
			}
		} else if action == "turn off" {
			for i := startRowIndex; i <= endRowIndex; i++ {
				for j := startColIndex; j <= endColIndex; j++ {
					if grid[i][j] > 0 {
						grid[i][j]--
					}
				}
			}
		}
	}

	// Count number of lights that are on
	var count int
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			count += grid[i][j]
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   6,
		Part:  2,
		Value: count,
	})
}
