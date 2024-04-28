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

var INSTRUCTION_TYPES = []string{
	"cpy",
	"inc",
	"dec",
	"jnz",
}

type instruction struct {
	instructionType string
	operand1        string
	operand2        string
}

type state struct {
	index        int
	instructions []instruction
	registers    map[string]int
}

func partOne(contents []string) {
	instructionList := parseInput(contents)
	registers := map[string]int{"a": 0, "b": 0, "c": 0, "d": 0}
	gameState := state{0, instructionList, registers}
	solution := solve(gameState)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   12,
		Part:  1,
		Value: solution,
	})
}

func partTwo(contents []string) {
	instructionList := parseInput(contents)
	registers := map[string]int{"a": 0, "b": 0, "c": 1, "d": 0}
	gameState := state{0, instructionList, registers}
	solution := solve(gameState)
	sharedstruct.PrintOutput(sharedstruct.Output{
		Day:   12,
		Part:  2,
		Value: solution,
	})
}

func solve(gameState state) int {
	for {
		handleInstruction(gameState.instructions[gameState.index], &gameState.registers, &gameState.index)
		if gameState.index >= len(gameState.instructions) {
			break
		}
	}
	return gameState.registers["a"]
}

func handleInstruction(instruction instruction, registers *map[string]int, index *int) {
	if instruction.instructionType == "cpy" {
		copyValue := getRegisterValOrInt(instruction.operand1, *registers)
		(*registers)[instruction.operand2] = copyValue
		// Bump the index by 1:
		*index++
		return
	}

	if instruction.instructionType == "inc" || instruction.instructionType == "dec" {
		increment := 1
		if instruction.instructionType == "dec" {
			increment = -1
		}
		(*registers)[instruction.operand1] += increment
		// Bump the index by 1:
		*index++
		return
	}

	if instruction.instructionType == "jnz" {
		// Is it in the registers key?
		op1Value := getRegisterValOrInt(instruction.operand1, *registers)
		if op1Value == 0 {
			// Bump the index by 1:
			*index++
			return
		}

		// Otherwise, we move the index by the amount specified in operand2:
		operand2Val, _ := strconv.Atoi(instruction.operand2)
		*index += operand2Val
		return
	}
}

func getRegisterValOrInt(operand string, registers map[string]int) int {
	if _, ok := registers[operand]; ok {
		return registers[operand]
	}

	result, _ := strconv.Atoi(operand)

	return result
}

func parseInput(contents []string) []instruction {
	instructionList := []instruction{}
	for _, line := range contents {
		instructionType := line[0:3]
		temp := strings.Split(line[4:], " ")
		operand1 := temp[0]
		operand2 := ""
		// temp[1] may or may not exist:
		if len(temp) > 1 {
			operand2 = temp[1]
		} else {
			operand2 = ""
		}

		instructionList = append(instructionList, instruction{
			instructionType: instructionType,
			operand1:        operand1,
			operand2:        operand2,
		})
	}

	return instructionList
}
