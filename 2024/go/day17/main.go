package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
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
	registers, instructions := parseInput(contents)

	fmt.Println(registers, instructions)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   17,
		Part:  1,
		Value: "TODO",
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   17,
		Part:  2,
		Value: "TODO",
	})
}

func parseInput(contents []string) (map[byte]int64, []int64) {
	registers := make(map[byte]int64)
	instructions := make([]int64, 0)

	for _, line := range contents {
		if len(line) == 0 {
			continue
		}

		if line[:8] == "Register" {
			register := line[10]
			val := line[12:]

			intVal, _ := strconv.ParseInt(val, 10, 64)

			registers[register] = intVal
		}

		if line[:7] == "Program" {
			stringInstructions := strings.Split(line, ",")
			for _, val := range stringInstructions {
				intVal, _ := strconv.ParseInt(val, 10, 64)
				instructions = append(instructions, intVal)
			}
		}
	}

	return registers, instructions
}
