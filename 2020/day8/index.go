package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

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

func runAccumulate(argument string, accumulator *int) {
	amount, multiplier := getArgumentValues(argument)

	*accumulator += multiplier * amount
}

func runJump(argument string, currentIndex *int) {
	amount, multiplier := getArgumentValues(argument)

	*currentIndex += multiplier * amount
}

func runInstruction(instruction string, accumulator *int, currentIndex *int) {
	splitOperation := strings.Split(instruction, " ")
	operation := splitOperation[0]
	argument := splitOperation[1]

	if operation == "acc" {
		runAccumulate(argument, accumulator)
		*currentIndex++
	} else if operation == "jmp" {
		runJump(argument, currentIndex)
	} else if operation == "nop" {
		*currentIndex++
	}

	// fmt.Printf("Ran %v, accumulator is now %d\n", instruction, *accumulator)
}

func runProgram(lines []string, accumulator int) int {
	instructionsRan := make(map[int]bool) // O(1) lookup time :)
	currentIndex := 0

	// while something is unrun
	for {
		// fmt.Printf("Current index is now at %d\n", currentIndex)
		// fmt.Printf("Map %#v", instructionsRan)
		instruction := lines[currentIndex]

		// if already ran this instruction, return value of accumulator
		_, exists := instructionsRan[currentIndex]
		if exists {
			return accumulator
		}

		instructionsRan[currentIndex] = true
		runInstruction(instruction, &accumulator, &currentIndex)
	}
}

func main() {
	lines := readFile("input.txt")
	accumulatorBeforeRerun := runProgram(lines, 0)

	fmt.Printf("The value of the accumulator before running a second time is %d\n", accumulatorBeforeRerun)
}
