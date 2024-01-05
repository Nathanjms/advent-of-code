package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
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

func main() {
	// If another cmd argument has been passed, use that as the input path:
	if len(os.Args) > 1 {
		inputPath = os.Args[1]
	}

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	totalToOrder := 0

	for _, line := range contents {
		// Explode line by 'x':
		temp := strings.Split(line, "x")
		length, _ := strconv.Atoi(temp[0])
		width, _ := strconv.Atoi(temp[1])
		height, _ := strconv.Atoi(temp[2])

		totalToOrder += 2 * (length*width + width*height + height*length)
		// Also need to add the smallest side area:
		totalToOrder += min(length*width, width*height, height*length)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   2,
		Part:  1,
		Value: totalToOrder,
	})
}

func partTwo(contents []string) {
	totalToOrder := 0

	for _, line := range contents {
		// Explode line by 'x':
		temp := strings.Split(line, "x")
		length, _ := strconv.Atoi(temp[0])
		width, _ := strconv.Atoi(temp[1])
		height, _ := strconv.Atoi(temp[2])

		smallestFace := 2 * min(length+width, width+height, height+length)
		volume := length * width * height
		totalToOrder += smallestFace + volume
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   2,
		Part:  1,
		Value: totalToOrder,
	})
}
