package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type programData struct {
	ranInstructions map[int]bool // O(1) lookup time :)
	accumulator     int
	currentIndex    int
	hasSwapped      bool
}

// was bigger initially, but now it's just a map lol
type swapData struct {
	swappedInstructions map[int]bool
}

type runnableOperation func(*programData, string)

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func getArgumentValues(argument string) (int, int) {
	sign := string(argument[0])
	amount, err := strconv.Atoi(argument[1:])
	multiplier := 0

	if err != nil {
		log.Fatal("expected number in operation string to be a number")
	}

	if sign == "+" {
		multiplier = 1
	} else if sign == "-" {
		multiplier = -1
	}

	return amount, multiplier
}

func runAccumulate(data *programData, argument string) {
	amount, multiplier := getArgumentValues(argument)

	data.accumulator += multiplier * amount
	data.currentIndex++
}

func runJump(data *programData, argument string) {
	amount, multiplier := getArgumentValues(argument)

	data.currentIndex += multiplier * amount
}

// we don't use argument, but pass it to match the other run* functions
func runNoop(data *programData, argument string) {
	data.currentIndex++
}

func splitInstruction(instruction string) (string, string) {
	splitOperation := strings.Split(instruction, " ")
	operation := splitOperation[0]
	argument := splitOperation[1]

	return operation, argument
}

func runInstruction(instruction string, data *programData) {
	operation, argument := splitInstruction(instruction)

	if operation == "acc" {
		runAccumulate(data, argument)
	} else if operation == "jmp" {
		runJump(data, argument)
	} else if operation == "nop" {
		runNoop(data, argument)
	}
}

// ooooooof
func trySwap(data *programData, swapData *swapData, argument string, originalFunc, swapFunc runnableOperation) {
	_, swappedInPreviousIteration := swapData.swappedInstructions[data.currentIndex]

	if !data.hasSwapped && !swappedInPreviousIteration {
		swapData.swappedInstructions[data.currentIndex] = true
		data.hasSwapped = true
		swapFunc(data, argument)
	} else {
		originalFunc(data, argument)
	}
}

func trySwapRunInstruction(instruction string, data *programData, swapData *swapData) {
	operation, argument := splitInstruction(instruction)

	if operation == "acc" {
		runAccumulate(data, argument)
	} else if operation == "jmp" {
		trySwap(data, swapData, argument, runJump, runNoop)
	} else if operation == "nop" {
		trySwap(data, swapData, argument, runNoop, runJump)
	}
}

func getAccumulatorByRunningProgram(lines []string) int {
	data := programData{make(map[int]bool), 0, 0, false}

	// while something is unrun
	for {
		instruction := lines[data.currentIndex]

		// if already ran this instruction, return value of accumulator
		_, exists := data.ranInstructions[data.currentIndex]
		if exists {
			return data.accumulator
		}

		data.ranInstructions[data.currentIndex] = true
		runInstruction(instruction, &data)
	}
}

func resetProgram(data *programData) {
	data.ranInstructions = make(map[int]bool)
	data.currentIndex = 0
	data.accumulator = 0
	data.hasSwapped = false
}

func getAccumulatorByChangingOperations(lines []string) int {
	swapData := swapData{make(map[int]bool)}
	data := programData{make(map[int]bool), 0, 0, false}

	// while something is unrun
	for {
		if data.currentIndex == len(lines) {
			return data.accumulator
		}

		instruction := lines[data.currentIndex]

		_, instructionExists := data.ranInstructions[data.currentIndex]
		// if exists, reset program data, and swapData has received a new entry in tried swap index
		if instructionExists {
			resetProgram(&data)
		} else {
			data.ranInstructions[data.currentIndex] = true
			trySwapRunInstruction(instruction, &data, &swapData)
		}
	}
}

func main() {
	lines := readFile("input.txt")

	accumulatorBeforeRerun := getAccumulatorByRunningProgram(lines)
	fmt.Printf("The value of the accumulator before running a second time is %d\n", accumulatorBeforeRerun)

	accumulatorAfterSwapping := getAccumulatorByChangingOperations(lines)
	fmt.Printf("The value of the accumulator after swapping jump and noop and terminating is %d\n", accumulatorAfterSwapping)
}
