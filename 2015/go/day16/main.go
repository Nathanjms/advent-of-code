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

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

type InnerMap struct {
	Value    int    `json:"value"`
	Operator string `json:"operator"`
}

var targetsByKey = map[string]InnerMap{
	"children":    {Value: 3, Operator: "eq"},
	"cats":        {Value: 7, Operator: "eq"},
	"samoyeds":    {Value: 2, Operator: "eq"},
	"pomeranians": {Value: 3, Operator: "eq"},
	"akitas":      {Value: 0, Operator: "eq"},
	"vizslas":     {Value: 0, Operator: "eq"},
	"goldfish":    {Value: 5, Operator: "eq"},
	"trees":       {Value: 3, Operator: "eq"},
	"cars":        {Value: 2, Operator: "eq"},
	"perfumes":    {Value: 1, Operator: "eq"},
}

func partOne(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   16,
		Part:  1,
		Value: getCorrectSue(contents, targetsByKey),
	})
}

func partTwo(contents []string) {
	partTwoTargetsByKey := make(map[string]InnerMap)
	for k, v := range targetsByKey {
		switch k {
		case "cats", "trees":
			partTwoTargetsByKey[k] = InnerMap{Value: v.Value, Operator: "gt"}
		case "goldfish", "pomeranian":
			partTwoTargetsByKey[k] = InnerMap{Value: v.Value, Operator: "lt"}
		default:
			partTwoTargetsByKey[k] = InnerMap{Value: v.Value, Operator: "eq"}
		}
	}
	fmt.Println(partTwoTargetsByKey)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   16,
		Part:  2,
		Value: getCorrectSue(contents, partTwoTargetsByKey),
	})
}

func getCorrectSue(contents []string, targetsByKey map[string]InnerMap) int {
	giftSueNum := 0

	for _, line := range contents {
		var sueNum, attOneVal, attTwoVal, attThreeVal int
		var attOne, attTwo, attThree string
		_, err := fmt.Sscanf(
			line,
			"Sue %d: %s %d, %s %d, %s %d",
			&sueNum,
			&attOne,
			&attOneVal,
			&attTwo,
			&attTwoVal,
			&attThree,
			&attThreeVal,
		)
		if err != nil {
			panic(err)
		}

		// remove the colons
		attOne = attOne[:len(attOne)-1]
		attTwo = attTwo[:len(attTwo)-1]
		attThree = attThree[:len(attThree)-1]

		// Check all three match, and break once we find a match
		if checkAttribute(attOne, attOneVal, &targetsByKey) &&
			checkAttribute(attTwo, attTwoVal, &targetsByKey) &&
			checkAttribute(attThree, attThreeVal, &targetsByKey) {
			giftSueNum = sueNum
			break
		}
	}

	return giftSueNum
}

func checkAttribute(attrName string, attrVal int, targetsByKey *map[string]InnerMap) bool {
	compareType := (*targetsByKey)[attrName].Operator

	if compareType == "gt" {
		return attrVal > (*targetsByKey)[attrName].Value
	} else if compareType == "lt" {
		return attrVal < (*targetsByKey)[attrName].Value
	} else {
		return attrVal == (*targetsByKey)[attrName].Value
	}
}
