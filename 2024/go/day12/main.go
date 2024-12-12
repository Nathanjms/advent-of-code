package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
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

type plotStruct struct {
	letter    byte
	points    [][2]int
	perimeter int
	sides     int
}

func partOne(contents []string) {
	// Use a BFS and not just a loop;
	plots := buildPlots(contents, false)

	totalPrice := 0
	for i := 0; i < len(plots); i++ {
		totalPrice += (len(plots[i].points) * plots[i].perimeter)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   12,
		Part:  1,
		Value: totalPrice,
	})
}

func partTwo(contents []string) {
	// Use a BFS and not just a loop;
	plots := buildPlots(contents, true)

	totalPrice := 0
	for i := 0; i < len(plots); i++ {
		totalPrice += (len(plots[i].points) * plots[i].sides)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   12,
		Part:  2,
		Value: totalPrice,
	})
}

type queueStruct struct {
	pos      [2]int
	currChar byte
}

func buildPlots(contents []string, partTwo bool) []plotStruct {
	plots := make([]plotStruct, 0)
	visited := make(map[[2]int]bool, 0)

	for i, line := range contents {
		for j := range line {
			if _, ok := visited[[2]int{i, j}]; ok {
				continue
			}

			plots = append(plots, buildPlot(i, j, contents, &visited, partTwo))
		}
	}

	return plots
}

func buildPlot(i int, j int, contents []string, visited *map[[2]int]bool, partTwo bool) plotStruct {
	plot := plotStruct{}

	directions := [4][2]int{
		{1, 0},  // Down
		{0, 1},  // Right
		{-1, 0}, // Up
		{0, -1}, // Left
	}

	currentValue := contents[i][j]
	plot.letter = currentValue
	queue := make([]queueStruct, 0)
	queue = append(queue, queueStruct{
		pos:      [2]int{i, j},
		currChar: currentValue,
	})

	var element queueStruct

	for {
		if len(queue) == 0 {
			break
		}

		// Grab the next element in queue
		element, queue = queue[0], queue[1:]

		// if visited, exit
		_, ok := (*visited)[element.pos]
		if ok {
			continue
		}

		(*visited)[element.pos] = true

		plot.points = append(plot.points, element.pos)

		for _, dir := range directions {
			newI := element.pos[0] + dir[0]
			newJ := element.pos[1] + dir[1]

			// Out of bounds checks
			if newI < 0 || newJ < 0 || newI > len(contents)-1 || newJ > len(contents[0])-1 {
				plot.perimeter++
				continue
			}

			// Rules check; it must be the same char
			if plot.letter != contents[newI][newJ] {
				plot.perimeter++
				continue
			}

			queue = append(queue, queueStruct{
				pos:      [2]int{newI, newJ},
				currChar: element.currChar,
			})
		}
	}

	if partTwo {
		corners := 0
		// For part 2, we count sides. This is equivalent to counting corners which is easier
		orthogonalPairs := [][2][2]int{}
		for i := 0; i < 4; i++ {
			orthogonalPairs = append(orthogonalPairs, [2][2]int{directions[i], directions[(i+1)%4]})
		}
		for _, point := range plot.points {
			cornerCount := 0
			for _, pair := range orthogonalPairs {
				// To check fo a corner, we check that the either: At least one orthogonal direction pair is not in plot OR both match but diagonal is not in plot
				posOne := [2]int{point[0] + pair[0][0], point[1] + pair[0][1]}
				posTwo := [2]int{point[0] + pair[1][0], point[1] + pair[1][1]}
				// 1. One orth. direction pair is not in plot
				if !isInSlice(plot.points, posOne) && !isInSlice(plot.points, posTwo) {
					cornerCount++
					continue
				}

				//2.  both match but diagonal is not in plot
				diagonal := [2]int{point[0] + pair[0][0] + pair[1][0], point[1] + pair[0][1] + pair[1][1]}
				if isInSlice(plot.points, posOne) && isInSlice(plot.points, posTwo) && !isInSlice(plot.points, diagonal) {
					cornerCount++
					continue
				}

			}
			if cornerCount > 0 {
				corners += cornerCount
			}
		}

		plot.sides = corners
	}

	return plot
}

func isInSlice(haystack [][2]int, needle [2]int) bool {
	for _, point := range haystack {
		if point == needle {
			return true
		}
	}
	return false
}
