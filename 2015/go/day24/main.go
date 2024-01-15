package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strconv"

	"gonum.org/v1/gonum/stat/combin"
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
	entanglement := solve(contents, 3)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   24,
		Part:  1,
		Value: entanglement,
	})
}

func partTwo(contents []string) {
	entanglement := solve(contents, 4)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   24,
		Part:  2,
		Value: entanglement,
	})
}

func solve(contents []string, numSections int) int {
	parcelWeights := make([]int, len(contents))
	totalWeight := 0
	for _, v := range contents {
		weight, _ := strconv.Atoi(v)
		parcelWeights = append(parcelWeights, weight)
		totalWeight += weight
	}

	// Start with 1 parcel, and iterate until we get a weight equal to total/numSections. We then *should* check if the others can be split properly (Update: no need for this question as was the right answer without this!)
	numParcels := 1
	validIndexes := make([][]int, 0)
	for {
		gen := combin.NewCombinationGenerator(len(parcelWeights), numParcels)
		for gen.Next() {
			indexes := gen.Combination(nil)
			sum := 0
			for _, idx := range indexes {
				sum += parcelWeights[idx]
			}
			if sum == totalWeight/numSections {
				validIndexes = append(validIndexes, indexes)
				break
			}
		}
		if len(validIndexes) > 0 {
			break
		}
		numParcels++
	}

	entanglement := math.MaxInt
	for _, option := range validIndexes {
		tempEntanglement := 1
		for _, idx := range option {
			tempEntanglement *= parcelWeights[idx]
		}
		entanglement = min(entanglement, tempEntanglement)
	}

	return entanglement
}
