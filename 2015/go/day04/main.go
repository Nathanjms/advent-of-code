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

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func partOne(contents string) {
	num := 0

	for {
		if GetMD5Hash(fmt.Sprintf("%s%d", contents, num))[:5] == "00000" {
			break
		}
		num++
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   4,
		Part:  1,
		Value: num,
	})
}

func partTwo(contents string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   4,
		Part:  1,
		Value: "TODO",
	})
}
