package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
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

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

type room struct {
	name     string
	sectorId int
	checksum string
}

func partOne(contents []string) {
	validCodeIds := 0
	for _, line := range contents {
		tmp := strings.Split(line, "[")
		// name := tmp[0]
		checksum := tmp[1]
		checksum = checksum[:len(checksum)-1] // Remove the ]

		tmp = strings.Split(tmp[0], "-")

		charPerFrequency := make(map[rune]int)

		// var builder strings.Builder
		var sectorId int
		for i, chars := range tmp {
			if i == len(tmp)-1 {
				sectorId, _ = strconv.Atoi(chars)
			} else {
				for _, char := range chars {
					if _, ok := charPerFrequency[char]; !ok {
						charPerFrequency[char] = 1
					} else {
						charPerFrequency[char]++
					}
				}
				// builder.WriteString(chars)
			}
		}
		// letters := builder.String()

		// roomObj := room{
		// 	name:     name,
		// 	sectorId: sectorId,
		// 	checksum: checksum,
		// }

		// To sort, we need an int and not

		keys := make([]rune, 0, len(charPerFrequency))

		for key := range charPerFrequency {
			keys = append(keys, key)
		}

		// First sort by letter order, from a...z
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})

		// Then sort by frequency. DOing it this way means we handle the 'equal' frequency letters too
		sort.SliceStable(keys, func(i, j int) bool {
			return charPerFrequency[keys[i]] > charPerFrequency[keys[j]]
		})

		charVals := make([]rune, 0)

		for i, key := range keys {
			if i > 4 {
				break
			}
			charVals = append(charVals, key)
		}

		invalid := false
		for i, v := range []byte(checksum) {
			if keys[i] != rune(v) {
				invalid = true
				break
			}
		}

		if !invalid {
			validCodeIds += sectorId
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   4,
		Part:  1,
		Value: validCodeIds,
	})
}

func partTwo(contents []string) {
	validSectorId := 0
	for _, line := range contents {
		tmp := strings.Split(line, "[")
		name := tmp[0]
		checksum := tmp[1]
		checksum = checksum[:len(checksum)-1] // Remove the ]

		tmp = strings.Split(tmp[0], "-")

		charPerFrequency := make(map[rune]int)

		var sectorId int
		for i, chars := range tmp {
			if i == len(tmp)-1 {
				sectorId, _ = strconv.Atoi(chars)
			} else {
				for _, char := range chars {
					if _, ok := charPerFrequency[char]; !ok {
						charPerFrequency[char] = 1
					} else {
						charPerFrequency[char]++
					}
				}
			}
		}

		// To sort, we need an int and not

		keys := make([]rune, 0, len(charPerFrequency))

		for key := range charPerFrequency {
			keys = append(keys, key)
		}

		// First sort by letter order, from a...z
		sort.SliceStable(keys, func(i, j int) bool {
			return keys[i] < keys[j]
		})

		// Then sort by frequency. DOing it this way means we handle the 'equal' frequency letters too
		sort.SliceStable(keys, func(i, j int) bool {
			return charPerFrequency[keys[i]] > charPerFrequency[keys[j]]
		})

		invalid := false
		for i, v := range []byte(checksum) {
			if keys[i] != rune(v) {
				invalid = true
				break
			}
		}

		if !invalid {
			var builder strings.Builder

			for i, char := range []byte(name) {
				// Sector ID is always 3 digits, so leave out the last 4 digits (sector id plus the dash)
				if i < (len(name) - 4) {
					// Replace dashes with spaces
					if char == '-' {
						builder.WriteString(" ")
					} else {
						// a is 97 as a byte, so we need to offset by 97 when doing the modulo of 26 (the alphabet length), then add it back
						char = 97 + (char-97+byte(sectorId%26))%26
						builder.WriteByte(char)
					}
				}
			}
			if builder.String() == "northpole object storage" {
				validSectorId = sectorId
			}
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   4,
		Part:  2,
		Value: validSectorId,
	})
}
