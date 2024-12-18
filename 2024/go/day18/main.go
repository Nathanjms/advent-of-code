package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"math"
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

type queueStruct struct {
	pos   [2]int
	steps int
}

func partOne(contents []string) {
	instructions := parseInput(contents)

	nanoSeconds := 1024
	gridSize := 71
	if isUsingExample {
		nanoSeconds = 12
		gridSize = 7
	}

	grid := buildGrid(instructions, nanoSeconds, gridSize)

	// displayGrid(grid)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   18,
		Part:  1,
		Value: bfs(instructions, grid),
	})
}

func bfs(instructions [][2]int, grid [][]byte) int {
	// BFS now!
	queue := make([]queueStruct, 0)
	visited := make(map[[2]int]bool)

	directions := [4][2]int{
		{1, 0},  // Down
		{0, 1},  // Right
		{-1, 0}, // Up
		{0, -1}, // Left
	}

	queue = append(queue, queueStruct{
		pos:   [2]int{0, 0},
		steps: 0,
	})

	var element queueStruct

	for {
		if len(queue) == 0 {
			break
		}

		// Grab the next element in queue
		element, queue = queue[0], queue[1:]

		// if visited, skip
		_, ok := visited[element.pos]
		if ok {
			continue
		}

		if element.pos == [2]int{len(grid) - 1, len(grid) - 1} {
			return element.steps
		}

		visited[element.pos] = true

		for _, dir := range directions {
			newI := element.pos[0] + dir[0]
			newJ := element.pos[1] + dir[1]

			// Out of bounds checks
			if newI < 0 || newJ < 0 || newI > len(grid)-1 || newJ > len(grid)-1 {
				continue
			}

			if grid[newI][newJ] == '#' {
				continue
			}

			queue = append(queue, queueStruct{
				pos:   [2]int{newI, newJ},
				steps: element.steps + 1,
			})
		}
	}

	return math.MaxInt
}

func buildGrid(instructions [][2]int, nanoseconds int, gridSize int) [][]byte {
	grid := make([][]byte, 0)
	for i := 0; i < gridSize; i++ {
		grid = append(grid, make([]byte, 0))
		for j := 0; j < gridSize; j++ {
			grid[i] = append(grid[i], '.')
		}
	}

	for t := 0; t < nanoseconds; t++ {
		instruction := instructions[t]
		grid[instruction[0]][instruction[1]] = '#'
	}

	return grid
}

func partTwo(contents []string) {
	instructions := parseInput(contents)

	nanoSeconds := 1024
	gridSize := 71
	if isUsingExample {
		nanoSeconds = 12
		gridSize = 7
	}

	grid := buildGrid(instructions, nanoSeconds, gridSize)

	nsWhenBreaks := math.MaxInt

	for i := nanoSeconds + 1; i < len(instructions)-1; i++ {
		grid[instructions[i][0]][instructions[i][1]] = '#' // Block the new coordinate spot!
		if steps := bfs(instructions, grid); steps == math.MaxInt {
			nsWhenBreaks = i
			break
		}
	}

	coordinate := fmt.Sprintf("%d,%d", instructions[nsWhenBreaks][1], instructions[nsWhenBreaks][0])
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   18,
		Part:  2,
		Value: coordinate,
	})
}

func parseInput(contents []string) [][2]int {
	coordinates := make([][2]int, 0)
	for _, line := range contents {
		var rowCoord, colCoord int
		fmt.Sscanf(line, "%d,%d", &colCoord, &rowCoord)
		coordinates = append(coordinates, [2]int{rowCoord, colCoord})
	}

	return coordinates
}

func displayGrid(grid [][]byte) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}
}
