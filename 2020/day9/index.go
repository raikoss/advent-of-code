package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile(path string) []int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []int

	for scanner.Scan() {
		line, err := strconv.Atoi(scanner.Text())

		if err != nil {
			log.Fatal("a line was not a number!")
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func isSumOfEarlierNumbers(numbers []int, sum int) bool {
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == sum {
				return true
			}
		}
	}

	return false
}

func getFirstNonSum(lines []int, preambleLength int) (int, error) {
	// first preamble numbers
	var currentNumbers []int // [including:excluding]
	notSumNumber := 0

	for i := 0; i < len(lines)-preambleLength; i++ {
		sumIndex := i + preambleLength
		currentNumbers = lines[i:sumIndex]

		currentNumber := lines[sumIndex]
		isSum := isSumOfEarlierNumbers(currentNumbers, currentNumber)

		if !isSum {
			fmt.Printf("Number is at i %d\n", sumIndex)
			return currentNumber, nil
		}
	}

	if notSumNumber == 0 {
		return 0, fmt.Errorf("no numbers were not equal to the sum of the %d earlier numbers", preambleLength)
	}

	return notSumNumber, nil
}

func main() {
	lines := readFile("input.txt")
	nonSum, err := getFirstNonSum(lines, 25)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The non-sum number is %d!\n", nonSum)
}
