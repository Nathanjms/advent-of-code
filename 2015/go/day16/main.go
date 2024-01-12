package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
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
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:  16,
		Part: 1,
		Value: getCorrectSue(
			contents,
			map[string]string{
				"children":    "eq",
				"cats":        "eq",
				"samoyeds":    "eq",
				"pomeranians": "eq",
				"akitas":      "eq",
				"vizslas":     "eq",
				"goldfish":    "eq",
				"trees":       "eq",
				"cars":        "eq",
				"perfumes":    "eq",
			},
		),
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:  16,
		Part: 2,
		Value: getCorrectSue(
			contents,
			map[string]string{
				"children":    "eq",
				"cats":        "gt",
				"samoyeds":    "eq",
				"pomeranians": "lt",
				"akitas":      "eq",
				"vizslas":     "eq",
				"goldfish":    "lt",
				"trees":       "gt",
				"cars":        "eq",
				"perfumes":    "eq",
			},
		),
	})
}

var targetsByKey = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func getCorrectSue(contents []string, operationKeys map[string]string) int {
	giftSueNum := 0

	for _, line := range contents {
		// Sue 1: cars: 9, akitas: 3, goldfish: 0
		var sueNum, attOneVal, attTwoVal, attThreeVal int
		var attOne, attTwo, attThree string
		_, err := fmt.Sscanf(
			line,
			"Sue %d: %s %d, %s %d, %s %d",
			&sueNum,
			&attOne,
			&attOneVal,
			&attTwo,
			&attTwoVal,
			&attThree,
			&attThreeVal,
		)
		if err != nil {
			panic(err)
		}

		// remove the colons
		attOne = attOne[:len(attOne)-1]
		attTwo = attTwo[:len(attTwo)-1]
		attThree = attThree[:len(attThree)-1]

		// Check all three match, and break once we find a match
		if checkAttribute(attOne, attOneVal, &operationKeys) &&
			checkAttribute(attTwo, attTwoVal, &operationKeys) &&
			checkAttribute(attThree, attThreeVal, &operationKeys) {
			giftSueNum = sueNum
			break
		}
	}

	return giftSueNum
}

func checkAttribute(attrName string, attrVal int, operationKeys *map[string]string) bool {
	compareType := (*operationKeys)[attrName]

	if compareType == "gt" {
		return attrVal > targetsByKey[attrName]
	} else if compareType == "lt" {
		return attrVal < targetsByKey[attrName]
	} else {
		return attrVal == targetsByKey[attrName]
	}
}
