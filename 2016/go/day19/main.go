package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
	"os"
	"path/filepath"
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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents string) {
	numElves, _ := strconv.Atoi(contents)
	// As this is the Josephus problem, if we express it as its largest power of 2 plus remainder r, the solution is 2*r + 1.

	l := 0
	power := 0
	for {
		// We could use some bitwise stuff here, but I don't really understand that so I'll keep it simple.
		value := math.Pow(2, float64(power))
		if value > float64(numElves) {
			l = int(float64(numElves) - math.Pow(2, float64(power-1)))
			break
		}
		power++
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   19,
		Part:  1,
		Value: 2*l + 1,
	})
}

func partTwo(contents string) {
	numElves, _ := strconv.Atoi(contents)

	elfList := make(map[int][2]int, 0)

	// This seems harder... we dont want a slice with >3 mil entries in if possible. Let's try it though, using a doubly-linked list
	for i := 0; i < numElves; i++ {
		prev := i - 1
		next := i + 1
		if i == 0 {
			prev = numElves - 1
		}

		if i == numElves-1 {
			next = 0
		}

		elfList[i] = [2]int{
			prev,
			next,
		}
	}

	/**
	Pseudocode from https://en.wikipedia.org/wiki/Doubly_linked_list
	node  := someNode
	do
		do something with node.value
		node  := node.next
	while node ≠ someNode
	*/

	index := 0
	current := elfList[index]
	numRemaining := numElves

	for {
		currentOpposite := elfList[int(math.Floor(float64(numRemaining)/2))] // Need to account for non-0 index!
		// Move the prev of the opposite to have the same prev value, but the next value from the removed one
		elfList[currentOpposite[0]] = [2]int{elfList[currentOpposite[0]][0], currentOpposite[1]}
		// Move the next of the opposite to have the same next value, but the prev value from the removed one
		elfList[currentOpposite[1]] = [2]int{currentOpposite[0], elfList[currentOpposite[1]][1]}
		// We remove this one by altering it's neighbours to skip it
		index = current[1]
		current = elfList[index]
		numRemaining--
		if numRemaining == 1 {
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   19,
		Part:  2,
		Value: index - 1,
	})
}
