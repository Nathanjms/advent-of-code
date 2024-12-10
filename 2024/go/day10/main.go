package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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
	grid := parseInput(contents)
	scoresSum := 0

	// Go through each point and if it's a 0, do a DFS on it
	for i, row := range grid {
		for j, point := range row {
			if point == 0 {
				scoresSum += determineScoreForPoint(&grid, i, j, false)
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   10,
		Part:  1,
		Value: scoresSum,
	})
}

type queueStruct struct {
	pos     [2]int
	currVal int
}

func partTwo(contents []string) {
	grid := parseInput(contents)
	scoresSum := 0

	// Go through each point and if it's a 0, do a DFS on it
	for i, row := range grid {
		for j, point := range row {
			if point == 0 {
				scoresSum += determineScoreForPoint(&grid, i, j, true)
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   10,
		Part:  2,
		Value: scoresSum,
	})
}

func determineScoreForPoint(grid *[][]int, i int, j int, isPartTwo bool) int {
	numOfNines := 0
	visited := make(map[[2]int]bool, 0)
	directions := [4][2]int{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}
	currentValue := (*grid)[i][j]

	queue := make([]queueStruct, 0)

	queue = append(queue, queueStruct{
		pos:     [2]int{i, j},
		currVal: currentValue,
	})

	for {
		if len(queue) == 0 {
			break
		}

		// Grab the last value from the queue
		element := queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		// if visited, for part 1 we exit. For part 2, we are counting the total number of ways to reach 9, so we can overlap
		if !isPartTwo {
			_, ok := visited[element.pos]
			if ok {
				continue
			}
		}

		// otherwise, mark as visited!
		visited[element.pos] = true

		if element.currVal == 9 {
			// We've made it, can increment and abort!
			numOfNines++
			continue
		}

		// Now expand out in all 4 directions and see if valid, add to queue if so
		for _, dir := range directions {
			newI := element.pos[0] + dir[0]
			newJ := element.pos[1] + dir[1]

			// Out of bounds checks
			if newI < 0 || newJ < 0 || newI > len(*grid)-1 || newJ > len((*grid)[0])-1 {
				continue
			}

			// Rules check; it must increment 1 from the current Value:
			if element.currVal != (*grid)[newI][newJ]-1 {
				continue
			}

			// Otherwise, we add to list
			queue = append(queue, queueStruct{
				pos:     [2]int{newI, newJ},
				currVal: (*grid)[newI][newJ],
			})
		}
	}

	return numOfNines
}

func parseInput(contents []string) [][]int {
	output := make([][]int, len(contents))
	for i, line := range contents {
		lineInt := make([]int, len(line))
		for j, char := range line {
			intVal, _ := strconv.Atoi(string(char))
			lineInt[j] = intVal
		}
		output[i] = lineInt
	}

	return output
}
