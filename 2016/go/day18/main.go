package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents string) {
	rows := 40
	safeTileCount := countSafeTilesInRow(contents)
	row := contents
	if isUsingExample {
		rows = 10
	}

	for i := 1; i < rows; i++ {
		// First we build the row, then we count the safe tiles
		row = buildRowFromPreviousRow(row)

		safeTileCount += countSafeTilesInRow(row)
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   18,
		Part:  1,
		Value: safeTileCount,
	})
}

func partTwo(contents string) {
	rows := 400000
	safeTileCount := countSafeTilesInRow(contents)
	row := contents

	for i := 1; i < rows; i++ {
		// First we build the row, then we count the safe tiles
		row = buildRowFromPreviousRow(row)

		safeTileCount += countSafeTilesInRow(row)
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   18,
		Part:  2,
		Value: safeTileCount,
	})
}

func countSafeTilesInRow(row string) int {
	safeTileCount := 0
	for i := 0; i < len(row); i++ {
		if row[i] == '.' {
			safeTileCount++
		}
	}
	return safeTileCount
}

func buildRowFromPreviousRow(previousRow string) string {
	var builder strings.Builder
	for i := 0; i < len(previousRow); i++ {
		// Check the left, center, and right tiles above
		/**
		  It is a trap only if:
		  Its left and center tiles are traps, but its right tile is not.
		  Its center and right tiles are traps, but its left tile is not.
		  Only its left tile is a trap.
		  Only its right tile is a trap.
		*/

		left := '.'
		center := previousRow[i]
		right := '.'

		if i > 0 {
			left = rune(previousRow[i-1])
		}

		if i < len(previousRow)-1 {
			right = rune(previousRow[i+1])
		}

		if left == '^' && center == '^' && right == '.' ||
			left == '.' && center == '^' && right == '^' ||
			left == '^' && center == '.' && right == '.' ||
			left == '.' && center == '.' && right == '^' {
			builder.WriteRune('^')
			continue
		}

		builder.WriteRune('.')
	}

	return builder.String()
}
