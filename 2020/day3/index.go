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
			xPos = xPos - lenLine
		}

		if rune(line[xPos]) == rune('#') {
			treesHit++
		}
	}

	return
}

func main() {
	lines := readFile("./input.txt")

	treesHit1 := traverseLines(lines, 1, 1)
	treesHit2 := traverseLines(lines, 1, 3)
	treesHit3 := traverseLines(lines, 1, 5)
	treesHit4 := traverseLines(lines, 1, 7)
	treesHit5 := traverseLines(lines, 2, 1)

	totalTreesHit := treesHit1 * treesHit2 * treesHit3 * treesHit4 * treesHit5

	fmt.Printf("We hit %d trees!\n", totalTreesHit)
}
