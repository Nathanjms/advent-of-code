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

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

type Calibration struct {
	result     int64
	testValues []int64
}

func partOne(contents []string) {
	calibrations := parseInput(contents)
	var totalValidResult int64

	for _, calibration := range calibrations {
		if recursiveSolve(&calibration, 0, 0, false) {
			totalValidResult += calibration.result
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   7,
		Part:  1,
		Value: totalValidResult,
	})
}

func partTwo(contents []string) {
	calibrations := parseInput(contents)
	var totalValidResult int64

	for _, calibration := range calibrations {
		if recursiveSolve(&calibration, 0, 0, true) {
			totalValidResult += calibration.result
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   7,
		Part:  2,
		Value: totalValidResult,
	})
}

func recursiveSolve(calibration *Calibration, currValue int64, index int, isPartTwo bool) bool {
	if index == len(calibration.testValues) {
		return currValue == calibration.result
	}

	// Otherwise, define the possibilities for the next index
	return recursiveSolve(calibration, currValue*calibration.testValues[index], index+1, isPartTwo) ||
		recursiveSolve(calibration, currValue+calibration.testValues[index], index+1, isPartTwo) ||
		(isPartTwo && recursiveSolve(calibration, concatenateThenInt(currValue, calibration.testValues[index]), index+1, true))

}

func parseInput(contents []string) []Calibration {
	calibrations := make([]Calibration, 0)
	for _, line := range contents {
		temp := strings.Split(line, ": ")

		result, _ := strconv.ParseInt(temp[0], 10, 64)

		strValues := strings.Fields(temp[1])
		intVals := make([]int64, 0)
		for _, val := range strValues {
			intVal, _ := strconv.ParseInt(val, 10, 64)
			intVals = append(intVals, intVal)
		}

		calibrations = append(calibrations, Calibration{
			result:     result,
			testValues: intVals,
		})
	}

	return calibrations

}

func concatenateThenInt(intOne int64, intTwo int64) int64 {
	strOne := strconv.FormatInt(intOne, 10)
	strTwo := strconv.FormatInt(intTwo, 10)

	// Now concatenate;
	res := strOne + strTwo

	// then convert and return
	intResult, err := strconv.ParseInt(res, 10, 64)
	if err != nil {
		panic(err)
	}
	return intResult
}
