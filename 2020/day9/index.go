package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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
	for i := 0; i < len(lines)-preambleLength; i++ {
		sumIndex := i + preambleLength

		currentNumber := lines[sumIndex]
		isSum := isSumOfEarlierNumbers(lines[i:sumIndex], currentNumber)

		if !isSum {
			return currentNumber, nil
		}
	}

	return 0, fmt.Errorf("no numbers were not equal to the sum of the %d earlier numbers", preambleLength)
}

func getFirstSetEqualsNonSum(lines []int, nonSum int) (int, error) {
	var set []int

	for startIndex := 0; startIndex < len(lines); startIndex++ {
		currentSum := 0

		for i := startIndex; i < len(lines); i++ {
			currentSum += lines[i]

			if currentSum == nonSum {
				set = lines[startIndex : i+1]
				break
			}
		}

		if len(set) > 2 {
			sort.Ints(set)

			return set[0] + set[len(set)-1], nil
		}
	}

	return 0, fmt.Errorf("no set totalled up to %d", nonSum)
}

func main() {
	lines := readFile("input.txt")
	nonSum, err := getFirstNonSum(lines, 25)

	if err != nil {
		log.Fatal(err)
	}

	setSum, err := getFirstSetEqualsNonSum(lines, nonSum)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The non-sum number is %d!\n", nonSum)
	fmt.Printf("The sum of the smallest and largest number in set that equalled non-sum is %d!\n", setSum)
}
