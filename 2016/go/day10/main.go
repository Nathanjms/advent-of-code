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
	usedUpIndexes := make(map[int]bool)
	currentIndex := 0
	robotParts := make(map[int][2]int)
	var robotWithValidParts int
	validParts := []int{17, 61}
	if isUsingExample {
		validParts = []int{2, 5}
	}

	for {
		if _, ok := usedUpIndexes[currentIndex]; ok {
			currentIndex++
			continue // already done this line previously
		}
		line := contents[currentIndex]
		if line[:5] == "value" {
			var partNumber, botNumber int
			fmt.Sscanf(line, "value %d goes to bot %d", &partNumber, &botNumber)
			givePartFromValue(botNumber, &robotParts, partNumber)
			if robotParts[botNumber][0] == validParts[0] && robotParts[botNumber][1] == validParts[1] {
				robotWithValidParts = botNumber
				break
			}
			usedUpIndexes[currentIndex] = true
		} else {
			var fromBot, lowToBot, highToBot int
			var lowType, highType string
			fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &fromBot, &lowType, &lowToBot, &highType, &highToBot)

			if robotParts[fromBot][0] != 0 && robotParts[fromBot][1] != 0 {
				usedUpIndexes[currentIndex] = true
				// Now give to the bots/outputs. We may need to consider that new robot now being ready to go (ie. recursion). But for now lets not
				if lowType == "bot" {
					givePartFromValue(lowToBot, &robotParts, robotParts[fromBot][0])
					if robotParts[lowToBot][0] == validParts[0] && robotParts[lowToBot][1] == validParts[1] {
						robotWithValidParts = lowToBot
						break
					}
				}
				if highType == "bot" {
					givePartFromValue(highToBot, &robotParts, robotParts[fromBot][1])
					if robotParts[highToBot][0] == validParts[0] && robotParts[highToBot][1] == validParts[1] {
						robotWithValidParts = highToBot
						break
					}
				}
				robotParts[fromBot] = [2]int{} // Used up so now empty - may want to use make here?
			}
		}

		currentIndex = (currentIndex + 1) % len(contents)
		if len(usedUpIndexes) == len(contents) {
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   10,
		Part:  1,
		Value: robotWithValidParts,
	})
}

func partTwo(contents []string) {
	usedUpIndexes := make(map[int]bool)
	currentIndex := 0
	robotParts := make(map[int][2]int)
	binParts := make(map[int]int)

	for {
		if _, ok := usedUpIndexes[currentIndex]; ok {
			currentIndex = (currentIndex + 1) % len(contents) // Go to next val, or reset to 0
			continue                                          // already done this line previously
		}
		line := contents[currentIndex]
		if line[:5] == "value" {
			var partNumber, botNumber int
			fmt.Sscanf(line, "value %d goes to bot %d", &partNumber, &botNumber)
			givePartFromValue(botNumber, &robotParts, partNumber)
			usedUpIndexes[currentIndex] = true
		} else {
			var fromBot, lowToBot, highToBot int
			var lowType, highType string
			fmt.Sscanf(line, "bot %d gives low to %s %d and high to %s %d", &fromBot, &lowType, &lowToBot, &highType, &highToBot)

			if robotParts[fromBot][0] != 0 && robotParts[fromBot][1] != 0 {
				usedUpIndexes[currentIndex] = true
				if lowType == "bot" {
					givePartFromValue(lowToBot, &robotParts, robotParts[fromBot][0])
				} else {
					givePartToBin(&binParts, lowToBot, robotParts[fromBot][0])
				}
				if highType == "bot" {
					givePartFromValue(highToBot, &robotParts, robotParts[fromBot][1])
				} else {
					givePartToBin(&binParts, highToBot, robotParts[fromBot][1])
				}
				robotParts[fromBot] = [2]int{} // Used up so now empty - may want to use make here?
			}
		}
		currentIndex = (currentIndex + 1) % len(contents) // Go to next val, or reset to 0
		if len(usedUpIndexes) == len(contents) {
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   10,
		Part:  2,
		Value: binParts[0] * binParts[1] * binParts[2],
	})
}

func givePartFromValue(robotIndex int, robotParts *map[int][2]int, partValue int) {
	if _, ok := (*robotParts)[robotIndex]; !ok {
		(*robotParts)[robotIndex] = [2]int{partValue}
	} else {
		// Assume we can only get here if the robot has only 1 part!
		if (*robotParts)[robotIndex][0] != 0 && (*robotParts)[robotIndex][1] != 0 {
			panic("Part number malfunction - already has 2 parts?")
		}
		// Compare to the 1st part, and determine whether to prepend or append
		if partValue >= (*robotParts)[robotIndex][0] {
			(*robotParts)[robotIndex] = [2]int{(*robotParts)[robotIndex][0], partValue}
		} else {
			(*robotParts)[robotIndex] = [2]int{partValue, (*robotParts)[robotIndex][0]}
		}
	}
}

func givePartToBin(binParts *map[int]int, binDex int, partValue int) {
	if _, ok := (*binParts)[binDex]; ok {
		panic("Already have value in bin :0")
	}
	(*binParts)[binDex] = partValue
}
