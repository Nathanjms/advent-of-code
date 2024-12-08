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

func partOne(contents []string) {
	// First go through and build a map of coordinates by char;
	coordinatesByChar := make(map[byte][][2]int, 0)

	antiNodesMap := make(map[[2]int]bool, 0)

	for i := 0; i < len(contents); i++ {
		for j := 0; j < len(contents[i]); j++ {
			if contents[i][j] != '.' {
				coordinatesByChar[contents[i][j]] = append(coordinatesByChar[contents[i][j]], [2]int{i, j})
			}
		}
	}

	// Now brute force each grid point to check if we satisfy the criteria;
	for i := 0; i < len(contents); i++ {
		for j := 0; j < len(contents[i]); j++ {
			testNode := [2]int{i, j}
			isValid := false

			// have we been here before?
			_, ok := antiNodesMap[testNode]
			// If the key exists
			if ok {
				continue // Already been confirmed as an antenna
			}

			// Otherwise, check whether it could work as an antinode
			// eg. test node 1,3: 3,4 is 2,1 | 5,5 is 4,2 which is double, and diff between 2,1 5,5 is 4,2
			for char, coords := range coordinatesByChar {
				if char == contents[i][j] {
					continue // Skip for the character type where there is a node here (but not for others)
				}
				// The difference between test coord and coord A must equal the diff between coord A cna coord B... maybe?
				for n := 0; n < len(coords); n++ {
					for m := 0; m < len(coords); m++ {
						if n == m {
							continue
						}

						diffCoordAAndTest := [2]int{
							coords[n][0] - testNode[0],
							coords[n][1] - testNode[1],
						}

						diffCoordBAndCoordA := [2]int{
							coords[m][0] - coords[n][0],
							coords[m][1] - coords[n][1],
						}

						if diffCoordAAndTest[0] == diffCoordBAndCoordA[0] && diffCoordAAndTest[1] == diffCoordBAndCoordA[1] {
							isValid = true
							break
						}
					}
				}
			}
			if isValid {
				antiNodesMap[testNode] = true
			}
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   8,
		Part:  1,
		Value: len(antiNodesMap),
	})
}

func partTwo(contents []string) {
	// First go through and build a map of coordinates by char;
	coordinatesByChar := make(map[byte][][2]int, 0)

	antiNodesMap := make(map[[2]int]bool, 0)

	for i := 0; i < len(contents); i++ {
		for j := 0; j < len(contents[i]); j++ {
			if contents[i][j] != '.' {
				coordinatesByChar[contents[i][j]] = append(coordinatesByChar[contents[i][j]], [2]int{i, j})
			}
		}
	}

	// Now brute force each grid point to check if we satisfy the criteria;
	for i := 0; i < len(contents); i++ {
		for j := 0; j < len(contents[i]); j++ {
			testNode := [2]int{i, j}
			isValid := false

			// have we been here before?
			_, ok := antiNodesMap[testNode]
			// If the key exists
			if ok {
				continue // Already been confirmed as an antenna
			}

			// Otherwise, check whether it could work as an antinode
			// eg. test node 1,3: 3,4 is 2,1 | 5,5 is 4,2 which is double, and diff between 2,1 5,5 is 4,2
			for char, coords := range coordinatesByChar {
				if char == contents[i][j] {
					isValid = true
					break
				}
				// The difference between test coord and coord A must equal a MULTIPLE OF the diff between coord A and coord B... maybe?
				for n := 0; n < len(coords); n++ {
					for m := 0; m < len(coords); m++ {
						if n == m {
							continue
						}

						diffCoordBAndCoordA := [2]int{
							coords[m][0] - coords[n][0],
							coords[m][1] - coords[n][1],
						}

						if isOnPath(testNode, diffCoordBAndCoordA, coords[n], contents) {
							isValid = true
							break
						}
					}
				}
			}
			if isValid {
				antiNodesMap[testNode] = true
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   8,
		Part:  2,
		Value: len(antiNodesMap),
	})
}

func isOnPath(testNode [2]int, diffCoordBAndCoordA [2]int, targetCoord [2]int, contents []string) bool {
	// Go through multiples until we exceed the range or have found it!
	probeCoord := [2]int{testNode[0], testNode[1]}
	multiplier := 1
	outOfRangeCount := 0

	for {
		outOfRangeCount = 0
		addition := [2]int{
			multiplier * diffCoordBAndCoordA[0],
			multiplier * diffCoordBAndCoordA[1],
		}

		probeCoordOne := [2]int{
			addition[0] + probeCoord[0],
			addition[1] + probeCoord[1],
		}

		probeCoordTwo := [2]int{
			(-1 * addition[0]) + probeCoord[0],
			(-1 * addition[1]) + probeCoord[1],
		}
		if probeCoordOne[0] < 0 || probeCoordOne[0] > len(contents)-1 || probeCoordOne[1] < 0 || probeCoordOne[1] > len(contents[0])-1 {
			outOfRangeCount++
		}

		if probeCoordTwo[0] < 0 || probeCoordTwo[0] > len(contents)-1 || probeCoordTwo[1] < 0 || probeCoordTwo[1] > len(contents[0])-1 {
			outOfRangeCount++
		}

		if outOfRangeCount == 2 {
			return false
		}

		if (probeCoordOne[0] == targetCoord[0] && probeCoordOne[1] == targetCoord[1]) ||
			(probeCoordTwo[0] == targetCoord[0] && probeCoordTwo[1] == targetCoord[1]) {
			return true
		}

		multiplier++
	}
}
