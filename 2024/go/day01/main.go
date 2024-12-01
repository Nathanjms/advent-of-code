package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	combos := parseInput(contents)

	sort.Slice(combos[0], func(i, j int) bool {
		return combos[0][i] < combos[0][j]
	})

	sort.Slice(combos[1], func(i, j int) bool {
		return combos[1][i] < combos[1][j]
	})

	difference := int64(0)

	for i := 0; i < len(combos[0]); i++ {
		difference += int64(math.Abs(float64(combos[0][i] - combos[1][i])))
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   1,
		Part:  1,
		Value: difference,
	})
}

func partTwo(contents []string) {
	combos := parseInput(contents)

	sort.Slice(combos[0], func(i, j int) bool {
		return combos[0][i] < combos[0][j]
	})

	sort.Slice(combos[1], func(i, j int) bool {
		return combos[1][i] < combos[1][j]
	})

	similarityScore := int64(0)

	for i := 0; i < len(combos[0]); i++ {
		// We go through all of the lhs list, and multiply the number by the qnty in the rhs list
		qty := int64(0)
		currVal := combos[0][i]

		for j := 0; j < len(combos[1]); j++ {
			if combos[1][j] > currVal {
				break
			}
			if combos[1][j] == currVal {
				qty++
			}
		}

		similarityScore += qty * currVal
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   1,
		Part:  2,
		Value: similarityScore,
	})
}

func parseInput(contents []string) [2][]int64 {
	leftCombos := make([]int64, 0)
	rightCombos := make([]int64, 0)

	for _, line := range contents {
		var numOne, numTwo int64

		tmp := strings.Split(line, "   ")

		numOne, _ = strconv.ParseInt(tmp[0], 10, 64)
		numTwo, _ = strconv.ParseInt(tmp[1], 10, 64)

		leftCombos = append(leftCombos, numOne)
		rightCombos = append(rightCombos, numTwo)
	}

	return [2][]int64{leftCombos, rightCombos}
}
