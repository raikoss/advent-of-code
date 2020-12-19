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
		take := int(float64((initialValues.ceil + 1)) / math.Pow(float64(2), float64(i+1))) // get half of ceil on i = 0, then 1/4 on i = 1, then 1/8 on i = 2, etc
		if string(char) == lowerHalfSep {
			currentValues.ceil = currentValues.ceil - take
		} else { // char == "B"
			currentValues.floor = currentValues.floor + take
		}
	}

	return currentValues.floor // or ceil, they are equal at the end
}

func getPassRowColumn(pass string) (int, int) {
	rowString := string(pass[:7])
	columnString := string(pass[7:])

	row := decodePart(rowString, 127, "F")
	column := decodePart(columnString, 7, "L")

	return row, column
}

func getSeatID(row, column int) int {
	return row*8 + column
}

func getSeatIDByPass(pass string) int {
	return getSeatID(getPassRowColumn(pass))
}

func getHighestSeatID(lines []string) int {
	higestSeatID := 0

	for _, pass := range lines {
		seatID := getSeatIDByPass(pass)

		if seatID > higestSeatID {
			higestSeatID = seatID
		}
	}

	return higestSeatID
}

// ugly but gets the job done I guess
func getMySeatID(lines []string) int {
	plane := make([][]string, 8) // columns
	for i := range plane {
		plane[i] = make([]string, 128) // rows
	}

	for _, pass := range lines {
		row, column := getPassRowColumn(pass)

		plane[column][row] = pass
	}

	// fmt.Printf("The plane is now %#v\n", plane)

	for column := range plane {
		for row := range plane[column] {
			if row == 0 || row == 127 {
				continue // so we don't check out of bounds, seat is not on edge anyways
			}

			isAtEdge := plane[column][row+1] == "" || plane[column][row-1] == ""

			if plane[column][row] == "" && !isAtEdge {
				// fmt.Printf("My seat id is %d,%d\n", column, row)
				return getSeatID(row, column)
			}
		}
	}

	return 0
}

func main() {
	lines := readFile("input.txt")

	highestSeatID := getHighestSeatID(lines)
	fmt.Printf("The highest seat id is %d\n", highestSeatID)

	mySeatID := getMySeatID(lines)
	fmt.Printf("My seat ID is %d\n", mySeatID)
}
