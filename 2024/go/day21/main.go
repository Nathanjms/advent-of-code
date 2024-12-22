package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"fmt"
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

var NUM_KEYPAD = [][3]string{
	{"7", "8", "9"},
	{"4", "5", "6"},
	{"1", "2", "3"},
	{"N", "0", "A"},
}

var DIR_KEYPAD = [][3]string{
	{"N", "^", "A"},
	{"<", "v", ">"},
}

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
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   21,
		Part:  1,
		Value: computeTotal(contents, 2),
	})
}

func partTwo(contents []string) {

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   21,
		Part:  2,
		Value: computeTotal(contents, 25),
	})
}

func computeTotal(contents []string, numRobots int) int {
	dirSequences := computeSequences(DIR_KEYPAD)
	numSequences := computeSequences(NUM_KEYPAD)

	// For the top-level, we pre-compute te lengths here as we'll use them later when we look at depths
	dirLengths := make(map[[2]string]int)
	for key, value := range dirSequences {
		dirLengths[key] = len(value[0])
	}

	cacheMap := make(map[string]int)
	total := 0

	for _, line := range contents {
		inputs := solve(line, &numSequences)
		minLength := math.MaxInt
		for _, input := range inputs {
			length := computeLength(input, numRobots, &dirSequences, &dirLengths, &cacheMap)
			if length < minLength {
				minLength = length
			}
		}
		num, _ := strconv.Atoi(line[:len(line)-1])
		total += minLength * num
	}

	return total
}

func solve(input string, sequences *map[[2]string][]string) []string {
	options := [][]string{}
	for i := 0; i < len(input); i++ {
		startChar := "A"
		if i > 0 {
			startChar = string(input[i-1])
		}
		endChar := string(input[i])
		options = append(options, (*sequences)[[2]string{startChar, endChar}])
	}

	var combine func([][]string) []string
	combine = func(arrays [][]string) []string {
		if len(arrays) == 0 {
			return []string{""}
		}
		first, rest := arrays[0], combine(arrays[1:])
		result := []string{}
		for _, f := range first {
			for _, r := range rest {
				result = append(result, f+r)
			}
		}
		return result
	}

	return combine(options)
}

// Length of certain sequence between 2 points, at a certain depth
func computeLength(seq string, depth int, dirSeqs *map[[2]string][]string, dirLengths *map[[2]string]int, cacheMap *map[string]int) int {
	if depth == 1 {
		length := 0
		for i := 0; i < len(seq); i++ {
			x := "A"
			if i > 0 {
				x = string(seq[i-1])
			}
			y := string(seq[i])
			length += (*dirLengths)[[2]string{x, y}]
		}
		return length
	}

	key := fmt.Sprintf("%s_%d", seq, depth)
	if val, exists := (*cacheMap)[key]; exists {
		return val
	}

	length := 0
	for i := 0; i < len(seq); i++ {
		x := "A"
		if i > 0 {
			x = string(seq[i-1])
		}
		y := string(seq[i])

		minLength := math.MaxInt
		for _, subSequence := range (*dirSeqs)[[2]string{x, y}] {
			subSequenceLength := computeLength(subSequence, depth-1, dirSeqs, dirLengths, cacheMap)
			if subSequenceLength < minLength {
				minLength = subSequenceLength
			}
		}
		length += minLength
	}

	(*cacheMap)[key] = length
	return length
}

// Build map of getting from A to B for every A,B on the keypad;
func computeSequences(keypad [][3]string) map[[2]string][]string {
	// First convert to i,j coords (row,col)
	keypadToCoord := make(map[string][2]int, 0)
	for i := 0; i < len(keypad); i++ {
		for j := 0; j < len(keypad[j]); j++ {
			// Can't ever be none, so skip these!
			if keypad[i][j] == "N" {
				continue
			}
			keypadToCoord[keypad[i][j]] = [2]int{i, j}
		}
	}

	// Now let's build the sequences for getting from A to B, could b multiple. Use BFS
	return getAllSequencesFromAtoB(keypad, keypadToCoord)
}

type queueStruct struct {
	pos             [2]int
	currentSequence string
}

var DIRECTIONS = [4][2]int{
	{-1, 0},
	{0, 1},
	{1, 0},
	{0, -1},
}

var DIRECTIONS_BY_ARROW = map[int]string{
	0: "^",
	1: ">",
	2: "v",
	3: "<",
}

func getAllSequencesFromAtoB(keypad [][3]string, keypadToCoord map[string][2]int) map[[2]string][]string {
	sequences := make(map[[2]string][]string, 0)

	for charStart := range keypadToCoord {
		for charTarget := range keypadToCoord {
			if charStart == charTarget {
				sequences[[2]string{charStart, charTarget}] = []string{"A"}
				continue
			}

			// BFS;
			sequences[[2]string{charStart, charTarget}] = bfs(keypad, charStart, charTarget, keypadToCoord)

		}

	}

	return sequences
}

// Get the possibilities of all paths from 1 char to the other
func bfs(keypad [][3]string, charStart string, charTarget string, keypadToCoords map[string][2]int) []string {
	queue := make([]queueStruct, 0)
	queue = append(queue, queueStruct{keypadToCoords[charStart], ""})
	var element queueStruct
	possibilities := make([]string, 0)
	optimalLength := math.MaxInt64

	for {
		if len(queue) == 0 {
			break
		}

		// Grab the next element in queue
		element, queue = queue[0], queue[1:]

		for dirIndex, dir := range DIRECTIONS {
			newI := element.pos[0] + dir[0]
			newJ := element.pos[1] + dir[1]

			// Range check;
			if newI < 0 || newJ < 0 || newI > len(keypad)-1 || newJ > len(keypad[0])-1 {
				continue
			}

			// Can't be Null;
			if keypad[newI][newJ] == "N" {
				continue
			}

			if keypad[newI][newJ] == charTarget {
				if len(element.currentSequence) > optimalLength {
					return possibilities
				}

				// Otherwise we'rw at the optimal!
				optimalLength = len(element.currentSequence) + 1
				newString := fmt.Sprintf("%s%s%s", element.currentSequence, string(DIRECTIONS_BY_ARROW[dirIndex]), "A")
				possibilities = append(possibilities, newString)
				continue
			}

			// Otherwise, append to the queue
			queue = append(queue, queueStruct{[2]int{newI, newJ}, element.currentSequence + string(DIRECTIONS_BY_ARROW[dirIndex])})
		}
	}

	return possibilities
}
