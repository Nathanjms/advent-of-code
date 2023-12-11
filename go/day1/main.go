package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type Output struct {
	day   int
	part  int
	value any
}

var inputPath = "./day1/example-input1"
var contents string
var contentsArray []string

func init() {
	// Import and read input data from ./example-input1
	input, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	contents = strings.TrimSpace(string(input))
	contentsArray = strings.Split(contents, "\n")
}

func main() {
	partOne()
}

func partOne() {
	// For each line, get the digits from the left and right, then add them together
	total := 0
	for _, line := range contentsArray {
		total += getDigitsFromLine(line)
	}

	output := Output{
		day:   1,
		part:  1,
		value: total,
	}

	fmt.Println(output)
	fmt.Printf("% +v\n", output)

}

func getDigitsFromLine(line string) int {
	leftIndex := 0
	rightIndex := len(line) - 1
	leftDigit := -1
	rightDigit := -1

	// Go from the left until the first digit is found, this digit could be zero so we can't just check for a non-zero digit
	for leftDigit == -1 {
		// If this is a digit, set leftDigit to the value of this digit
		if unicode.IsDigit(rune(line[leftIndex])) {
			leftDigit, _ = strconv.Atoi(string(line[leftIndex]))
			break
		}
		leftIndex++
	}

	for rightDigit == -1 {
		// If this is a digit, set leftDigit to the value of this digit
		if unicode.IsDigit(rune(line[rightIndex])) {
			rightDigit, _ = strconv.Atoi(string(line[rightIndex]))
			break
		}
		rightIndex--
	}

	combinedInt, _ := strconv.Atoi(strconv.Itoa(leftDigit) + strconv.Itoa(rightDigit))

	fmt.Println(combinedInt)
	// return the string concatenation of the two digits
	return combinedInt
}
