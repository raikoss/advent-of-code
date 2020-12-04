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
	minLength int
	maxLength int
	letter    rune
	password  string
}

func testPolicy(p policy) bool {
	matches := 0

	for _, char := range p.password {
		if char == p.letter {
			matches++
		}
	}

	return p.minLength <= matches && matches <= p.maxLength
}

func getPolicy(policyString string) policy {
	regex := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)
	passwordBytes := []byte(policyString)
	minLength, _ := strconv.Atoi(string(regex.FindSubmatch(passwordBytes)[1])) // match 0 is whole string
	maxLength, _ := strconv.Atoi(string(regex.FindSubmatch(passwordBytes)[2])) // match 0 is whole string
	letter, _ := utf8.DecodeRune(regex.FindSubmatch(passwordBytes)[3])
	password := string(regex.FindSubmatch(passwordBytes)[4])

	return policy{minLength, maxLength, letter, password}
}

func main() {
	policies := readFile("./input.txt")

	var validPasswords []string

	for _, policyString := range policies {
		policy := getPolicy(policyString)
		isValid := testPolicy(policy)

		if isValid {
			validPasswords = append(validPasswords, policy.password)
		}
	}

	fmt.Printf("Number of valid passwords: %d\n", len(validPasswords))
}
