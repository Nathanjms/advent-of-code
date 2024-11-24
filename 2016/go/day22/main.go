package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
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
	parseDisks(contents)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   22,
		Part:  1,
		Value: "TODO",
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   22,
		Part:  2,
		Value: "TODO",
	})
}

type disk struct {
	totalSize int
	used      int
}

func parseDisks(contents []string) {
	var disks [][]disk

	for i, line := range contents {
		var x, y, totalSize, used int
		if i <= 1 {
			continue
		}
		// Parse line from eg '/dev/grid/node-x0-y0     88T   67T    21T   76%'          Size  Used  Avail  Use%' using regex
		re := regexp.MustCompile(`/dev/grid/node-x(\d+)-y(\d+)\s+(\d+)T\s+(\d+)T\s+(\d+)T\s+\d+%`)
		matches := re.FindStringSubmatch(line)

		x, _ = strconv.Atoi(matches[1])
		y, _ = strconv.Atoi(matches[2])
		totalSize, _ = strconv.Atoi(matches[2])
		used, _ = strconv.Atoi(matches[3])

		// if

		fmt.Println(matches)
	}
}

func (d disk) free() int {
	return d.totalSize - d.used
}
