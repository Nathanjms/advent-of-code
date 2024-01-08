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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents string) {
	for {
		fmt.Println(contents)
		if satisfiesRequirements(contents) {
			break
		}

		// Increment the password
		contents = incrementString(contents)
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  1,
		Value: contents,
	})
}

func partTwo(contents string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   11,
		Part:  2,
		Value: "TODO",
	})
}

func incrementString(s string) string {
	runes := []rune(s)
	length := len(runes)

	// Start from the rightmost character and increment
	for i := length - 1; i >= 0; i-- {
		if runes[i] < 'z' {
			runes[i]++
			// We can skip if the character is i, o, or l
			if runes[i] == 'i' || runes[i] == 'o' || runes[i] == 'l' {
				runes[i]++
			}
			return string(runes)
		} else {
			runes[i] = 'a'
		}
	}

	// If all characters are 'z', add a new 'a' at the beginning
	return "a" + string(runes)
}

/*
* Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot
skip letters; abd doesn't count.
* Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore
confusing.
* Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.
*/
func satisfiesRequirements(contents string) bool {
	// Check for increasing straight of at least three letters
	hasThreeInARow := false
	for i := 0; i < len(contents)-2; i++ {
		if contents[i]+1 == contents[i+1] && contents[i+1]+1 == contents[i+2] {
			hasThreeInARow = true
		}
	}
	if !hasThreeInARow {
		return false
	}

	// Check for i, o, or l
	for _, ch := range []byte{'i', 'o', 'l'} {
		if contents[0] == ch || contents[1] == ch || contents[2] == ch {
			return false
		}
	}

	// We need to have 2 distinct pairs of letters
	pairs := 0
	pairsUsed := make(map[string]bool)
	for i := 0; i < len(contents)-1; i++ {
		if contents[i] == contents[i+1] {
			// Check if we've already used this pair
			if pairsUsed[string(contents[i])+string(contents[i+1])] {
				continue
			}
			pairsUsed[string(contents[i])+string(contents[i+1])] = true
			pairs++
			i++ // Skip the next character
		}
	}
	return pairs >= 2
}
