package sharedcode

import (
	"os"
	"strings"
)

/**
 * ParseFile parses a file and returns (1) the entire string and (2) an array of strings, one for each line
 */
func ParseFile(inputPath string) (string, []string) {
	// Import and read input data from ./example-input1
	input, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	contents := strings.TrimSpace(string(input))
	contentsArray := strings.Split(contents, "\n")
	return contents, contentsArray
}
