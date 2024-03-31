package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
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

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

var namesToId = map[string]string{
	"lithium":    "LI",
	"thulium":    "TH",
	"plutonium":  "PL",
	"strontium":  "ST",
	"promethium": "PR",
	"ruthenium":  "RU",
	"hydrogen":   "HY",
}

type item struct {
	Type string
	Name string
	ID   string
}

type state struct {
	steps         int
	floorState    map[int][]item
	elevatorFloor int
}

type stateKey struct {
	floors   string
	elevator int
}

func partOne(contents []string) {
	// Get items per floor. Note we are offset by 1 due to zero-indexing
	itemsPerFloor := parseInput(contents)
	displayFloors(&itemsPerFloor, 0)

	numSteps := solve(state{
		steps:         0,
		floorState:    itemsPerFloor,
		elevatorFloor: 0,
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

func solve(initial state) int {
	queue := []state{initial}
	visited := make(map[stateKey]bool)
	iterations := 0

	for len(queue) > 0 {
		iterations++
		current := queue[0]
		queue = queue[1:]

		if isFinished(&current.floorState) {
			fmt.Println(iterations, "iterations")
			return current.steps
		}

		nextStates := generateNextStates(current)
		for _, nextState := range nextStates {
			key := getStateKey(nextState)
			if !visited[key] {
				queue = append(queue, nextState)
				visited[key] = true
			}
		}
	}

	fmt.Println(iterations, "iterations")

	return -1 // No solution found if we've made it here
}

// This is the hard one... move each item ON THE ELEVATOR FLOOR up and down and check if valid
func generateNextStates(current state) []state {
	nextStates := make([]state, 0)
	itemsOnCurrentFloor := current.floorState[current.elevatorFloor]

	// We can't simply loop items, because you can move more than one item at a time! Let's get the combinations of items we can move
	combinations := generateCombinations(itemsOnCurrentFloor)
	for _, combo := range combinations {
		for dir := -1; dir <= 1; dir += 2 {
			newElevator := current.elevatorFloor + dir
			if newElevator < 0 || newElevator >= len(current.floorState) {
				continue // Skip invalid elevator positions
			}

			// Try moving items and check if resulting state is valid
			newState := moveItems(current, combo, newElevator)
			if isValidState(newState) {
				newState.steps = current.steps + 1
				nextStates = append(nextStates, newState)
			}
		}
	}

	return nextStates
}

func generateCombinations(items []item) [][]item {
	// Generate all possible combinations of items to move (subset generation)
	combinations := [][]item{{}}
	for _, item := range items {
		for i := len(combinations) - 1; i >= 0; i-- {
			newCombo := append(combinations[i][:len(combinations[i]):len(combinations[i])], item)
			// Can only carry up to 2 items in the elevator
			if len(newCombo) <= 2 && len(newCombo) != 0 {
				combinations = append(combinations, newCombo)
			}
		}
	}
	return combinations
}

// Take the current state and move the items to the new floor
func moveItems(current state, items []item, newElevator int) state {
	// Create a new state with items moved to newElevator
	newState := state{
		elevatorFloor: newElevator,
		steps:         current.steps,
		floorState:    make(map[int][]item),
	}

	newState.floorState = cloneFloorState(current.floorState)

	for _, item := range items {
		newState.floorState[current.elevatorFloor] = removeItem(newState.floorState[current.elevatorFloor], item.ID)
		newState.floorState[newElevator] = append(newState.floorState[newElevator], item)
	}

	return newState
}

func cloneFloorState(floorState map[int][]item) map[int][]item {
	newFloorState := make(map[int][]item)
	for floor, items := range floorState {
		// Deep copy items slice for each floor
		newItems := make([]item, len(items))
		copy(newItems, items)
		newFloorState[floor] = newItems
	}
	return newFloorState
}

func removeItem(items []item, itemToRemoveID string) []item {
	result := []item{}
	for _, item := range items {
		if item.ID != itemToRemoveID {
			result = append(result, item)
		}
	}
	return result
}

func isValidState(state state) bool {
	// Check if state is valid based on problem constraints
	for _, floor := range state.floorState {
		if containsInvalidPair(floor) {
			return false
		}
	}
	return true
}

func containsInvalidPair(items []item) bool {
	generators := make([]item, 0)
	// Add all generators first...
	for _, item := range items {
		// Generators are safe by themselves
		if item.Type == "G" {
			generators = append(generators, item)
		}
	}
	// Then look at microchips...
	for _, item := range items {
		// All microchips cannot be left with a generator unless IT HAS ITS OWN PAIR
		if item.Type == "M" && len(generators) > 0 && !containsGeneratorPair(generators, item) {
			return true
		}
	}
	return false
}

func containsGeneratorPair(generators []item, microchip item) bool {
	for _, generator := range generators {
		if microchip.Name == generator.Name {
			return true
		}
	}
	return false
}

func parseInput(contents []string) map[int][]item {
	itemsPerFloor := make(map[int][]item)
	currentFloor := 0
	patternsByType := map[string]string{
		"M": `(\w+)-compatible microchip`,
		"G": `(\w+) generator`,
	}
	for i := 0; i < 4; i++ {
		itemsPerFloor[currentFloor] = make([]item, 0)
		line := contents[i]
		for itemType, pattern := range patternsByType {
			entries := regexp.MustCompile(pattern)
			matches := entries.FindAllStringSubmatch(line, -1)
			for _, match := range matches {
				val := match[1]
				itemsPerFloor[currentFloor] = append(itemsPerFloor[currentFloor],
					item{
						Name: val,
						Type: itemType,
						ID:   fmt.Sprintf("%s-%s", namesToId[val], itemType),
					},
				)
			}
		}
		currentFloor++
	}

	return itemsPerFloor
}

func displayFloors(itemsPerFloor *map[int][]item, elevatorFloor int) {
	for i := 3; i >= 0; i-- {
		items := (*itemsPerFloor)[i]
		strTmp := ""
		if elevatorFloor == i {
			strTmp += "E "
		} else {
			strTmp += "  "
		}
		for _, item := range items {
			strTmp += item.ID + " "
		}
		fmt.Printf("\r%s%d: %s\n", "F", i+1, strTmp)
	}
	fmt.Println("")
	fmt.Println("")
}

func isFinished(itemsPerFloor *map[int][]item) bool {
	// If the top floor has all of the items, we're finished. We'll go up each floor and check (could optimise by using count of items if needed)
	// Check the first 3 floors for items:
	for i := 0; i < 3; i++ {
		if len((*itemsPerFloor)[i]) > 0 {
			return false
		}
	}
	// If we've made it here, all are on the top floor!
	return true
}

func getStateKey(state state) stateKey {
	var floorsStr strings.Builder
	for _, floor := range state.floorState {
		for _, item := range floor {
			floorsStr.WriteString(item.ID)
			floorsStr.WriteString("|")
		}
		floorsStr.WriteString(",")
	}
	return stateKey{floors: floorsStr.String(), elevator: state.elevatorFloor}
}
