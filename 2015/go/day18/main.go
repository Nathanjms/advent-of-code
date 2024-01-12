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
	sizeOfGrid := 100
	if isUsingExample {
		sizeOfGrid = 6
	}
	var grid [][]bool

	// Create a slice of slices to represent the 2D grid
	for i := 0; i < sizeOfGrid; i++ {
		grid = append(grid, make([]bool, sizeOfGrid))
	}

	for i, line := range contents {
		for j, char := range line {
			if char == '#' {
				grid[i][j] = true
			} else {
				grid[i][j] = false
			}
		}
	}

	numSteps := 100
	if isUsingExample {
		numSteps = 4
	}

	for i := 0; i < numSteps; i++ {
		// For each step, go through each piece and handle the logic
		updateGrid(&grid, false)
	}

	// Get the count of lights that are now on
	onCount := 0
	for _, line := range grid {
		for _, isOn := range line {
			if isOn {
				onCount++
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   18,
		Part:  1,
		Value: onCount,
	})
}

func partTwo(contents []string) {
	sizeOfGrid := 100
	if isUsingExample {
		sizeOfGrid = 6
	}
	var grid [][]bool

	// Create a slice of slices to represent the 2D grid
	for i := 0; i < sizeOfGrid; i++ {
		grid = append(grid, make([]bool, sizeOfGrid))
	}

	for i, line := range contents {
		for j, char := range line {
			if char == '#' {
				grid[i][j] = true
			} else {
				grid[i][j] = false
			}
		}
	}

	numSteps := 100
	if isUsingExample {
		numSteps = 4
	}

	for i := 0; i < numSteps; i++ {
		// For each step, go through each piece and handle the logic
		updateGrid(&grid, true)
	}

	// Get the count of lights that are now on
	onCount := 0
	for _, line := range grid {
		for _, isOn := range line {
			if isOn {
				onCount++
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   18,
		Part:  2,
		Value: onCount,
	})
}

func updateGrid(grid *[][]bool, cornersAlwaysOn bool) {
	// We want to change grid each time, but need to have the original for the entire set. So deep clone it
	originalGrid := cloneGrid(*grid)
	for i, line := range *grid {
		for j := range line {
			handleUpdate(grid, i, j, &originalGrid, cornersAlwaysOn)
		}
	}
}

func handleUpdate(grid *[][]bool, i int, j int, originalGrid *[][]bool, cornersAlwaysOn bool) {
	// If a corner, always leave on:
	if cornersAlwaysOn && (i == 0 || i == len(*grid)-1) && (j == 0 || j == len(*grid)-1) {
		return
	}
	// First determine if the current value is on (true) or false (off)
	if (*originalGrid)[i][j] {
		// On
		neighbourOnCount := getOnNeighboursCount(originalGrid, i, j)
		// Stay on if 2 OR 3 neighbours are on
		if neighbourOnCount != 2 && neighbourOnCount != 3 {
			(*grid)[i][j] = false
		}
	} else {
		// Off
		neighbourOnCount := getOnNeighboursCount(originalGrid, i, j)
		// Turn on if 3 neighbours are on
		if neighbourOnCount == 3 {
			(*grid)[i][j] = true
		}
	}
}

func getOnNeighboursCount(grid *[][]bool, i int, j int) int {
	neighbourOnCount := 0
	for iIdx := i - 1; iIdx <= i+1; iIdx++ {
		for jIdx := j - 1; jIdx <= j+1; jIdx++ {
			// Skip if out of bounds OR if the centre piece
			if (iIdx == i && jIdx == j) || (iIdx < 0 || iIdx > len(*grid)-1) || (jIdx < 0 || jIdx > len((*grid)[i])-1) {
				continue
			}
			if (*grid)[iIdx][jIdx] {
				neighbourOnCount++
			}
		}
	}
	return neighbourOnCount
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
