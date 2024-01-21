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
	rowSize := 6
	colSize := 50
	if isUsingExample {
		rowSize = 3
		colSize = 7
	}

	grid := make([][]bool, rowSize)
	for i := range grid {
		grid[i] = make([]bool, colSize)
	}

	for _, line := range contents {
		if line[0:4] == "rect" {
			split := strings.Split(line[5:], "x")
			width, _ := strconv.Atoi(split[0])
			height, _ := strconv.Atoi(split[1])
			for i := 0; i < height; i++ {
				for j := 0; j < width; j++ {
					grid[i][j] = true
				}
			}
		} else if line[:8] == "rotate c" {
			split := strings.Split(line[16:], " by ")
			colIndex, _ := strconv.Atoi(split[0])
			amount, _ := strconv.Atoi(split[1])
			changedRowIndexes := make(map[int]bool)
			for i := 0; i < rowSize; i++ {
				grid[(i+amount)%(rowSize)][colIndex] = originalGrid[i][colIndex]
				// Set to false only if it has not already been changed
				changedRowIndexes[(i+amount)%(rowSize)] = true
				if _, ok := changedRowIndexes[i]; !ok {
					grid[i][colIndex] = false
				}
			}
		} else {
			split := strings.Split(line[13:], " by ")
			rowIndex, _ := strconv.Atoi(split[0])
			amount, _ := strconv.Atoi(split[1])
			changedColIndexes := make(map[int]bool)
			for i := 0; i < colSize; i++ {
				grid[rowIndex][(i+amount)%colSize] = originalGrid[rowIndex][i]
				// Set to false only if it has not already been changed
				changedColIndexes[(i+amount)%colSize] = true
				if _, ok := changedColIndexes[i]; !ok {
					grid[rowIndex][i] = false
				}
			}
		}
	}

	pixelCount := 0

	for _, line := range grid {
		for _, isTrue := range line {
			if isTrue {
				pixelCount++
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   8,
		Part:  1,
		Value: pixelCount,
	})
	display(&grid) // Display it here for pt 2!
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   8,
		Part:  2,
		Value: "See Above!",
	})
}

func display(grid *[][]bool) {
	for _, line := range *grid {
		ln := ""
		for _, char := range line {
			symbol := "."
			if char {
				symbol = "#"
			}
			ln = ln + symbol
		}
		fmt.Println(ln)
	}
}

func cloneGrid(grid [][]bool) [][]bool {
	// Initialize gridOriginal with the same structure as grid
	newGrid := make([][]bool, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]bool, len(grid[i]))
	}

	// Copy the values from grid to newGrid
	for i := range grid {
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
