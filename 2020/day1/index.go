package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readFile() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		// logs it couldn't find the file
		log.Fatal(err)
	}

	// closes the file I guess?
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int // slice = JS []

	for scanner.Scan() {
		line := scanner.Text()            // reads one line?
		number, err := strconv.Atoi(line) // convert line to int
		if err != nil {
			// if input isn't an int, shut down?
			os.Exit(2)
		}

		// add line int to list of numbers
		numbers = append(numbers, number)
	}

	// logs if the scanner encountered an error at some point?
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}

func getTwiceMultipliedSum(numbers []int) (int, error) {
	for _, value1 := range numbers {
		for _, value2 := range numbers {
			sum := value1 + value2

			if sum == 2020 {
				fmt.Printf("Value1 is %d\n", value1)
				fmt.Printf("Value2 is %d\n", value2)
				return value1 * value2, nil
			}
		}
	}

	return 0, errors.New("no numbers add up to 2020")
}

func getThriceMultipliedSum(numbers []int) (int, error) {
	for _, value1 := range numbers {
		for _, value2 := range numbers {
			for _, value3 := range numbers {
				sum := value1 + value2 + value3

				if sum == 2020 {
					fmt.Printf("Value1 is %d\n", value1)
					fmt.Printf("Value2 is %d\n", value2)
					fmt.Printf("Value3 is %d\n", value3)
					return value1 * value2 * value3, nil
				}
			}
		}
	}

	return 0, errors.New("no numbers add up to 2020")
}

func main() {
	var numbers []int = readFile()

	twiceSum, twiceErr := getTwiceMultipliedSum(numbers)
	thriceSum, thriceErr := getThriceMultipliedSum(numbers)

	// if no numbers equal 2020, log and exit
	if twiceErr != nil {
		log.Fatal(twiceErr)
		os.Exit(1)
	}

	// if no numbers equal 2020, log and exit
	if thriceErr != nil {
		log.Fatal(thriceErr)
		os.Exit(1)
	}

	fmt.Printf("The multiplied sum twice is %d\n", twiceSum)
	fmt.Printf("The multiplied sum thrice is %d\n", thriceSum)
}
