package main

import (
	"aoc-shared/pkg/sharedcode"
	"aoc-shared/pkg/sharedstruct"
	"math"
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
	registers, instructions := parseInput(contents)

	instructionPointer := 0
	output := make([]int64, 0)

	// Go through the instructions;
	for {
		if instructionPointer > len(instructions)-1 {
			break
		}

		handleInstruction(&registers, instructions[instructionPointer], instructions[instructionPointer+1], &instructionPointer, &output)
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   17,
		Part:  1,
		Value: sliceToString(output),
	})

}

func partTwo(contents []string) {
	originalRegisters, instructions := parseInput(contents)
	instructionsAsString := sliceToString(instructions)
	hasFoundA := false

	A := int64(117440)
	iterations := int64(0)
	// Trial different A values;
	for {
		if iterations > 0 {
			break
		}
		if hasFoundA {
			break
		}
		instructionPointer := 0
		output := make([]int64, 0)
		registers := make(map[byte]int64)
		registers['A'] = A
		registers['B'] = originalRegisters['B']
		registers['C'] = originalRegisters['C']
		for {
			if instructionPointer > len(instructions)-1 {
				break
			}

			handleInstruction(&registers, instructions[instructionPointer], instructions[instructionPointer+1], &instructionPointer, &output)
		}

		if instructionsAsString == sliceToString(output) {
			hasFoundA = true
			break
		}

		// fmt.Println(A, instructionsAsString, sliceToString(output))

		A++
		iterations++
	}

	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   17,
		Part:  2,
		Value: A,
	})
}

func sliceToString(slice []int64) string {
	var builder strings.Builder

	for i, outputVal := range slice {
		stringVal := strconv.FormatInt(outputVal, 10)
		builder.WriteString(stringVal)
		if i != len(slice)-1 {
			builder.WriteByte(',')
		}
	}

	return builder.String()
}

func handleInstruction(registers *map[byte]int64, opcode int64, operand int64, pointer *int, output *[]int64) {
	if opcode == 0 {
		// (*registers)['A'] = (*registers)['A'] >> getComboOperand(operand, registers) // Apparently this is the same!
		numerator := (*registers)['A']
		denominator := math.Pow(2, float64(getComboOperand(operand, registers)))

		result := math.Floor(float64(numerator) / float64(denominator))
		(*registers)['A'] = int64(result)
		*pointer = *pointer + 2
		return
	}
	if opcode == 1 {
		result := (*registers)['B'] ^ operand
		(*registers)['B'] = result
		*pointer = *pointer + 2
		return
	}
	if opcode == 2 {
		(*registers)['B'] = getComboOperand(operand, registers) % 8
		*pointer = *pointer + 2
		return
	}
	if opcode == 3 {
		if (*registers)['A'] == 0 {
			*pointer = *pointer + 2
			return
		}
		*pointer = int(operand)
		return
	}
	if opcode == 4 {
		result := (*registers)['B'] ^ (*registers)['C']
		(*registers)['B'] = result
		*pointer = *pointer + 2
		return
	}
	if opcode == 5 {
		outputVal := getComboOperand(operand, registers) % 8
		*pointer = *pointer + 2
		*output = append(*output, outputVal)
		return
	}
	if opcode == 6 {
		numerator := (*registers)['A']
		denominator := math.Pow(2, float64(getComboOperand(operand, registers)))

		result := math.Floor(float64(numerator) / float64(denominator))
		(*registers)['B'] = int64(result)
		*pointer = *pointer + 2
		return
	}
	if opcode == 7 {
		numerator := (*registers)['A']
		denominator := math.Pow(2, float64(getComboOperand(operand, registers)))

		result := math.Floor(float64(numerator) / float64(denominator))
		(*registers)['C'] = int64(result)
		*pointer = *pointer + 2
		return
	} else {
		panic("Invalid opcode?")
	}
}

func getComboOperand(operand int64, register *map[byte]int64) int64 {
	if operand <= 3 {
		return operand
	}

	if operand == 4 {
		return (*register)['A']
	}

	if operand == 5 {
		return (*register)['B']
	}

	if operand == 6 {
		return (*register)['C']
	}

	panic("Invalid!")
}

func parseInput(contents []string) (map[byte]int64, []int64) {
	registers := make(map[byte]int64)
	instructions := make([]int64, 0)

	for _, line := range contents {
		if len(line) == 0 {
			continue
		}

		if line[:8] == "Register" {
			register := line[9]
			val := line[12:]

			intVal, _ := strconv.ParseInt(val, 10, 64)

			registers[register] = intVal
		}

		if line[:7] == "Program" {
			stringInstructions := strings.Split(line[9:], ",")
			for i := range stringInstructions {
				intVal, _ := strconv.ParseInt(stringInstructions[i], 10, 64)
				instructions = append(instructions, intVal)
			}
		}
	}

	return registers, instructions
}
