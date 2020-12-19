package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type data struct {
	floor int
	ceil  int
}

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func decodePart(part string, ceil int, lowerHalfSep string) int {
	initialValues := data{0, ceil}
	currentValues := data{initialValues.floor, initialValues.ceil}

	for i, char := range part {
		take := int(float64((initialValues.ceil + 1)) / math.Pow(float64(2), float64(i+1)))
		// fmt.Printf("Taking %d\n", take)
		// fmt.Printf("Current values %#v\n", currentValues)
		if string(char) == lowerHalfSep {
			currentValues.ceil = currentValues.ceil - take
		} else { // char == "B"
			currentValues.floor = currentValues.floor + take
		}
	}

	return currentValues.floor // or ceil, they are equal at the end
}

func getSeatID(pass string) int {
	rowString := string(pass[:7])
	columnString := string(pass[7:])

	row := decodePart(rowString, 127, "F")
	column := decodePart(columnString, 7, "L")

	// fmt.Printf("Code %v equals row %d, column %d\n", pass, row, column)

	return row*8 + column
}

func getHighestSeatID(lines []string) int {
	higestSeatID := 0

	for _, pass := range lines {
		seatID := getSeatID(pass)

		if seatID > higestSeatID {
			higestSeatID = seatID
		}
	}

	return higestSeatID
}

func main() {
	lines := readFile("input.txt")

	highestSeatID := getHighestSeatID(lines)

	fmt.Printf("The highest seat id is %d\n", highestSeatID)
}
