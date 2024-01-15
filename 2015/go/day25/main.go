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
	var rowLocation, colLocation int
	_, err := fmt.Sscanf(
		contents,
		"To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.", &rowLocation, &colLocation,
	)
	if err != nil {
		panic("Error :c")
	}

	/* Note that we'll be offset by 1! */
	// We probably don't need a huge matrix, and can just store the previous value?
	value := 20151125
	rowIdx := 1 // we have the start value, so start from '[1][0]'
	colIdx := 0
	maxRow := 0

	for {
		maxRow = max(rowIdx, maxRow)
		// Get the value for this location
		value = (value * 252533) % 33554393
		// Are we where we need to be (minus 1 to account for 0 indexing)?
		if rowIdx == rowLocation-1 && colIdx == colLocation-1 {
			break
		}
		// Iterate the relevant row/col
		rowIdx-- // Row goes up one
		colIdx++ // Col goes right one
		if rowIdx < 0 {
			// If we're now below zero, we start the next diagonal
			colIdx = 0          // This goes back to the start
			rowIdx = maxRow + 1 // This goes to one more than the previously largest row
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   25,
		Part:  1,
		Value: value,
	})
}

func partTwo(contents string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   25,
		Part:  2,
		Value: "Merry Christmas!",
	})
}
