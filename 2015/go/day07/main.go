package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

func partOne(contents []string) {
	// Mapping of key of a variable, and a value of a 16 bit int:
	var variables = map[string]uint16{}
	// Assume that every RHS variable only comes up ONCE, so we need to repeat until all have values
	for {
		allHaveMatch := true
		for _, line := range contents {
			// Parse line . First split the string around ' -> ':
			var splitLine = strings.Split(line, " -> ")

			// We now have two sides. LHS will either be a single number, or something like 'x LSHIFT y', or 1 num and a var like '1 LSHIFT y'. RHS will be a single variable.
			var lhs = splitLine[0]
			var rhs = splitLine[1]

			// The key is the RHS. If it has already go an entry in variables, we skip!
			_, ok := variables[rhs]
			// If the key exists
			if ok {
				continue // Skip just this line if we've already got a value for it
			}

			/* Now to get the value for the RHS, we need to handle the cases that the LHS can be in: */

			var splitLhs = strings.Split(lhs, " ")

			if len(splitLhs) == 1 {
				// 1. Simple operation assignment
				// Is it an int or a variable?
				if intVal, err := strconv.Atoi(splitLhs[0]); err == nil {
					// Is an int
					variables[rhs] = uint16(intVal)
				} else {
					// Is a variable
					// Ensure variable has map value, else skip
					if lhsVal, ok := variables[splitLhs[0]]; ok {
						// We are good to do the operation
						variables[rhs] = lhsVal
					} else {
						allHaveMatch = false
						continue // Skip if not got a map value yet
					}
				}
			} else if len(splitLhs) == 2 {
				// 2. NOT operation
				// Is it an int or a variable?
				if intVal, err := strconv.Atoi(splitLhs[1]); err == nil {
					// Is an int
					variables[rhs] = ^uint16(intVal)
				} else {
					// Is a variable
					// Ensure variable has map value, else skip
					if lhsVal, ok := variables[splitLhs[1]]; ok {
						// We are good to do the operation
						variables[rhs] = ^lhsVal
					} else {
						allHaveMatch = false
						continue // Skip if not got a map value yet
					}
				}
			} else if len(splitLhs) == 3 {
				// AND, OR, LSHIFT, RSHIFT operation
				var firstOperand = splitLhs[0]
				var operator = splitLhs[1]
				var secondOperand = splitLhs[2]

				var firstOperandAsInt uint16
				var secondOperandAsInt uint16

				// Check if a letter or a number:
				if firstOperandIntVal, err := strconv.Atoi(firstOperand); err == nil {
					firstOperandAsInt = uint16(firstOperandIntVal)
				} else {
					// Not an int - get from map, or skip if not ready
					if firstOperandIntVal, ok := variables[firstOperand]; ok {
						firstOperandAsInt = uint16(firstOperandIntVal)
					} else {
						allHaveMatch = false
						continue
					}

				}
				if secondOperandIntVal, err := strconv.Atoi(secondOperand); err == nil {
					secondOperandAsInt = uint16(secondOperandIntVal)
				} else {
					// Not an int - get from map, or skip if not ready
					if secondOperandIntVal, ok := variables[secondOperand]; ok {
						secondOperandAsInt = uint16(secondOperandIntVal)
					} else {
						allHaveMatch = false
						continue
					}
				}

				switch operator {
				case "AND":
					variables[rhs] = firstOperandAsInt & secondOperandAsInt
				case "OR":
					variables[rhs] = firstOperandAsInt | secondOperandAsInt
				case "LSHIFT":
					variables[rhs] = firstOperandAsInt << secondOperandAsInt
				case "RSHIFT":
					variables[rhs] = firstOperandAsInt >> secondOperandAsInt
				}
			}
		}
		if allHaveMatch {
			// All done!
			break
		}
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   7,
		Part:  1,
		Value: variables["a"],
	})
}

