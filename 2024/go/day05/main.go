package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"reflect"
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

type Instruction [2]int
type Update []int

func partOne(contents []string) {
	instructions, updates := parseInput(contents)

	middleNumSum := 0

	for _, update := range updates {
		isValid := true
		for _, instruction := range instructions {
			firstPageNum := getIndexOfInt(update, instruction[0])
			if firstPageNum == -1 {
				// Not exists so can skip
				continue
			}

			secondPageNum := getIndexOfInt(update, instruction[1])

			if secondPageNum == -1 {
				continue
			}

			if secondPageNum < firstPageNum {
				isValid = false
				break
			}

		}

		if isValid {
			middleNumSum += update[(len(update)-1)/2] // Assume that there's no even ones...
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   5,
		Part:  1,
		Value: middleNumSum,
	})
}

func partTwo(contents []string) {
	instructions, updates := parseInput(contents)

	invalidUpdates := make([]Update, 0)

	for _, update := range updates {
		isValid := true
		for _, instruction := range instructions {
			firstPageNum := getIndexOfInt(update, instruction[0])
			if firstPageNum == -1 {
				// Not exists so can skip
				continue
			}

			secondPageNum := getIndexOfInt(update, instruction[1])

			if secondPageNum == -1 {
				continue
			}

			if secondPageNum < firstPageNum {
				isValid = false
				break
			}
		}

		if !isValid {
			invalidUpdates = append(invalidUpdates, update)
		}
	}

	middleNumSum := 0
	for _, update := range invalidUpdates {
		update := bruteForceUntilValid(update, instructions)
		sort.Slice(update, func(i, j int) bool {
			return getSortScore(update, i, instructions) < getSortScore(update, j, instructions)
		})
		// Assume all have the ability to be sorted correctly;
		middleNumSum += update[(len(update)-1)/2] // Assume that there's no even ones...
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   5,
		Part:  2,
		Value: middleNumSum,
	})
}

func parseInput(contents []string) ([]Instruction, []Update) {
	isInstruction := true
	instructions := make([]Instruction, 0)
	updates := make([]Update, 0)

	for _, line := range contents {
		if line == "" {
			isInstruction = false
			continue
		}

		if isInstruction {
			var inOne, inTwo int

			temp := strings.Split(line, "|")
			inOne, _ = strconv.Atoi(temp[0])
			inTwo, _ = strconv.Atoi(temp[1])

			instructions = append(instructions, [2]int{inOne, inTwo})
			continue
		}

		temp := strings.Split(line, ",")
		ints := make(Update, len(temp))

		for i := range temp {
			ints[i], _ = strconv.Atoi(temp[i])
		}

		updates = append(updates, ints)
	}

	return instructions, updates
}

func getIndexOfInt(array []int, val int) int {
	for i := 0; i < len(array); i++ {
		if array[i] == val {
			return i
		}
	}
	return -1
}

func getSortScore(update Update, index int, instructions []Instruction) int {
	// Find all the rules where it's on the left OR right, and for each time it's on the left, add 0, for the right, add one
	val := update[index]
	sortScore := 0
	for _, instruction := range instructions {
		if instruction[1] == val {
			sortScore++
		}
	}

	return sortScore
}

func bruteForceUntilValid(update Update, instructions []Instruction) Update {
	isValid := true

	for {
		for _, instruction := range instructions {
			firstPageNum := getIndexOfInt(update, instruction[0])
			if firstPageNum == -1 {
				// Not exists so can skip
				continue
			}

			secondPageNum := getIndexOfInt(update, instruction[1])

			if secondPageNum == -1 {
				continue
			}

			if secondPageNum < firstPageNum {
				swapF := reflect.Swapper(update)
				swapF(firstPageNum, secondPageNum)

				update = bruteForceUntilValid(update, instructions)
			}
		}

		if isValid {
			return update
		}
	}
}
