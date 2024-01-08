package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
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
	minDistance := getMaxOrMinDistance(contents, false)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   9,
		Part:  1,
		Value: minDistance,
	})
}

func partTwo(contents []string) {
	maxDistance := getMaxOrMinDistance(contents, true)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   9,
		Part:  2,
		Value: maxDistance,
	})
}

// Generate permutations of a string using Heap's Algorithm (https://en.wikipedia.org/wiki/Heap%27s_algorithm)
func permutations(arr []string) [][]string {
	var helper func([]string, int)
	res := [][]string{}

	helper = func(arr []string, n int) {
		if n == 1 {
			tmp := make([]string, len(arr))
			copy(tmp, arr)
			res = append(res, tmp)
		} else {
			for i := 0; i < n; i++ {
				helper(arr, n-1)
				if n%2 == 1 {
					tmp := arr[i]
					arr[i] = arr[n-1]
					arr[n-1] = tmp
				} else {
					tmp := arr[0]
					arr[0] = arr[n-1]
					arr[n-1] = tmp
				}
			}
		}
	}
	helper(arr, len(arr))
	return res
}

func getMaxOrMinDistance(contents []string, isMax bool) int {
	var mappings = make(map[string]map[string]int)
	for _, line := range contents {
		var from, to, distance = "", "", 0
		_, err := fmt.Sscanf(line, "%s to %s = %d", &from, &to, &distance)
		if err != nil {
			panic("Error :c")
		}

		if _, ok := mappings[from]; !ok {
			mappings[from] = make(map[string]int)
			mappings[from][to] = distance
		} else {
			mappings[from][to] = distance
		}

		if _, ok := mappings[to]; !ok {
			mappings[to] = make(map[string]int)
			mappings[to][from] = distance
		} else {
			mappings[to][from] = distance
		}
	}

	// Get all permutations of the keys
	var keys = make([]string, len(mappings))
	var i = 0
	for k := range mappings {
		keys[i] = k
		i++
	}
	permutations := permutations(keys)

	var limitDistance int
	if isMax {
		limitDistance = 0
	} else {
		limitDistance = int(^uint(0) >> 1)
	}
	for _, permutation := range permutations {
		var distance = 0
		for j := 0; j < len(permutation)-1; j++ {
			distance += mappings[permutation[j]][permutation[j+1]]
		}
		if isMax {
			limitDistance = max(limitDistance, distance)
		} else {
			limitDistance = min(limitDistance, distance)
		}
	}

	return limitDistance
}
