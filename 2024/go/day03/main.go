package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents string) {
	multiplier := 0

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllString(contents, -1)
	for _, match := range matches {
		var valOne, valTwo int
		_, err := fmt.Sscanf(match, "mul(%d,%d)", &valOne, &valTwo)
		if err != nil {
			panic(err)
		}
		fmt.Println(valOne, valTwo)
		multiplier += valOne * valTwo
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   3,
		Part:  1,
		Value: multiplier,
	})
}

func partTwo(contents string) {
	multiplier := 0
	addEnabled := true

	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|(do\(\))|(don't\(\))`)
	matches := re.FindAllString(contents, -1)
	for _, match := range matches {
		if match == "do()" {
			addEnabled = true
			continue
		} else if match == "don't()" {
			addEnabled = false
			continue
		}

		if addEnabled {
			var valOne, valTwo int
			_, err := fmt.Sscanf(match, "mul(%d,%d)", &valOne, &valTwo)
			if err != nil {
				panic(err)
			}
			fmt.Println(valOne, valTwo)
			multiplier += valOne * valTwo
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   3,
		Part:  2,
		Value: multiplier,
	})
}
