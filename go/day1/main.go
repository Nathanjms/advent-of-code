package day1

import (
	"strconv"
	"unicode"

	"github.com/nathanjms/advent-of-code-2023/go/pkg/sharedcode"
	"github.com/nathanjms/advent-of-code-2023/go/pkg/sharedstruct"
)

var inputPath = "./day1/example-input1"

// var contents string
var contentsArray []string

func init() {
	_, contentsArray = sharedcode.ParseFile(inputPath)
}

func PartOne() {
	// For each line, get the digits from the left and right, then add them together
	total := 0
	for _, line := range contentsArray {
		total += getDigitsFromLine(line)
	}

	output := sharedstruct.Output{
		Day:   1,
		Part:  1,
		Value: total,
	}

	sharedstruct.PrintOutput(output)

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

	// return the string concatenation of the two digits
	return combinedInt
}
