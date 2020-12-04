package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func readFile(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		// logs it couldn't find the file
		log.Fatal(err)
	}

	// closes the file I guess?
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string

	for scanner.Scan() {
		line := scanner.Text() // reads one line

		if err != nil {
			// if input isn't an int, shut down?
			os.Exit(2)
		}

		// add line int to list of lines
		lines = append(lines, line)
	}

	// logs if the scanner encountered an error at some point?
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

type policy struct {
	firstNumber  int
	secondNumber int
	letter       rune
	password     string
}

func testPolicyLength(p policy) bool {
	matches := 0

	for _, char := range p.password {
		if char == p.letter {
			matches++
		}
	}

	return p.firstNumber <= matches && matches <= p.secondNumber
}

func testPolicyPosition(p policy) bool {
	firstNumberMatch := false
	secondNumberMatch := false

	for index, char := range p.password {
		if char != p.letter {
			continue
		}

		if index+1 == p.firstNumber {
			firstNumberMatch = true
		} else if index+1 == p.secondNumber {
			secondNumberMatch = true
		}
	}

	return (firstNumberMatch || secondNumberMatch) && !(firstNumberMatch && secondNumberMatch) // native nand somewhere?
}

func getPolicy(policyString string) policy {
	regex := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

	passwordBytes := []byte(policyString)
	firstNumber, _ := strconv.Atoi(string(regex.FindSubmatch(passwordBytes)[1]))  // match 0 is whole string
	secondNumber, _ := strconv.Atoi(string(regex.FindSubmatch(passwordBytes)[2])) // byte[] -> string -> int
	letter, _ := utf8.DecodeRune(regex.FindSubmatch(passwordBytes)[3])            // byte[] -> string -> int
	password := string(regex.FindSubmatch(passwordBytes)[4])

	return policy{firstNumber, secondNumber, letter, password}
}

func getNumLengthPolicies(policies []string) int {
	var validPasswords []string

	for _, policyString := range policies {
		policy := getPolicy(policyString)
		isValid := testPolicyLength(policy)

		if isValid {
			validPasswords = append(validPasswords, policy.password)
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", len(validPasswords))

	return len(validPasswords)
}

func getNumPositionPolicies(policies []string) int {
	var validPasswords []string

	for _, policyString := range policies {
		policy := getPolicy(policyString)
		isValid := testPolicyPosition(policy)

		if isValid {
			validPasswords = append(validPasswords, policy.password)
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", len(validPasswords))

	return len(validPasswords)
}

func main() {
	policies := readFile("./input.txt")

	numLengthPolicies := getNumLengthPolicies(policies)
	numPositionPolicies := getNumPositionPolicies(policies)

	fmt.Printf("Number of valid passwords with length requirement: %d\n", numLengthPolicies)
	fmt.Printf("Number of valid passwords with position requirement: %d\n", numPositionPolicies)
}
