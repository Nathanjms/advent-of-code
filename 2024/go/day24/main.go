package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
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

	_, contents := sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	valuesMap, instructions := parseInput(contents)
	counter := 0

	for {
		instruction := instructions[counter]

		// If both inputs exist in the map, we can run!
		if valuesMap[instruction.inputOne] == -1 || valuesMap[instruction.inputTwo] == -1 {
			counter = (counter + 1) % len(instructions)
			continue
		}

		if instruction.operator == "OR" {
			if valuesMap[instruction.inputOne] == 1 || valuesMap[instruction.inputTwo] == 1 {
				valuesMap[instruction.outputKey] = 1
			} else {
				valuesMap[instruction.outputKey] = 0
			}
		} else if instruction.operator == "AND" {
			if valuesMap[instruction.inputOne] == 1 && valuesMap[instruction.inputTwo] == 1 {
				valuesMap[instruction.outputKey] = 1
			} else {
				valuesMap[instruction.outputKey] = 0
			}
		} else {
			valuesMap[instruction.outputKey] = valuesMap[instruction.inputOne] ^ valuesMap[instruction.inputTwo]
		}

		if haveAllProcessed(&valuesMap) {
			break
		}

		counter = (counter + 1) % len(instructions)
	}

	// Step 1: Filter the keys that start with 'z'
	zKeys := make([]string, 0)
	for key := range valuesMap {
		if strings.HasPrefix(key, "z") {
			zKeys = append(zKeys, key)
		}
	}

	sort.Strings(zKeys) // Sort the filtered keys

	var binaryValues strings.Builder
	for i := len(zKeys) - 1; i >= 0; i-- {
		key := zKeys[i]
		newString := strconv.FormatInt(int64(valuesMap[key]), 10)
		fmt.Println(key, valuesMap[key], newString)
		binaryValues.WriteString(newString)
	}

	fmt.Println(binaryValues.String())

	intValue, err := strconv.ParseInt(binaryValues.String(), 2, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   24,
		Part:  1,
		Value: intValue,
	})
}

func haveAllProcessed(valuesMap *map[string]int) bool {
	for _, val := range *valuesMap {
		if val == -1 {
			return false
		}
	}
	return true
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   24,
		Part:  2,
		Value: "TODO",
	})
}

type instructionStruct struct {
	inputOne  string
	operator  string
	inputTwo  string
	outputKey string
}

func parseInput(contents []string) (map[string]int, []instructionStruct) {
	valuesMap := make(map[string]int, 0)
	instructions := make([]instructionStruct, 0)
	for _, line := range contents {
		if len(line) == 0 {
			continue
		}
		if line[3] == ':' {
			value := line[len(line)-1]
			intVal, err := strconv.Atoi(string(value))
			if err != nil {
				panic(err)
			}
			valuesMap[line[0:3]] = intVal
			continue
		}

		// Else it's an instructions map:
		values := strings.Fields(line)
		instructions = append(instructions, instructionStruct{
			inputOne:  values[0],
			operator:  values[1],
			inputTwo:  values[2],
			outputKey: values[4],
		})

		// Also add to the values map as -1 for the 3 keys;
		if _, ok := valuesMap[values[0]]; !ok {
			valuesMap[values[0]] = -1 // Default to -1
		}

		if _, ok := valuesMap[values[2]]; !ok {
			valuesMap[values[2]] = -1 // Default to -1
		}

		if _, ok := valuesMap[values[4]]; !ok {
			valuesMap[values[4]] = -1 // Default to -1
		}
	}

	return valuesMap, instructions
}
