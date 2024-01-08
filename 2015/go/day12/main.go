package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
	// We don't need to parse into json, just use regex to extract all numbers:
	var numbers = regexp.MustCompile(`-?\d+`)
	var matches = numbers.FindAllString(contents, -1)
	sum := 0
	for _, match := range matches {
		intVal, _ := strconv.Atoi(match)
		sum += intVal
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   12,
		Part:  1,
		Value: sum,
	})
}

func partTwo(contents string) {
	// We need to parse it this time, because we need to ignore any object that has a value of "red" anywhere in it.
	// We can't just use regex to extract all numbers, because we need to know the context of the number.
	// Parse the JSON string

	// Parse the JSON string
	data, err := parseJSON(contents)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	count := 0
	var getCount func(interface{})

	getCount = func(data interface{}) {
		switch v := data.(type) {
		case float64:
			count += int(v)
		case map[string]interface{}:
			// If is an object, first check red is not a value in any of it. If it is, return
			for key, value := range v {
				if value == "red" || key == "red" {
					return
				}
			}
			// If not, call getCount on each value
			for _, value := range v {
				getCount(value)
			}
		case []interface{}:
			for _, value := range v {
				getCount(value)
			}
		}
	}

	getCount(data)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   12,
		Part:  2,
		Value: count,
	})
}

func parseJSON(jsonString string) (interface{}, error) {
	var data interface{}
	err := json.Unmarshal([]byte(jsonString), &data)
	return data, err
}
