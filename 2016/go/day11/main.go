package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

var ITEM_TYPE_TO_INDEX = map[string]int{
	"M": 0,
	"G": 1,
}

type state struct {
	steps         int
	elevatorFloor int
	floorState    []itemPair
}

type itemPair [2]int

// [FLOOR]-[ORDERED_MICROCHIPS]-[ORDERED_GENERATORS]
type stateKey map[string]bool

func partOne(contents []string) {
	// Get items per floor. Note we are offset by 1 due to zero-indexing
	items := parseInput(contents)

	// To optimise, change the itemsPerFloor into a less memory-intensive format

	numSteps := solve(state{
		steps:         0,
		elevatorFloor: 0,
		floorState:    items,
	})

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  1,
		Value: numSteps,
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  2,
		Value: "TODO",
	})
}

var visited = make(stateKey)

func solve(initial state) int {
	queue := []state{initial}
	iterations := 0

	for len(queue) > 0 {
		iterations++
		current := queue[0]
		queue = queue[1:]

		if isFinished(&current.floorState) {
			fmt.Println(iterations, "iterations")
			return current.steps
		}

		key := getStateKey(current)

		visited[key] = true // Only visit each state once

		fmt.Println(current.steps, "steps")

		nextStates := generateNextValidStates(current)
		queue = append(queue, nextStates...)
	}

	return -1 // No solution found if we've made it here
}

// This is the hard one... move each item ON THE ELEVATOR FLOOR up and down and check if valid
func generateNextValidStates(current state) []state {
	nextStates := make([]state, 0)

	// We can't simply loop items, because you can move more than one item at a time! Let's get the combinations of items we can move
	combinations := generateCombinations(current.floorState, current.elevatorFloor)
	for _, combo := range combinations {
		for dir := -1; dir <= 1; dir += 2 {
			newElevator := current.elevatorFloor + dir
			if newElevator < 0 || newElevator >= 4 {
				continue // Skip invalid elevator positions
			}

			// Try moving items and check if resulting state is valid

			// Clone the state, ensureing we do not pass ANYTHING by reference:
			newState := state{
				steps:         current.steps + 1,
				elevatorFloor: newElevator,
				floorState:    cloneFloorState(current.floorState),
			}

			// Move items
			// 2 Items:
			if len(combo) == 2 {
				newState.floorState[combo[0][0]][combo[0][1]] = newElevator
				newState.floorState[combo[1][0]][combo[1][1]] = newElevator
			} else {
				// 1 Item:
				newState.floorState[combo[0][0]][combo[0][1]] = newElevator
			}

			// Check it's not been visited:
			if visited[getStateKey(newState)] {
				continue
			}

			if isValidState(newState) {
				newState.steps = current.steps + 1
				nextStates = append(nextStates, newState)
			}
		}
	}

	return nextStates
}

func generateCombinations(items []itemPair, currentFloor int) [][][2]int {
	// First sweep through and get i,j coordinates of every item pert of the current floor:
	itemPartCoordsForCurrentFloor := make([][2]int, 0)
	for i := range items {
		for j := 0; j < 2; j++ {
			if items[i][j] == currentFloor {
				itemPartCoordsForCurrentFloor = append(itemPartCoordsForCurrentFloor, [2]int{i, j})
			}
		}
	}

	// We now have all the coordinates of the items on the current floor. We can take between 1 or 2 of these.
	// Generate all possible combinations of this:
	combinations := [][][2]int{{}}
	for _, item := range itemPartCoordsForCurrentFloor {
		for i := len(combinations) - 1; i >= 0; i-- {
			newCombo := append(combinations[i][:len(combinations[i]):len(combinations[i])], item)

			// Can only carry up to 2 items in the elevator
			if len(newCombo) <= 2 && len(newCombo) != 0 {
				combinations = append(combinations, newCombo)
			}
		}
	}

	return combinations[1:] // Skip the first empty combination
}

func cloneFloorState(floorState []itemPair) []itemPair {
	newFloorState := make([]itemPair, len(floorState))
	copy(newFloorState, floorState)
	return newFloorState
}

func isValidState(state state) bool {
	// Check if state is valid based on problem constraints
	// Valid if no generators on same floor as chip, unless they are paired

	dangerousGenerators := make([]int, 0)
	for _, pairs := range state.floorState {
		if pairs[0] != pairs[1] {
			// Not on the same floor... generator is potentially dangerous:
			dangerousGenerators = append(dangerousGenerators, pairs[ITEM_TYPE_TO_INDEX["G"]])
		}
	}
	for _, pairs := range state.floorState {
		if pairs[0] != pairs[1] {
			// Not on the same floor... are any generators on the same floor as the chip
			for _, gen := range dangerousGenerators {
				// Is the generator on this floor AND DOESNT HAVE ITS CHIP on the same floor
				if gen == pairs[ITEM_TYPE_TO_INDEX["M"]] {
					return false
				}
			}
		}
	}

	return true
}

var itemNameToIndex = map[string]int{}

func parseInput(contents []string) []itemPair {
	itemsState := make([]itemPair, 0)
	currentFloor := 0
	patternsByType := map[string]string{
		"M": `(\w+)-compatible microchip`,
		"G": `(\w+) generator`,
	}
	for i := 0; i < 4; i++ {
		line := contents[i]
		for itemType, pattern := range patternsByType {
			entries := regexp.MustCompile(pattern)
			matches := entries.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				val := match[1]
				if _, ok := itemNameToIndex[val]; !ok {
					itemsState = append(itemsState, itemPair{-1, -1})
					// Update the index map
					itemNameToIndex[val] = len(itemsState) - 1
				}
				// Retrieve the existing array from the map
				itemsState[itemNameToIndex[val]][ITEM_TYPE_TO_INDEX[itemType]] = currentFloor
			}
		}
		currentFloor++
	}

	return itemsState
}

func isFinished(items *[]itemPair) bool {
	// If the top floor has all of the items, we're finished. We'll go up each floor and check (could optimise by using count of items if needed)
	isFinished := true
	// Check the first 3 floors for items:
	for _, pairs := range *items {
		if pairs[0] != 3 || pairs[1] != 3 {
			isFinished = false
			break
		}
	}
	return isFinished
}

func getStateKey(state state) string {
	// HINT: THE MOST IMPORTANT, ABSOLUTELY ESSENTIAL: ALL PAIRS ARE INTERCHANGEABLE
	// To do this, order the items by firstly their values in index 0, then in index 1 IF THE VALUES ARE THE SAME IN INDEX 0
	floorState := state.floorState
	sort.Slice(floorState, func(i, j int) bool {
		if floorState[i][0] != floorState[j][0] {
			return floorState[i][0] < floorState[j][0]
		}
		return floorState[i][1] < floorState[j][1]
	})

	// The string is elevatorFloor-state:
	return fmt.Sprintf("%d-%v", state.elevatorFloor, floorState)
}
