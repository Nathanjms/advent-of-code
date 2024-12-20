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

func partOne(contents []string) {
	grid := parseToGrid(contents)
	distancesToPoint := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		distancesToPoint[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			distancesToPoint[i][j] = -1
		}
	}

	start, end := findStartAndEnd(contents)
	distancesToPoint[start[0]][start[1]] = 0
	bfsDetermineDistances(grid, start, end, &distancesToPoint)

	count := 0

	// To avoid double counting, we only look at half the points per space;
	halfDistances := [4][2]int{
		{2, 0},
		{1, 1},
		{0, 2},
		{-1, 1},
	}

	picoSecondsToSave := 100
	if isUsingExample {
		picoSecondsToSave = 40 // Expect 2 as the output
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if distancesToPoint[i][j] == -1 {
				continue
			}

			for _, dir := range halfDistances {
				newI := i + dir[0]
				newJ := j + dir[1]

				// Out of bounds checks
				if newI < 0 || newJ < 0 || newI > len(grid)-1 || newJ > len(grid[0])-1 {
					continue
				}

				if grid[newI][newJ] == '#' {
					continue
				}

				if int(math.Abs(float64(distancesToPoint[i][j]-distancesToPoint[newI][newJ]))) >= picoSecondsToSave+2 {
					count++
				}
			}

		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   20,
		Part:  1,
		Value: count,
	})
}

func partTwo(contents []string) {
	grid := parseToGrid(contents)
	distancesToPoint := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		distancesToPoint[i] = make([]int, len(grid[i]))
		for j := 0; j < len(grid[i]); j++ {
			distancesToPoint[i][j] = -1
		}
	}

	start, end := findStartAndEnd(contents)
	distancesToPoint[start[0]][start[1]] = 0
	bfsDetermineDistances(grid, start, end, &distancesToPoint)

	count := 0

	picoSecondsToSave := 100

	visited := make(map[[2][2]int]bool, 0) // start and end

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for radius := 2; radius < 21; radius++ {
				for di := 0; di < radius+1; di++ {
					dj := radius - di
					for _, dir := range [4][2]int{{i + di, j + dj}, {i + di, j - dj}, {i - di, j + dj}, {i - di, j - dj}} {
						newI := dir[0]
						newJ := dir[1]
						if _, ok := visited[[2][2]int{{i, j}, {newI, newJ}}]; ok {
							continue
						}

						visited[[2][2]int{{i, j}, {newI, newJ}}] = true

						// Out of bounds checks
						if newI < 0 || newJ < 0 || newI > len(grid)-1 || newJ > len(grid[0])-1 {
							continue
						}

						if grid[newI][newJ] == '#' {
							continue
						}

						if distancesToPoint[i][j]-distancesToPoint[newI][newJ] >= picoSecondsToSave+radius {
							count++
						}
					}
				}
			}

		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   20,
		Part:  2,
		Value: count,
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

func bfsDetermineDistances(grid [][]byte, start [2]int, end [2]int, distancesToPoint *[][]int) int {
	queue := make([]queueStruct, 0)

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

			if (*distancesToPoint)[newI][newJ] != -1 {
				continue
			}

			(*distancesToPoint)[newI][newJ] = (*distancesToPoint)[element.pos[0]][element.pos[1]] + 1

			queue = append(queue, queueStruct{
				pos:   [2]int{newI, newJ},
				steps: element.steps + 1,
			})
		}
	}

	return -1
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
