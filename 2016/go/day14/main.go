package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"crypto/md5"
	"encoding/hex"
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
		hash := getMD5Hash(contents + strconv.Itoa(index))
		if charToMatch := threeInRow(hash); charToMatch != 0 {
			if checkIfValid(contents, index, charToMatch, false) {
				validIndexes[index] = true
			}

		}

		if len(validIndexes) >= 64 {
			break
		}
		index++
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   14,
		Part:  1,
		Value: index,
	})
}

var hashMapCachePtTwo = make(map[string]string) // the result of hashing the given key 1 + 2016 times

func partTwo(contents string) {
	index := 0
	validIndexes = make(map[int]bool)

	for {
		hash := getMD5HashPt2(contents + strconv.Itoa(index))
		if charToMatch := threeInRow(hash); charToMatch != 0 {
			if checkIfValid(contents, index, charToMatch, true) {
				validIndexes[index] = true
			}

		}

		if len(validIndexes) >= 64 {
			break
		}
		index++
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   14,
		Part:  2,
		Value: index,
	})
}

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func getMD5HashPt2(text string) string {
	originalText := text
	if hash, ok := hashMapCachePtTwo[originalText]; ok {
		return hash
	}

	// Hash it, then hash again 2016 times... 2017 hashes in total

	for i := 0; i < 2017; i++ {
		hash := md5.Sum([]byte(text))
		text = hex.EncodeToString(hash[:])
	}

	hashMapCachePtTwo[originalText] = text
	return text
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
		if input[i] == charToMatch && input[i] == input[i+1] && input[i+1] == input[i+2] && input[i+2] == input[i+3] && input[i+3] == input[i+4] {
			return true
		}
	}
	return false
}

func checkIfValid(input string, index int, charToMatch byte, isPartTwo bool) bool {
	// Check if any of the next 1000 hashes contain 5 characters in a row
	for i := 0; i < 1000; i++ {
		var hash string
		if isPartTwo {
			hash = getMD5HashPt2(input + strconv.Itoa(index+i+1))
		} else {
			hash = getMD5Hash(input + strconv.Itoa(index+i+1))
		}
		if fiveInRow(hash, charToMatch) {
			return true
		}
	}

	return false
}
