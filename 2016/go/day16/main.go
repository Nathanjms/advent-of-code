package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"bytes"
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

	var contents, _ = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents string) {
	diskSizeTarget := 272
	if isUsingExample {
		diskSizeTarget = 20
	}

	// First iterate until we reach the disk size target, or larger
	dragonCurve := contents
	for len(dragonCurve) < diskSizeTarget {
		dragonCurve = makeDragonCurve(dragonCurve)
	}

	// Truncate to only the diskSizeTarget characters
	dragonCurve = dragonCurve[:diskSizeTarget]

	// Now we're ready to generate the checksum:
	checksum := generateChecksum([]byte(dragonCurve))
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   16,
		Part:  1,
		Value: string(checksum),
	})
}

func partTwo(contents string) {
	diskSizeTarget := 35651584

	// First iterate until we reach the disk size target, or larger
	dragonCurve := contents
	for len(dragonCurve) < diskSizeTarget {
		dragonCurve = makeDragonCurve(dragonCurve)
	}

	// Truncate to only the diskSizeTarget characters
	dragonCurve = dragonCurve[:diskSizeTarget]

	// Now we're ready to generate the checksum:
	checksum := generateChecksum([]byte(dragonCurve))
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   16,
		Part:  2,
		Value: string(checksum),
	})
}

/*
*
Instructions:

	Call the data you have at this point "a".
	Make a copy of "a"; call this copy "b".
	Reverse the order of the characters in "b".
	In "b", replace all instances of 0 with 1 and all 1s with 0.
	The resulting data is "a", then a single 0, then "b".
*/
func makeDragonCurve(input string) string {
	var builder bytes.Buffer
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == '0' {
			builder.WriteByte('1')
		} else {
			builder.WriteByte('0')
		}
	}
	return input + "0" + builder.String()
}

func generateChecksum(input []byte) []byte {
	var builder bytes.Buffer
	for i := 0; i < len(input); i += 2 {
		if input[i] == input[i+1] {
			builder.WriteByte('1')
		} else {
			builder.WriteByte('0')
		}
	}

	if builder.Len()%2 == 0 {
		return generateChecksum(builder.Bytes())
	}

	return builder.Bytes()
}
