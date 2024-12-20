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
	grid := parseToGrid(contents)
	start, end := findStartAndEnd(contents)
	numSeconds := bfs(grid, start, end)
	savedSecondsCount := make(map[int]int, 0)

	displayGrid(grid)

	// We'll go through all options of removals. If they can remove just 1, or 1&2 or 2, is this 3 options??
	// We might need to avoid wasting time and duplicating, 1,2 != 2,1?? Will try without first
	for i := 1; i < len(grid)-2; i++ {
		for j := 1; j < len(grid[i])-1; j++ {
			nextMustBeHash := false
			if grid[i][j] == 'E' || grid[i][j] == 'S' {
				continue
			}
			if grid[i][j] != '#' {
				continue
				// // Go 4 directions and only valid if it is a #
				// nextMustBeHash = true
			}

			for _, dir := range directions {
				newI := i + dir[0]
				newJ := j + dir[1]

				// Out of bounds checks
				if newI < 0 || newJ < 0 || newI > len(grid)-1 || newJ > len(grid[i])-1 {
					continue
				}

				// Can't finish on an out of bounds tile!
				if grid[newI][newJ] == '#' {
					continue
				}

				if nextMustBeHash && grid[newI][newJ] != '#' {
					// Would not be a cheat, so skip!
					continue
				}

				newGrid := cloneGrid(grid)
				newGrid[i][j] = '.'
				if grid[newI][newJ] != 'E' && grid[newI][newJ] != 'S' {
					newGrid[newI][newJ] = '.'
				}

				seconds := bfs(newGrid, start, end)
				// Took 100s, now takes 55s. So saved 100-55=45 seconds
				savedSeconds := numSeconds - seconds
				if savedSeconds > 0 {
					if savedSeconds == 64 {
						fmt.Println()
						displayGrid(newGrid)
						fmt.Println("debug")
					}
					savedSecondsCount[savedSeconds]++
				}
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   20,
		Part:  1,
		Value: savedSecondsCount,
	})
}

var directions = [4][2]int{
	{1, 0},  // Down
	{0, 1},  // Right
	{-1, 0}, // Up
	{0, -1}, // Left
}

type queueStruct struct {
	pos   [2]int
	steps int
}

func bfs(grid [][]byte, start [2]int, end [2]int) int {
	// BFS now!
	queue := make([]queueStruct, 0)
	visited := make(map[[2]int]bool)

	queue = append(queue, queueStruct{
		pos:   start,
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

		if element.pos == end {
			return element.steps
		}

		visited[element.pos] = true

		for _, dir := range directions {
			newI := element.pos[0] + dir[0]
			newJ := element.pos[1] + dir[1]

			// Out of bounds checks
			if newI < 0 || newJ < 0 || newI > len(grid)-1 || newJ > len(grid[0])-1 {
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

	return -1
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   20,
		Part:  2,
		Value: "TODO",
	})
}

func findStartAndEnd(contents []string) ([2]int, [2]int) {
	var start, end [2]int
	foundCount := 0
	for i, line := range contents {
		for j, char := range line {
			if char == 'S' {
				start = [2]int{i, j}
				foundCount++
			}
			if char == 'E' {
				end = [2]int{i, j}
				foundCount++
			}
		}
		if foundCount == 2 {
			break
		}
	}

	return start, end
}

func parseToGrid(contents []string) [][]byte {
	grid := make([][]byte, len(contents))
	for i := 0; i < len(contents); i++ {
		grid[i] = make([]byte, len(contents[i]))
		for j := 0; j < len(contents[i]); j++ {
			grid[i][j] = contents[i][j]
		}
	}

	return grid
}

func displayGrid(grid [][]byte) {
	for _, row := range grid {
		for _, char := range row {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}

func cloneGrid(grid [][]byte) [][]byte {
	newGrid := make([][]byte, len(grid))
	for i := range grid {
		newGrid[i] = make([]byte, len(grid[i]))
		copy(newGrid[i], grid[i])
	}
	return newGrid
}
