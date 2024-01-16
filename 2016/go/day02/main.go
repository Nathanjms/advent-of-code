package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
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

var DIRECTIONS = map[rune][2]int{
	'U': {-1, 0},
	'R': {0, 1},
	'D': {1, 0},
	'L': {0, -1},
}

func partOne(contents []string) {
	COORD_TO_NUM := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}
	pin := ""
	coord := [2]int{1, 1} // Start at 5
	for _, line := range contents {
		for j, char := range line {
			newCoord := [2]int{
				coord[0] + DIRECTIONS[char][0],
				coord[1] + DIRECTIONS[char][1],
			}
			if isInRange(newCoord[0], newCoord[1], COORD_TO_NUM[:]) {
				coord = newCoord
			}
			if j == len(line)-1 {
				pin += COORD_TO_NUM[coord[0]][coord[1]]
				break
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   2,
		Part:  1,
		Value: pin,
	})
}

func partTwo(contents []string) {
	COORD_TO_NUM := [][]string{
		{"", "", "1", "", ""},
		{"", "2", "3", "4", ""},
		{"5", "6", "7", "8", "9"},
		{"", "A", "B", "C", ""},
		{"", "", "D", "", ""},
	}

	pin := ""
	coord := [2]int{2, 0} // Start at 5
	for _, line := range contents {
		for j, char := range line {
			newCoord := [2]int{
				coord[0] + DIRECTIONS[char][0],
				coord[1] + DIRECTIONS[char][1],
			}
			if isInRange(newCoord[0], newCoord[1], COORD_TO_NUM[:]) {
				coord = newCoord
			}
			if j == len(line)-1 {
				pin += COORD_TO_NUM[coord[0]][coord[1]]
				break
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   2,
		Part:  2,
		Value: pin,
	})
}

func isInRange(row int, col int, keypad [][]string) bool {
	if row >= 0 && row < len(keypad) && col >= 0 && col < len(keypad[0]) {
		return keypad[row][col] != ""
	}
	return false
}
