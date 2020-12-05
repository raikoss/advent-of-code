package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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
		lines = append(lines, scanner.Text())
	}

	// logs if the scanner encountered an error at some point?
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func getPassports(lines []string) (passports []string) {
	currentPassport := ""

	// build one-liners of passport per passport
	for _, line := range lines {
		// when newline is met, move currentPassport into passport list
		if line == "" {
			passports = append(passports, currentPassport)
			currentPassport = ""
			continue
		}

		currentPassport += line + " "
	}

	return
}

func isValidPassport(passport string) bool {
	fields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid" /*"cid" not required! */}
	isValid := true

	for _, field := range fields {
		// fmt.Printf("Checking %v for field %v\n", passport, field)
		if !strings.Contains(passport, field+":") {
			isValid = false
			break
		}
	}

	return isValid
}

func getNumValidPassports(lines []string) int {
	passports := getPassports(lines)
	fmt.Printf("Amount of passports: %d\n", len(passports))
	var validPassports []string

	for _, passport := range passports {
		if isValidPassport(passport) {
			validPassports = append(validPassports, passport)
		}
	}

	return len(validPassports)
}

func main() {
	lines := readFile("./input.txt")

	numValidPassports := getNumValidPassports(lines)

	fmt.Printf("There are %d valid passorts!\n", numValidPassports)
}
