package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"bytes"
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

var DIRECTION_MAPPINGS = map[byte][2]int{
	'^': {-1, 0},
	'v': {1, 0},
	'>': {0, 1},
	'<': {0, -1},
}

func partOne(contents []string) {
	grid, instructions := parseInput(contents, false)

	found := false

	// Get start position;
	startCoordinate := [2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' {
				startCoordinate = [2]int{i, j}
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	coordinate := startCoordinate

	// Now, go through the instructions;
	for _, instruction := range instructions {
		coordinate = moveRobotByInstruction(coordinate, instruction, &grid)
		// fmt.Println(string(instruction))
		// displayGrid(&grid)
	}

	// Now we compute the sum from the rock positions
	gpsSum := 0
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == 'O' {
				gpsSum += 100*i + j
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   15,
		Part:  1,
		Value: gpsSum,
	})
}

func partTwo(contents []string) {
	grid, instructions := parseInput(contents, true)

	found := false

	// Get start position;
	startCoordinate := [2]int{}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '@' {
				startCoordinate = [2]int{i, j}
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	coordinate := startCoordinate

	// Now, go through the instructions;
	for _, instruction := range instructions {
		coordinate = moveRobotByInstruction(coordinate, instruction, &grid)
		// fmt.Println(string(instruction))
		// displayGrid(&grid)
	}

	// Now we compute the sum from the rock positions
	gpsSum := 0
	for i, line := range grid {
		for j := range line {
			if grid[i][j] == '[' {
				gpsSum += 100*i + j
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   15,
		Part:  2,
		Value: gpsSum,
	})
}

func moveRobotByInstruction(currentPosition [2]int, direction byte, grid *[][]byte) [2]int {

	movementMappings := make(map[[2]int]moveMapItem, 0)
	if canMoveDirection(currentPosition, direction, grid, &movementMappings) {
		newPosition := [2]int{
			currentPosition[0] + DIRECTION_MAPPINGS[direction][0],
			currentPosition[1] + DIRECTION_MAPPINGS[direction][1],
		}

		// First set them all to empty, and then fill them all in. Might be overkill, but it's not a huge operation so it'll do!
		for _, mapItem := range movementMappings {
			(*grid)[mapItem.oldPosition[0]][mapItem.oldPosition[1]] = '.'
		}

		for newPos, mapItem := range movementMappings {
			(*grid)[newPos[0]][newPos[1]] = mapItem.newChar
		}

		return newPosition
	}

	return currentPosition
}

func canMoveDirection(currentPosition [2]int, direction byte, grid *[][]byte, movementMappings *map[[2]int]moveMapItem) bool {
	newPosition := [2]int{
		currentPosition[0] + DIRECTION_MAPPINGS[direction][0],
		currentPosition[1] + DIRECTION_MAPPINGS[direction][1],
	}

	if (*grid)[newPosition[0]][newPosition[1]] == '#' {
		// I think this covers the 'out of bounds' case and so we do not need 2 if statements c:
		return false
	}

	(*movementMappings)[newPosition] = moveMapItem{
		newChar:     (*grid)[currentPosition[0]][currentPosition[1]],
		oldPosition: currentPosition,
	}

	// Part 1 only if statement!
	if (*grid)[newPosition[0]][newPosition[1]] == 'O' {
		// It's a boulder, so maybe we can move it?
		return canMoveDirection(newPosition, direction, grid, movementMappings)
	}

	// Part 2 only if statement!
	if (*grid)[newPosition[0]][newPosition[1]] == '[' {
		// It's a box, so maybe we can move it, need to now check 2 positions (unless we're going right)
		if direction == '>' {
			newLeftBracketPosition := [2]int{
				newPosition[0] + DIRECTION_MAPPINGS[direction][0],
				newPosition[1] + DIRECTION_MAPPINGS[direction][1],
			}
			(*movementMappings)[newLeftBracketPosition] = moveMapItem{
				newChar:     '[',
				oldPosition: [2]int{currentPosition[0], currentPosition[1] + 1},
			}
			return canMoveDirection([2]int{newPosition[0], newPosition[1] + 1}, direction, grid, movementMappings)
		}
		return canMoveDirection(newPosition, direction, grid, movementMappings) &&
			canMoveDirection([2]int{newPosition[0], newPosition[1] + 1}, direction, grid, movementMappings)
	}

	// Part 2 only if statement!
	if (*grid)[newPosition[0]][newPosition[1]] == ']' {
		// It's a box, so maybe we can move it, need to now check 2 positions (unless we're going left)
		if direction == '<' {
			newRightBracketPosition := [2]int{
				newPosition[0] + DIRECTION_MAPPINGS[direction][0],
				newPosition[1] + DIRECTION_MAPPINGS[direction][1],
			}
			(*movementMappings)[newRightBracketPosition] = moveMapItem{
				newChar:     ']',
				oldPosition: [2]int{currentPosition[0], currentPosition[1] - 1},
			}
			return canMoveDirection([2]int{newPosition[0], newPosition[1] - 1}, direction, grid, movementMappings)
		}
		return canMoveDirection(newPosition, direction, grid, movementMappings) &&
			canMoveDirection([2]int{newPosition[0], newPosition[1] - 1}, direction, grid, movementMappings)
	}

	// Only other case is it's safe to move!
	return true
}

type moveMapItem struct {
	newChar     byte
	oldPosition [2]int
}

func parseInput(contents []string, isPartTwo bool) ([][]byte, []byte) {
	grid := make([][]byte, 0)
	var builder bytes.Buffer

	isBuildingGrid := true // first par of input is the grid

	for _, line := range contents {
		if isBuildingGrid {
			if len(line) == 0 {
				isBuildingGrid = false
				continue
			}

			if isPartTwo {
				var gridBuilder bytes.Buffer

				for i := 0; i < len(line); i++ {
					if line[i] == '.' || line[i] == '#' {
						gridBuilder.WriteByte(line[i])
						gridBuilder.WriteByte(line[i])
					} else if line[i] == '@' {
						gridBuilder.WriteByte('@')
						gridBuilder.WriteByte('.')
					} else if line[i] == 'O' {
						gridBuilder.WriteByte('[')
						gridBuilder.WriteByte(']')
					}
				}

				grid = append(grid, gridBuilder.Bytes())
			} else {
				grid = append(grid, []byte(line))
			}
		} else {
			builder.WriteString(line)
		}
	}

	return grid, builder.Bytes()

}

func getColumn(grid *[][]byte, jIndex int) []byte {
	column := make([]byte, 0)

	for i := 0; i < len(*grid); i++ {
		column = append(column, (*grid)[i][jIndex])
	}

	return column
}

func displayGrid(grid *[][]byte) {
	for _, line := range *grid {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println()
	}
}
