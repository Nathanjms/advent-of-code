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

func partOne(contents []string) {
	itemsPerFloor := parseInput(contents)
	displayFloors(itemsPerFloor, 0)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  1,
		Value: "TODO",
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  2,
		Value: "TODO",
	})
}

type item struct {
	Type string
	Name string
	ID   string
}

func parseInput(contents []string) map[int][]item {
	itemsPerFloor := make(map[int][]item)
	currFloor := 0
	patternsByType := map[string]string{
		"M": `(\w+)-compatible microchip`,
		"G": `(\w+) generator`,
	}
	for i := 0; i < 4; i++ {
		itemsPerFloor[currFloor] = make([]item, 0)
		line := contents[i]
		for itemType, pattern := range patternsByType {
			entries := regexp.MustCompile(pattern)
			matches := entries.FindAllString(line, -1)
			for _, match := range matches {
				tmp := strings.Split(match, "-")
				itemsPerFloor[currFloor] = append(itemsPerFloor[currFloor],
					item{
						Name: tmp[0],
						Type: itemType,
						ID:   strings.ToUpper(string(tmp[0][0])) + itemType,
					},
				)
			}
		}
		currFloor++
	}

	return itemsPerFloor
}

func displayFloors(itemsPerFloor map[int][]item, elevatorFloor int) {
	for i := 3; i >= 0; i-- {
		items := itemsPerFloor[i]
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
