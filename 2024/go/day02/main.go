package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
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

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	reports := parseInput(contents)

	safeReports := make([][]int, 0)

	for _, line := range reports {
		if isSafe, _ := passes(line); isSafe {
			safeReports = append(safeReports, line)
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   2,
		Part:  1,
		Value: len(safeReports),
	})
}

func partTwo(contents []string) {
	reports := parseInput(contents)

	safeReports := make([][]int, 0)
	// reportIndexesToCheck := make([]int, 0)

	for _, line := range reports {
		// Check if pass, then check where it failed (and also zero);
		// isSafe, failureIndex := passes(line)
		// if isSafe {
		// 	safeReports = append(safeReports, line)
		// } else if isNowSafe, _ := passes(removeAtIndex(line, failureIndex)); isNowSafe {
		// 	safeReports = append(safeReports, line)
		// } else if isNowSafe, _ := passes(removeAtIndex(line, 0)); isNowSafe {
		// 	safeReports = append(safeReports, line)
		// }

		// Brute force each one as checking just the failure doesn't seem to work (must be some edge cases)
		if isSafe, _ := passes(line); isSafe {
			safeReports = append(safeReports, line)
		} else {
			for i := 0; i < len(line); i++ {
				newReport := removeAtIndex(line, i)
				if isNowSafe, _ := passes(newReport); isNowSafe {
					safeReports = append(safeReports, line)
					break
				}
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   2,
		Part:  2,
		Value: len(safeReports),
	})
}

func passes(report []int) (bool, int) {
	prevDiff := 0
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// If diff > 3, then we skip this line
		if math.Abs(float64(diff)) > 3 ||
			math.Abs(float64(diff)) < 1 ||
			diff == 0 ||
			(prevDiff < 0 && diff > 0) ||
			(prevDiff > 0 && diff < 0) {
			return false, i
		}

		prevDiff = diff
	}
	return true, 0
}

func parseInput(contents []string) [][]int {
	output := make([][]int, 0)

	for _, line := range contents {
		tmp := strings.Fields(line)
		lineOutput := []int{}

		for i := 0; i < len(tmp); i++ {
			tmp, err := strconv.Atoi(tmp[i])
			if err != nil {
				panic(err)
			}
			lineOutput = append(lineOutput, tmp)
		}
		output = append(output, lineOutput)
	}

	return output
}

func removeAtIndex(s []int, index int) []int {
	clone := make([]int, len(s))
	copy(clone, s)
	return append(clone[:index], clone[index+1:]...)
}
