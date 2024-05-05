package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"crypto/md5"
	"encoding/hex"
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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

type state struct {
	initialHash      []byte
	position         [2]int
	directionHistory []byte
}

var directionToCoordMap = map[byte][2]int{
	'U': {0, -1},
	'D': {0, 1},
	'L': {-1, 0},
	'R': {1, 0},
}

var visited = make(map[[2]int]bool)

func partOne(contents string) {
	initialState := state{
		initialHash:      []byte(contents),
		position:         [2]int{0, 0},
		directionHistory: []byte{},
	}

	solution := solve(initialState)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   17,
		Part:  1,
		Value: solution,
	})
}

func partTwo(contents string) {
	initialState := state{
		initialHash:      []byte(contents),
		position:         [2]int{0, 0},
		directionHistory: []byte{},
	}

	solution := solvePt2(initialState)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   17,
		Part:  2,
		Value: solution,
	})
}

func solve(initial state) string {
	queue := []state{initial}
	iterations := 0

	for len(queue) > 0 {
		iterations++
		current := queue[0]
		queue = queue[1:]

		if current.position[0] == 3 && current.position[1] == 3 {
			return string(current.directionHistory)
		}

		// Note: We DO NOT block revisits

		nextStates := generateNextValidStates(current)
		queue = append(queue, nextStates...)
	}

	return ""
}

func solvePt2(initial state) int {
	finishedPaths := make(map[string]bool)
	results := make([]int, 0)
	queue := []state{initial}
	iterations := 0

	for len(queue) > 0 {
		iterations++
		current := queue[0]
		queue = queue[1:]

		if current.position[0] == 3 && current.position[1] == 3 {
			results = append(results, len(current.directionHistory))
			finishedPaths[string(current.directionHistory)] = true
			continue
		}

		if finishedPaths[string(current.directionHistory)] {
			continue
		}

		nextStates := generateNextValidStates(current)
		queue = append(queue, nextStates...)
	}

	return results[len(results)-1]
}

func generateNextValidStates(current state) []state {
	nextStates := make([]state, 0)

	isLockedStates := generateIsLockedStates(append(current.initialHash, current.directionHistory...))

	for _, direction := range []byte{'U', 'D', 'L', 'R'} {
		pos := [2]int{current.position[0] + directionToCoordMap[direction][0], current.position[1] + directionToCoordMap[direction][1]}

		// First, is it out of bounds?
		if pos[0] < 0 || pos[0] > 3 || pos[1] < 0 || pos[1] > 3 {
			continue
		}

		if isLockedStates[direction] {
			continue
		}

		newDirection := make([]byte, len(current.directionHistory))
		copy(newDirection, current.directionHistory)
		newDirection = append(newDirection, direction)

		nextStates = append(nextStates, state{
			initialHash:      current.initialHash,
			position:         pos,
			directionHistory: newDirection,
		})
	}

	return nextStates
}

func generateIsLockedStates(input []byte) map[byte]bool {
	lockStates := map[byte]bool{
		'U': true,
		'D': true,
		'L': true,
		'R': true,
	}

	hash := getMD5Hash(input)

	for i, direction := range []byte{'U', 'D', 'L', 'R'} {
		// It is OPEN if Any b, c, d, e, or f means that the corresponding door is open:
		if hash[i] == 'b' || hash[i] == 'c' || hash[i] == 'd' || hash[i] == 'e' || hash[i] == 'f' {
			lockStates[direction] = false
		}
	}

	return lockStates
}

func getMD5Hash(text []byte) string {
	hash := md5.Sum(text)
	return hex.EncodeToString(hash[:])
}
