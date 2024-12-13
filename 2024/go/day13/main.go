package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
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
	equations := parseInput(contents)

	validEquationButtonPresses := make([][2]int64, 0)

	// Now we have equations, use cramer's rule to determine X and Y (aka A btn presses and B btn presses)
	for _, eqn := range equations {
		x := ((eqn.cOne * eqn.bTwo) - (eqn.bOne - eqn.cTwo)) / ((eqn.aOne * eqn.bTwo) - (eqn.bOne * eqn.aTwo))
		y := ((eqn.aOne * eqn.cTwo) - (eqn.cOne - eqn.aTwo)) / ((eqn.aOne * eqn.bTwo) - (eqn.bOne * eqn.aTwo))

		fmt.Println(x, y)

		if x < 100 && y < 100 {
			validEquationButtonPresses = append(validEquationButtonPresses, [2]int64{x, y})
		}
	}

	fmt.Println(validEquationButtonPresses)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   13,
		Part:  1,
		Value: "TODO",
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   13,
		Part:  2,
		Value: "TODO",
	})
}

type equationsStruct struct {
	aOne int64
	bOne int64
	cOne int64
	aTwo int64
	bTwo int64
	cTwo int64
}

func parseInput(contents []string) []equationsStruct {
	equations := make([]equationsStruct, 0)
	currentIndex := 0

	for _, line := range contents {
		if len(line) == 0 {
			currentIndex++
		} else if line[:6] == "Button" {
			if len(equations) == currentIndex {
				// First time at this index so increment slice
				equations = append(equations, equationsStruct{})
			}

			entries := regexp.MustCompile(`^Button (\w): X\+(\d+), Y\+(\d+)`)
			match := entries.FindStringSubmatch(line)

			a, _ := strconv.ParseInt(match[2], 10, 64)
			b, _ := strconv.ParseInt(match[3], 10, 64)

			if match[1] == "A" {
				equations[currentIndex].aOne = a
				equations[currentIndex].bOne = b
			} else {
				equations[currentIndex].aTwo = a
				equations[currentIndex].bTwo = b
			}

		} else if line[:5] == "Prize" {
			var prizeX, prizeY int64
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &prizeX, &prizeY)

			equations[currentIndex].cOne = prizeX
			equations[currentIndex].cTwo = prizeY

			fmt.Println(equations[currentIndex])
		}
	}

	return equations
}
