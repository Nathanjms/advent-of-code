package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

var DIRECTIONS = map[string][2]int{
	"N": {-1, 0},
	"E": {0, 1},
	"S": {1, 0},
	"W": {0, -1},
}

func partOne(contents string) {
	instructions := strings.Split(contents, ", ")

	facingDirection := "N"
	rowIndex := 0
	colIndex := 0

	for _, instruction := range instructions {
		direction := string(instruction[0])
		steps, _ := strconv.Atoi(instruction[1:])

		facingDirection = changeDirection(facingDirection, direction)
		rowIndex += DIRECTIONS[facingDirection][0] * steps
		colIndex += DIRECTIONS[facingDirection][1] * steps
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   1,
		Part:  1,
		Value: math.Abs(float64(rowIndex)) + math.Abs(float64(colIndex)),
	})
}

func partTwo(contents string) {
	instructions := strings.Split(contents, ", ")

	facingDirection := "N"
	rowIndex := 0
	colIndex := 0

	var twoVisits [2]int
	visited := make(map[string]bool, 0)
	visited["0,0"] = true

	for _, instruction := range instructions {
		direction := string(instruction[0])
		steps, _ := strconv.Atoi(instruction[1:])

		facingDirection = changeDirection(facingDirection, direction)

		for i := 0; i < steps; i++ {
			rowIndex += DIRECTIONS[facingDirection][0]
			colIndex += DIRECTIONS[facingDirection][1]

			if _, ok := visited[key(rowIndex, colIndex)]; !ok {
				visited[key(rowIndex, colIndex)] = true
			} else {
				twoVisits = [2]int{rowIndex, colIndex}
				break
			}
		}
		if twoVisits != [2]int{0, 0} {
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   1,
		Part:  2,
		Value: math.Abs(float64(twoVisits[0])) + math.Abs(float64(twoVisits[1])),
	})
}

func changeDirection(facingDirection string, direction string) string {
	directions := [4]string{"N", "E", "S", "W"}
	// Get the index of the current direction
	currIndex := 0
	for i, dir := range directions {
		if dir == facingDirection {
			currIndex = i
			break
		}
	}

	if direction == "R" {
		return directions[(currIndex+1)%4]
	} else {
		return directions[(currIndex-1+4)%4]
	}
}

func key(row int, col int) string {
	rowStr := strconv.Itoa(row)
	colStr := strconv.Itoa(col)
	return rowStr + "," + colStr
}
