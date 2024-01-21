package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
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

	var _, contents = sharedcode.ParseFile(inputPath)
	contents2 := contents
	if isUsingExample {
		_, contents2 = sharedcode.ParseFile(inputPath + "2")
	}

	partOne(contents)
	partTwo(contents2)
}

type entry struct {
	content          string
	inSquareBrackets bool
}

func partOne(contents []string) {
	numSupported := 0
	for _, line := range contents {
		entries := make([]entry, 0)
		// Build the entries array
		var builder strings.Builder

		for i, char := range line {
			if char == '[' {
				entries = append(entries, entry{builder.String(), false})
				builder.Reset()
			} else if char == ']' {
				entries = append(entries, entry{builder.String(), true})
				builder.Reset()
			} else {
				builder.WriteRune(char)
			}

			// Handle the last line
			if i == len(line)-1 {
				entries = append(entries, entry{builder.String(), false})
				builder.Reset()
			}
		}
		squareBracketsValid := true
		outsideValid := false
		for _, entry := range entries {
			if entry.inSquareBrackets && hasAbba(entry.content) && squareBracketsValid {
				// One inside invalid means the line is invalid from the inside
				squareBracketsValid = false
				break // We now know it can never be valid, so can skip the rest
			} else if !entry.inSquareBrackets && hasAbba(entry.content) && !outsideValid {
				// One outside valid true means the line is valid from the outside
				outsideValid = true
				// We can't skip here because we need to test all remaining square brackets
			}
		}

		if squareBracketsValid && outsideValid {
			numSupported++
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   7,
		Part:  1,
		Value: numSupported,
	})
}

func partTwo(contents []string) {
	numSupported := 0
	for _, line := range contents {
		var builder strings.Builder
		abasInLine := make(map[string]bool)
		babsAsAbasInLine := make(map[string]bool) // log these in aba form so we do not have to convert later!

		for i, char := range line {
			if char == '[' {
				// We've reached the end of the word
				word := builder.String()
				updateAbas(word, &abasInLine) // Add to abas
				builder.Reset()
			} else if char == ']' {
				word := builder.String()
				updateBabs(word, &babsAsAbasInLine)
				builder.Reset()
			} else {
				builder.WriteRune(char)
			}

			// Handle the last line - never in square brackets
			if i == len(line)-1 {
				word := builder.String()
				updateAbas(word, &abasInLine) // Add to abas
				builder.Reset()
			}
		}

		// Now we check if we have a match for the map, and if so it is valid
		valid := false
		for chars := range abasInLine {
			// Is this in the babsAsAbasInLine?
			for babsChars := range babsAsAbasInLine {
				if chars == babsChars {
					valid = true
					break
				}
			}
			if valid {
				break
			}
		}

		if valid {
			numSupported++
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   7,
		Part:  2,
		Value: numSupported,
	})
}

func hasAbba(chars string) bool {
	for i := 0; i < len(chars)-3; i++ {
		if chars[i] == chars[i+3] && chars[i] != chars[i+1] && chars[i+1] == chars[i+2] {
			return true
		}
	}
	return false
}

func updateAbas(chars string, abas *map[string]bool) {
	for i := 0; i < len(chars)-2; i++ {
		if chars[i] == chars[i+2] && chars[i] != chars[i+1] {
			aba := string(chars[i]) + string(chars[i+1]) + string(chars[i+2])
			(*abas)[aba] = true
		}
	}
}

func updateBabs(chars string, babsAsAba *map[string]bool) {
	for i := 0; i < len(chars)-2; i++ {
		if chars[i] == chars[i+2] && chars[i] != chars[i+1] {
			asAba := string(chars[i+1]) + string(chars[i]) + string(chars[i+1])

			(*babsAsAba)[asAba] = true
		}
	}
}
