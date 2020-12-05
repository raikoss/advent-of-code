package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func traverseLines(lines []string, down, right int) (treesHit int) {
	xPos := 0

	for yPos := 0; yPos < len(lines)-down; {
		yPos += down
		line := lines[yPos]
		lenLine := len(line)
		xPos += right

		if xPos >= lenLine {
			fmt.Printf("Would have hit %d, length of line is %d\n", xPos, lenLine)
			xPos = xPos - lenLine
		} else {

		}

		if rune(line[xPos]) == rune('#') {
			fmt.Printf("There was a tree at %d,%d\n", yPos, xPos)
			treesHit++
		}
	}

	return
}

func main() {
	lines := readFile("./input.txt")

	treesHit := traverseLines(lines, 1, 3)

	fmt.Printf("We hit %d trees!\n", treesHit)
}
