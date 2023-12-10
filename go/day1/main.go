package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

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
	// If the digits are the same, add them to the total
	total := 0
	for _, line := range contentsArray {
		digits := getDigitsFromLine(line)
		total += digits
	}

	fmt.Println(total)

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
			fmt.Println(line[leftIndex])
			leftDigit = int(line[leftIndex])
			break
		}
		leftIndex++
	}

	for rightDigit == -1 {
		// If this is a digit, set leftDigit to the value of this digit
		if unicode.IsDigit(rune(line[rightIndex])) {
			fmt.Println(rune(line[rightIndex]))
			rightDigit = int(line[rightIndex])
			break
		}
		rightIndex--
	}

	fmt.Println(leftDigit, rightDigit)

	combinedInt, _ := strconv.Atoi(strconv.Itoa(leftDigit) + strconv.Itoa(rightDigit))
	// return the string concatenation of the two digits
	return combinedInt
}
