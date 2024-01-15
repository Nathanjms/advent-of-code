package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
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

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	variablesValueMap := make(map[byte]int)
	handleInstructions(contents, &variablesValueMap)

	targetVariable := 'b'
	if isUsingExample {
		targetVariable = 'a'
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   23,
		Part:  1,
		Value: variablesValueMap[byte(targetVariable)],
	})
}

func partTwo(contents []string) {
	variablesValueMap := make(map[byte]int)
	variablesValueMap[byte('a')] = 1 // Start a at 1 instead of 0

	handleInstructions(contents, &variablesValueMap)

	targetVariable := 'b'
	if isUsingExample {
		targetVariable = 'a'
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   23,
		Part:  2,
		Value: variablesValueMap[byte(targetVariable)],
	})
}

func handleInstructions(contents []string, variablesValueMap *map[byte]int) {
	currentIndex := 0

	for {
		line := contents[currentIndex]
		switch line[0:3] {
		case "inc":
			if _, ok := (*variablesValueMap)[line[4]]; !ok {
				(*variablesValueMap)[line[4]] = 1
			} else {
				(*variablesValueMap)[line[4]]++
			}
			currentIndex++
		case "tpl":
			if _, ok := (*variablesValueMap)[line[4]]; !ok {
				(*variablesValueMap)[line[4]] = 0
			} else {
				(*variablesValueMap)[line[4]] *= 3 // TODO: Check this works?
			}
			currentIndex++
		case "hlf":
			if _, ok := (*variablesValueMap)[line[4]]; !ok {
				(*variablesValueMap)[line[4]] = 0
			} else {
				(*variablesValueMap)[line[4]] /= 2 // TODO: Check this works?
			}
			currentIndex++
		case "jmp":
			amount, _ := strconv.Atoi(line[5:])
			if line[4] == '+' {
				currentIndex += amount
			} else {
				currentIndex -= amount
			}
		case "jie":
			variable := line[4]
			// Probably need existence check here but works for my input without sooo... nah
			if (*variablesValueMap)[variable]%2 == 0 {
				amount, _ := strconv.Atoi(line[8:])
				if line[7] == '+' {
					currentIndex += amount
				} else {
					currentIndex -= amount
				}
			} else {
				currentIndex++
			}
		case "jio":
			variable := line[4]
			// Probably need existence check here but works for my input without sooo... nah
			if (*variablesValueMap)[variable] == 1 {
				amount, _ := strconv.Atoi(line[8:])
				if line[7] == '+' {
					currentIndex += amount
				} else {
					currentIndex -= amount
				}
			} else {
				currentIndex++
			}
		}
		if currentIndex > len(contents)-1 || currentIndex < 0 {
			break
		}
	}
}
