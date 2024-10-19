package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
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

	var _, contents = sharedcode.ParseFile(inputPath)

	partOne(contents)
	partTwo(contents)
}

func partOne(contents []string) {
	var password []byte
	if isUsingExample {
		password = []byte("abcde")
	} else {
		password = []byte("abcdefgh")
	}

	password = doInstructions(contents, password)

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   21,
		Part:  1,
		Value: string(password),
	})
}

func partTwo(contents []string) {
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   21,
		Part:  2,
		Value: "TODO",
	})
}

func doInstructions(instructions []string, password []byte) []byte {
	for _, instruction := range instructions {
		// Does it start with 'swap position'
		if instruction[0:6] == "swap p" {
			indexOne, _ := strconv.Atoi(string(instruction[14]))
			indexTwo, _ := strconv.Atoi(string(instruction[len(instruction)-1]))

			password[indexOne], password[indexTwo] = password[indexTwo], password[indexOne]

			continue
		}

		// Does it start with 'swap letter'
		if instruction[0:6] == "swap l" {
			indexOne := getIndexOfLetter(instruction[12], password)
			indexTwo := getIndexOfLetter(instruction[len(instruction)-1], password)

			password[indexOne], password[indexTwo] = password[indexTwo], password[indexOne]

			continue
		}

		// Does it start with 'rotate based'
		if instruction[0:8] == "rotate b" {
			// rotate based on position of letter X means that the whole string should be rotated to the right based on the index of letter X (counting from 0) as determined before this instruction does any rotations. Once the index is determined, rotate the string to the right one time, plus a number of times equal to that index, plus one additional time if the index was at least 4.
			index := getIndexOfLetter(instruction[len(instruction)-1], password)

			rotations := index + 1

			if index >= 4 {
				rotations++
			}

			rotations = rotations % len(password)

			reversalAmount := len(password) - rotations

			tempPassword := make([]byte, len(password))

			for i := 0; i < len(password); i++ {
				tempPassword[i] = password[(i+reversalAmount)%len(password)]
			}

			password = tempPassword

			continue
		}

		// Does it start with 'rotate'
		if instruction[0:6] == "rotate" {
			var numSteps int
			var moveRightAmount int
			tempPassword := make([]byte, len(password))
			if instruction[7] == 'l' {
				numSteps, _ = strconv.Atoi(string(instruction[12]))
				moveRightAmount = len(password) - numSteps
			} else if instruction[7] == 'r' {
				// rotate right
				moveRightAmount, _ = strconv.Atoi(string(instruction[13]))
			}

			reversalAmount := len(password) - moveRightAmount
			for i := 0; i < len(password); i++ {
				tempPassword[i] = password[(i+reversalAmount)%len(password)]
			}

			password = tempPassword

			continue
		}

		// Does it start with 'reverse'
		if instruction[0:7] == "reverse" {
			indexOne, _ := strconv.Atoi(string(instruction[18]))
			indexTwo, _ := strconv.Atoi(string(instruction[len(instruction)-1]))

			diff := indexTwo - indexOne

			for i := 0; i <= diff/2; i++ {
				password[indexOne+i], password[indexTwo-i] = password[indexTwo-i], password[indexOne+i]
			}

			continue
		}

		// Does it start with 'move'
		if instruction[0:4] == "move" {
			// move position 1 to position 2
			indexOne, _ := strconv.Atoi(string(instruction[14]))
			indexTwo, _ := strconv.Atoi(string(instruction[len(instruction)-1]))

			passwordTemp := make([]byte, len(password))

			hasFlipped := false
			// Shortcut: if indexOne is RHS of indexTwo, flip password and then flip back at end:
			if indexOne > indexTwo {
				hasFlipped = true
				password = flipPassword(password)

				indexOne = len(password) - 1 - indexOne
				indexTwo = len(password) - 1 - indexTwo
			}

			for i := 0; i < len(password); i++ {
				// Everything before it remains as it is
				if i < indexOne {
					passwordTemp[i] = password[i]
					continue
				}

				if i < indexTwo {
					passwordTemp[i] = password[i+1]
					continue
				}

				if i == indexTwo {
					passwordTemp[i] = password[indexOne]
					continue
				}

				if i > indexTwo {
					passwordTemp[i] = password[i]
					continue
				}
			}

			if hasFlipped {
				password = flipPassword(passwordTemp)
			} else {
				password = passwordTemp
			}

			continue
		}
	}

	return password
}

func getIndexOfLetter(letter byte, password []byte) int {
	for i, char := range password {
		if char == letter {
			return i
		}
	}
	return -1
}

func flipPassword(password []byte) []byte {
	flipped := make([]byte, len(password))
	for i := 0; i < len(password); i++ {
		flipped[i] = password[len(password)-i-1]
	}
	return flipped
}
