package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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
	lowestPoints := bfs(contents)

	// BFS today I think!
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   16,
		Part:  1,
		Value: lowestPoints,
	})
}

func partTwo(contents []string) {
	totalSpotsOnAnyPath := dfs(contents, bfs(contents))

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   16,
		Part:  2,
		Value: totalSpotsOnAnyPath,
	})
}

type queueStruct struct {
	pos                [2]int
	currPoints         int64
	currDirectionIndex int
}

type cacheKey struct {
	pos      [2]int
	dirIndex int
}

var directions = [4][2]int{
	{1, 0},  // Down
	{0, 1},  // Right
	{-1, 0}, // Up
	{0, -1}, // Left
}

func bfs(contents []string) int64 {
	queue := make([]queueStruct, 0)
	visited := make(map[cacheKey]bool)
	start, end := findStartAndEnd(contents)

	queue = append(queue, queueStruct{
		pos:                start,
		currPoints:         int64(0),
		currDirectionIndex: 1, // Starting facing right according to input
	})

	var element queueStruct

	for {
		if len(queue) == 0 {
			break
		}

		// Take the element with the lowest points from the queue (should be the first)
		element, queue = queue[0], queue[1:]

		if element.pos == end {
			return element.currPoints
		}

		// If visited, then we've been here with less points and so can skip
		if _, ok := visited[cacheKey{element.pos, element.currDirectionIndex}]; ok {
			continue
		}

		visited[cacheKey{element.pos, element.currDirectionIndex}] = true

		// We only can go 90 degrees or stay the same, so just do 3 indexes:
		for i := -1; i <= 1; i++ {
			newDirIndex := (element.currDirectionIndex + i + 4) % 4
			dir := directions[newDirIndex]

			newI := element.pos[0] + dir[0]
			newJ := element.pos[1] + dir[1]

			// Out of bounds checks - we can actually just check for a '#' symbol as we are surrounded by walls
			if contents[newI][newJ] == '#' {
				continue
			}

			// Otherwise, we make the move and add to queue
			if i == 0 {
				// Just stepped forward - 1 point added;
				queue = append(queue, queueStruct{
					pos:                [2]int{newI, newJ},
					currPoints:         element.currPoints + 1,
					currDirectionIndex: element.currDirectionIndex,
				})
			} else {
				// Turned and stepped - 1000 + 1 point added
				queue = append(queue, queueStruct{
					pos:                [2]int{newI, newJ},
					currPoints:         element.currPoints + 1001,
					currDirectionIndex: newDirIndex,
				})
			}

		}

		// Sort the queue by points (might be expensive but enables BFS so probably worth it?):
		// Sort by lowest cost and then go upwards
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].currPoints < queue[j].currPoints
		})
	}

	return math.MaxInt64
}

type queueStructPt2 struct {
	pos                [2]int
	currPoints         int64
	currDirectionIndex int
	visitedPoints      map[[2]int]bool
}

func dfs(contents []string, pointsTarget int64) int {
	queue := make([]queueStructPt2, 0)
	start, end := findStartAndEnd(contents)

	allShortestPathCoordinates := make(map[[2]int]bool)
	visitedPoints := map[[2]int]bool{}

	queue = append(queue, queueStructPt2{
		pos:                start,
		currPoints:         int64(0),
		currDirectionIndex: 1, // Starting facing right according to input
		visitedPoints:      visitedPoints,
	})

	var element queueStructPt2

	for {
		if len(queue) == 0 {
			break
		}

		// Take the element with the highest points from the queue (should be the first)
		element = queue[len(queue)-1]
		queue = queue[:len(queue)-1]

		if element.currPoints > pointsTarget {
			// If it's exceeded the shortest path's score, we can drop this path
			continue
		}
		fmt.Println(len(queue), element.currPoints, len(element.visitedPoints))

		// has this element been on this spot before?
		if _, ok := element.visitedPoints[element.pos]; ok {
			continue
		}

		element.visitedPoints[element.pos] = true

		if element.pos == end {
			addToShortestPathCoordinates(&allShortestPathCoordinates, element.visitedPoints)
			continue
		}

		// We only can go 90 degrees or stay the same, so just do 3 indexes:
		for i := -1; i <= 1; i++ {
			newDirIndex := (element.currDirectionIndex + i + 4) % 4
			dir := directions[newDirIndex]

			newI := element.pos[0] + dir[0]
			newJ := element.pos[1] + dir[1]

			// Out of bounds checks - we can actually just check for a '#' symbol as we are surrounded by walls
			if contents[newI][newJ] == '#' {
				continue
			}

			if _, ok := element.visitedPoints[[2]int{newI, newJ}]; ok {
				continue
			}

			// Otherwise, we make the move and add to queue
			if i == 0 && element.currPoints+1 <= pointsTarget {
				// Just stepped forward - 1 point added;
				queue = append(queue, queueStructPt2{
					pos:                [2]int{newI, newJ},
					currPoints:         element.currPoints + 1,
					currDirectionIndex: element.currDirectionIndex,
					visitedPoints:      copyMap(element.visitedPoints),
				})
			} else if element.currPoints+1000 <= pointsTarget {
				// Turned and stepped - 1000 + 1 point added
				queue = append(queue, queueStructPt2{
					pos:                [2]int{newI, newJ},
					currPoints:         element.currPoints + 1001,
					currDirectionIndex: newDirIndex,
					visitedPoints:      copyMap(element.visitedPoints),
				})
			}

		}

		// Sort the queue by points (might be expensive but enables BFS so probably worth it?):
		// Sort by lowest cost and then go upwards
		sort.Slice(queue, func(i, j int) bool {
			return queue[i].currPoints < queue[j].currPoints
		})
	}

	return len(allShortestPathCoordinates)
}

func addToShortestPathCoordinates(visited *map[[2]int]bool, newVisitedCoordinates map[[2]int]bool) {
	for coord := range newVisitedCoordinates {
		(*visited)[coord] = true
	}
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

func copyMap(m map[[2]int]bool) map[[2]int]bool {
	m2 := make(map[[2]int]bool, len(m))
	var id [2]int
	for id = range m {
		m2[id] = true
	}
	return m2
}
