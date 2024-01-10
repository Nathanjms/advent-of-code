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
	// Parse the input. Each line is "Alice would gain 54 happiness units by sitting next to Bob." - we want name, gain/lose, amount, name
	mapping := make(map[string]map[string]int)

	for _, line := range contents {
		var name, gainOrLose, name2 string
		var amount int

		line = line[:len(line)-1] // Remove the period at the end
		_, err := fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &name, &gainOrLose, &amount, &name2)
		if err != nil {
			panic(err)
		}

		// Add to map for the source name, creating for this key if it doesn't exist
		if mapping[name] == nil {
			mapping[name] = make(map[string]int)
		}
		if gainOrLose == "lose" {
			amount = -amount
		}
		mapping[name][name2] = amount
	}

	// As in day 9, we can use Heap's Algorithm to generate all permutations of the names, and find the max by brute force.
	// Permutations is a bit overkill quite needed, because 1,2,3,4 is the same as 2,3,4,1, but it's easier to use what we already have.
	var names []string
	for name := range mapping {
		names = append(names, name)
	}

	maxHappiness := 0
	for _, permutation := range permutations(names) {
		// TO the left is the index -1, to the right is +1. If we're at the start or end, wrap around.
		permutationHappiness := 0
		for i, name := range permutation {
			leftIndex := (i - 1 + len(permutation)) % len(permutation)
			if leftIndex < 0 {
				leftIndex = len(permutation) - 1
			}
			rightIndex := (i + 1) % len(permutation)
			if rightIndex >= len(permutation) {
				rightIndex = 0
			}
			permutationHappiness += mapping[name][permutation[leftIndex]] + mapping[name][permutation[rightIndex]]
		}
		maxHappiness = max(maxHappiness, permutationHappiness)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   13,
		Part:  1,
		Value: maxHappiness,
	})
}

func partTwo(contents []string) {
	// Parse the input. Each line is "Alice would gain 54 happiness units by sitting next to Bob." - we want name, gain/lose, amount, name
	mapping := make(map[string]map[string]int)

	for _, line := range contents {
		var name, gainOrLose, name2 string
		var amount int

		line = line[:len(line)-1] // Remove the period at the end
		_, err := fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &name, &gainOrLose, &amount, &name2)
		if err != nil {
			panic(err)
		}

		// Add to map for the source name, creating for this key if it doesn't exist
		if mapping[name] == nil {
			mapping[name] = make(map[string]int)
		}
		if gainOrLose == "lose" {
			amount = -amount
		}
		mapping[name][name2] = amount
	}

	// Need to add "Nathan" to the map, with 0 happiness for each other person, and each other person with 0 happiness for Nathan
	mapping["Nathan"] = make(map[string]int)
	for name := range mapping {
		mapping[name]["Nathan"] = 0
		mapping["Nathan"][name] = 0
	}

	// As in day 9, we can use Heap's Algorithm to generate all permutations of the names, and find the max by brute force.
	// Permutations is a bit overkill quite needed, because 1,2,3,4 is the same as 2,3,4,1, but it's easier to use what we already have.
	var names []string
	for name := range mapping {
		names = append(names, name)
	}

	maxHappiness := 0
	for _, permutation := range permutations(names) {
		// TO the left is the index -1, to the right is +1. If we're at the start or end, wrap around.
		permutationHappiness := 0
		for i, name := range permutation {
			leftIndex := (i - 1 + len(permutation)) % len(permutation)
			if leftIndex < 0 {
				leftIndex = len(permutation) - 1
			}
			rightIndex := (i + 1) % len(permutation)
			if rightIndex >= len(permutation) {
				rightIndex = 0
			}
			permutationHappiness += mapping[name][permutation[leftIndex]] + mapping[name][permutation[rightIndex]]
		}
		maxHappiness = max(maxHappiness, permutationHappiness)
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   13,
		Part:  2,
		Value: maxHappiness,
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
