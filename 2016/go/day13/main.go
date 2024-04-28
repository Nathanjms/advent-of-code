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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

type state struct {
	coord [2]int // x, y
	steps int
}

var visited = make(map[[2]int]bool)
var isWallMap = make(map[[2]int]bool)

func partOne(contents string) {
	favouriteNumber, _ := strconv.Atoi(contents)

	steps := solve(state{
		coord: [2]int{1, 1},
		steps: 0,
	}, favouriteNumber)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   13,
		Part:  1,
		Value: steps,
	})
}

var visitedPt2 = make(map[[2]int]bool)

func partTwo(contents string) {
	favouriteNumber, _ := strconv.Atoi(contents)

	solvePt2(state{
		coord: [2]int{1, 1},
		steps: 0,
	}, favouriteNumber)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   13,
		Part:  2,
		Value: len(visitedPt2),
	})
}

func solve(initial state, favouriteNumber int) int {
	queue := []state{initial}
	iterations := 0

	for len(queue) > 0 {
		iterations++
		current := queue[0]
		queue = queue[1:]

		if isFinished(current) {
			return current.steps
		}

		if visited[current.coord] {
			continue // Skip if we've already visited this state
		}

		visited[current.coord] = true

		nextStates := generateNextStates(current, favouriteNumber)
		queue = append(queue, nextStates...)
	}

	return -1 // No solution found if we've made it here
}

func solvePt2(initial state, favouriteNumber int) int {
	queue := []state{initial}
	iterations := 0

	for len(queue) > 0 {
		iterations++
		current := queue[0]
		queue = queue[1:]

		if current.steps > 50 {
			continue
		}

		if visitedPt2[current.coord] {
			continue
		}

		// We no longer skip if already visited (i think?)

		visitedPt2[current.coord] = true

		nextStates := generateNextStates(current, favouriteNumber)
		queue = append(queue, nextStates...)
	}

	return -1
}

func isFinished(current state) bool {
	if isUsingExample {
		return current.coord[0] == 7 && current.coord[1] == 4
	}

	return current.coord[0] == 31 && current.coord[1] == 39
}

func generateNextStates(current state, favouriteNumber int) []state {
	nextStates := make([]state, 0)
	directions := [4][2]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

	for _, dir := range directions {
		newCoord := [2]int{current.coord[0] + dir[0], current.coord[1] + dir[1]}

		if newCoord[0] < 0 || newCoord[1] < 0 {
			// Can't go below 0
			continue
		}

		if isWall(newCoord[0], newCoord[1], favouriteNumber) {
			// Can't go into a wall
			continue
		}

		nextStates = append(nextStates, state{
			coord: newCoord,
			steps: current.steps + 1,
		})
	}

	return nextStates
}

/*
Find x*x + 3*x + 2*x*y + y + y*y.
Add the office designer's favorite number (your puzzle input).
Find the binary representation of that sum; count the number of bits that are 1.

	If the number of bits that are 1 is even, it's an open space.
	If the number of bits that are 1 is odd, it's a wall.
*/
func isWall(x, y, designerNumber int) bool {
	// First check the cache for whether it exists
	if isWallCoord, ok := isWallMap[[2]int{x, y}]; ok {
		return isWallCoord
	}

	result := x*x + 3*x + 2*x*y + y + y*y + designerNumber
	bitCount := countBits(result)

	isWallCoord := bitCount%2 == 1

	// Update the cache
	isWallMap[[2]int{x, y}] = isWallCoord
	return isWallCoord
}

func countBits(num int) int {
	binaryStr := strconv.FormatInt(int64(num), 2)
	count := 0
	for _, bit := range binaryStr {
		if bit == '1' {
			count++
		}
	}
	return count
}
