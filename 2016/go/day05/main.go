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

func partOne(contents string) {
	password := ""
	currNum := 0
	for {
		md5Hash := GetMD5Hash(fmt.Sprintf("%s%d", contents, currNum))
		if md5Hash[:5] == "00000" {
			password += string(md5Hash[5])
		}

		currNum++

		if len(password) >= 8 {
			break
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   5,
		Part:  1,
		Value: password,
	})
}

func partTwo(contents string) {
	var passwordArray [8]byte
	validCount := 0
	currNum := 0
	isValid := false
	for {
		md5Hash := GetMD5Hash(fmt.Sprintf("%s%d", contents, currNum))
		// fmt.Println(md5Hash)
		if md5Hash[:5] == "00000" {
			// We use the 6-1th character for position, and 7-1th for value
			index, err := strconv.Atoi(string(md5Hash[5]))
			if err != nil || index > len(passwordArray)-1 || passwordArray[index] != 0 {
				// fmt.Println(index, err)
				// Ignore if error in conversion, out of range, or already been input
			} else {
				passwordArray[index] += md5Hash[6]
				validCount++
				if validCount == 8 {
					isValid = true
					break
				}
			}
		}

		if isValid {
			break
		}

		currNum++
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   5,
		Part:  2,
		Value: string(passwordArray[:]),
	})
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
