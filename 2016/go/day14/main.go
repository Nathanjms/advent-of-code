package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"crypto/md5"
	"encoding/hex"
	"fmt"
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

var validIndexes = make(map[int]bool)

func partOne(contents string) {
	index := 0

	for {
		hash := getmd5hash(contents + strconv.Itoa(index))
		if charToMatch := threeInRow(hash), charToMatch != 0 {
			// Send of goroutine to check for five in a row, we dont need to stop the main process
			checkIfValid(contents, index, )
		}

		if len(validIndexes) >= 64 {
			break
		}
		index++
	}

	fmt.Println(validIndexes)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   14,
		Part:  1,
		Value: "TODO",
	})
}

func partTwo(contents string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   14,
		Part:  2,
		Value: "TODO",
	})
}

func getmd5hash(input string) string {
	hash := md5.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}

func threeInRow(input string) byte {
	// Does the string contain 3 characters in a row?
	for i := 0; i < len(input)-2; i++ {
		if input[i] == input[i+1] && input[i+1] == input[i+2] {
			return input[i]
		}
	}
	return 0
}

func fiveInRow(input string, charToMatch byte) bool {
	// Does the string contain 5 characters in a row?
	for i := 0; i < len(input)-4; i++ {
		if input[i] == input[i+1] && input[i+1] == input[i+2] && input[i+2] == input[i+3] && input[i+3] == input[i+4] && input[i] == charToMatch {
			return true
		}
	}
	return false
}

func checkIfValid(input string, index int, charToMatch byte) {
	// Check if any of the next 1000 hashes contain 5 characters in a row
	for i := 0; i < 1000; i++ {
		hash := getmd5hash(input + strconv.Itoa(index+i+1))
		if fiveInRow(hash, charToMatch) {
			validIndexes[index] = true
			return
		}
	}
}
