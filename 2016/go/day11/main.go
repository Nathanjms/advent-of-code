package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"math"
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

var namesToId = map[string]int{
	"thulium":    0,
	"plutonium":  1,
	"strontium":  2,
	"promethium": 3,
	"ruthenium":  4,
	"hydrogen":   5,
}

var itemTypeToId = map[string]int{
	"M": 64,
	"G": 63,
}

type item struct {
	Type string
	Name string
	ID   string
}

type queueItem struct {
	Steps         int
	CurrState     map[int][]item
	CurrLiftFloor int
}

func partOne(contents []string) {
	// Get items per floor. Note we are offset by 1 due to zero-indexing
	itemsPerFloor := parseInput(contents)
	displayFloors(&itemsPerFloor, 0)

	// Lets try for a BFS
	queue := make([]queueItem, 0)

	queue = append(queue, queueItem{
		Steps:         0,
		CurrState:     itemsPerFloor,
		CurrLiftFloor: 0,
	})

	var seen []map[int][]item // We could optimise this to just be the IDs - try this for now!

	steps := math.MaxInt
	var nextItem queueItem
	for {
		nextItem, queue = queue[0], queue[1:] // Grab the next entry in the queue

		// Have we already seen this?
		if isFinished(&nextItem.CurrState) {
			steps = nextItem.Steps
			break
		}

		// Get all possibilities and add them to queue if not visited:
		// This is the hard bit...
		possibilities := getPossibilities(&nextItem)

	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  1,
		Value: steps,
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  2,
		Value: "TODO",
	})
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
			matches := entries.FindAllString(line, -1)
			for _, match := range matches {
				tmp := strings.Split(match, "-")
				itemsPerFloor[currentFloor] = append(itemsPerFloor[currentFloor],
					item{
						Name: tmp[0],
						Type: itemType,
						ID:   fmt.Sprintf("%d-%d", namesToId[tmp[0]], itemTypeToId[itemType]),
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
}

func isFinished(itemsPerFloor *map[int][]item) bool {
	// Check the first 3 floors for items:
	for i := 0; i < 3; i++ {
		if len((*itemsPerFloor)[i]) > 0 {
			return false
		}
	}
	// If we've made it here, all are on the top floor!
	return true
}

func getPossibilities(itemsOnFloor *queueItem) []map[int][]item {
	possibilities := make([]map[int][]item, 0)
	for _, item := range itemsOnFloor.CurrState[itemsOnFloor.CurrLiftFloor] {
		// For each item, we can move it up or down. Keep it in range though.
		newPossibility := make(map[int][]item)
		// UP:
		if itemsOnFloor.CurrLiftFloor != 3 {
			// Move item up a floor!
			newPossibility = itemsOnFloor.CurrState
			// Remove the item from its current floor, move it up to the next floor:

		}
		// DOWN:
		if itemsOnFloor.CurrLiftFloor != 0 {
			newPossibility = ite
		}
	}
}
