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

	fmt.Println(inputPath)

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

type Boundary struct {
	lower int
	upper int
}

func partOne(contents []string) {
	var boundaries []Boundary
	for _, line := range contents {
		temp := strings.Split(line, "-")
		lower, _ := strconv.Atoi(temp[0])
		upper, _ := strconv.Atoi(temp[1])
		boundaries = append(boundaries, Boundary{
			lower: int(lower),
			upper: int(upper),
		})
	}

	sort.Slice(boundaries, func(i, j int) bool {
		return boundaries[i].lower < boundaries[j].lower
	})

	testIp := int(0)
	currentRelevantBoundaryIndex := 0

	for {
		if testIp > math.MaxUint32 {
			panic("Too Big")
		}
		if testIp < boundaries[currentRelevantBoundaryIndex].lower {
			break
		} else if testIp <= boundaries[currentRelevantBoundaryIndex].upper {
			// within the range so give up and bump the testIp up to the first value outside the bound
			testIp = boundaries[currentRelevantBoundaryIndex].upper + 1
			currentRelevantBoundaryIndex++
			continue
		} else if testIp > boundaries[currentRelevantBoundaryIndex].upper {
			currentRelevantBoundaryIndex++
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   20,
		Part:  1,
		Value: testIp,
	})
}

func partTwo(contents []string) {
	var boundaries []Boundary
	for _, line := range contents {
		temp := strings.Split(line, "-")
		lower, _ := strconv.Atoi(temp[0])
		upper, _ := strconv.Atoi(temp[1])
		boundaries = append(boundaries, Boundary{
			lower: int(lower),
			upper: int(upper),
		})
	}

	sort.Slice(boundaries, func(i, j int) bool {
		return boundaries[i].lower < boundaries[j].lower
	})

	testIp := int(0)
	currentRelevantBoundaryIndex := 0
	numAllowed := 0

	for {
		if testIp > math.MaxUint32 {
			break
		}
		if currentRelevantBoundaryIndex > len(boundaries)-1 {
			numAllowed += (math.MaxUint32 - testIp)
			break
		}
		if testIp < boundaries[currentRelevantBoundaryIndex].lower {
			numAllowed++
			testIp++
			continue
		} else if testIp <= boundaries[currentRelevantBoundaryIndex].upper {
			// within the range so give up and bump the testIp up to the first value outside the bound
			testIp = boundaries[currentRelevantBoundaryIndex].upper + 1
			currentRelevantBoundaryIndex++
			continue
		} else if testIp > boundaries[currentRelevantBoundaryIndex].upper {
			currentRelevantBoundaryIndex++
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   20,
		Part:  2,
		Value: numAllowed,
	})
}