func partTwo(contents []string) {
	// Mapping of key of a variable, and a value of a 16 bit int:
	var variables = map[string]uint16{}
	// Assume that every RHS variable only comes up ONCE, so we need to repeat until all have values
	for {
		allHaveMatch := true
		for _, line := range contents {
			// Parse line . First split the string around ' -> ':
			var splitLine = strings.Split(line, " -> ")

			// We now have two sides. LHS will either be a single number, or something like 'x LSHIFT y', or 1 num and a var like '1 LSHIFT y'. RHS will be a single variable.
			var lhs = splitLine[0]
			var rhs = splitLine[1]

			if rhs == "b" {
				// Override b to the a value from Pt1
				lhs = "16076"
			}

			// The key is the RHS. If it has already go an entry in variables, we skip!
			_, ok := variables[rhs]
			// If the key exists
			if ok {
				continue // Skip just this line if we've already got a value for it
			}

			/* Now to get the value for the RHS, we need to handle the cases that the LHS can be in: */

			var splitLhs = strings.Split(lhs, " ")

			if len(splitLhs) == 1 {
				// 1. Simple operation assignment
				// Is it an int or a variable?
				if intVal, err := strconv.Atoi(splitLhs[0]); err == nil {
					// Is an int
					variables[rhs] = uint16(intVal)
				} else {
					// Is a variable
					// Ensure variable has map value, else skip
					if lhsVal, ok := variables[splitLhs[0]]; ok {
						// We are good to do the operation
						variables[rhs] = lhsVal
					} else {
						allHaveMatch = false
						continue // Skip if not got a map value yet
					}
				}
			} else if len(splitLhs) == 2 {
				// 2. NOT operation
				// Is it an int or a variable?
				if intVal, err := strconv.Atoi(splitLhs[1]); err == nil {
					// Is an int
					variables[rhs] = ^uint16(intVal)
				} else {
					// Is a variable
					// Ensure variable has map value, else skip
					if lhsVal, ok := variables[splitLhs[1]]; ok {
						// We are good to do the operation
						variables[rhs] = ^lhsVal
					} else {
						allHaveMatch = false
						continue // Skip if not got a map value yet
					}
				}
			} else if len(splitLhs) == 3 {
				// AND, OR, LSHIFT, RSHIFT operation
				var firstOperand = splitLhs[0]
				var operator = splitLhs[1]
				var secondOperand = splitLhs[2]

				var firstOperandAsInt uint16
				var secondOperandAsInt uint16

				// Check if a letter or a number:
				if firstOperandIntVal, err := strconv.Atoi(firstOperand); err == nil {
					firstOperandAsInt = uint16(firstOperandIntVal)
				} else {
					// Not an int - get from map, or skip if not ready
					if firstOperandIntVal, ok := variables[firstOperand]; ok {
						firstOperandAsInt = uint16(firstOperandIntVal)
					} else {
						allHaveMatch = false
						continue
					}

				}
				if secondOperandIntVal, err := strconv.Atoi(secondOperand); err == nil {
					secondOperandAsInt = uint16(secondOperandIntVal)
				} else {
					// Not an int - get from map, or skip if not ready
					if secondOperandIntVal, ok := variables[secondOperand]; ok {
						secondOperandAsInt = uint16(secondOperandIntVal)
					} else {
						allHaveMatch = false
						continue
					}
				}

				switch operator {
				case "AND":
					variables[rhs] = firstOperandAsInt & secondOperandAsInt
				case "OR":
					variables[rhs] = firstOperandAsInt | secondOperandAsInt
				case "LSHIFT":
					variables[rhs] = firstOperandAsInt << secondOperandAsInt
				case "RSHIFT":
					variables[rhs] = firstOperandAsInt >> secondOperandAsInt
				}
			}
		}
		if allHaveMatch {
			// All done!
			break
		}
	}
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   7,
		Part:  2,
		Value: variables["a"],
	})
}
