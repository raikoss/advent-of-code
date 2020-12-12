package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

type ship struct {
	posX  int
	posY  int
	angle int
}

func (s *ship) runInstruction(instruction string) {
	fmt.Printf("Running instruction %v\n", instruction)
	operation := string(instruction[0])
	value, _ := strconv.Atoi(instruction[1:])

	switch operation {
	case "N", "E", "S", "W":
		s.goDirection(operation, value)
	case "L":
		s.turn(-value)
	case "R":
		s.turn(value)
	case "F":
		s.goForward(value)
	}
}

func (s *ship) turn(degrees int) {
	var newAngle int

	if degrees >= 0 {
		newAngle = (s.angle + degrees) % 360
	} else {
		newAngle = ((s.angle+degrees)%360 + 360) % 360
	}

	s.angle = newAngle
	// fmt.Printf("New ship angle is %d\n", s.angle)
}

func (s *ship) goDirection(letter string, value int) {
	switch letter {
	case "N":
		s.posY += value
	case "E":
		s.posX += value
	case "S":
		s.posY -= value
	case "W":
		s.posX -= value
	}
}

func (s *ship) goForward(value int) {
	directionMap := map[int]string{0: "N", 90: "E", 180: "S", 270: "W"} // creating every time the function is called but I want this to be "global"
	direction := directionMap[s.angle]

	s.goDirection(direction, value)
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
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func navigateShip(lines []string) (int, int) {
	ship := ship{0, 0, 90} // facing east initially

	for _, line := range lines {
		ship.runInstruction(line)
		fmt.Printf("New position: %d,%d\n", ship.posX, ship.posY)
	}

	return ship.posX, ship.posY
}

func getManhattanDistance(lines []string) int {
	x, y := navigateShip(lines)

	absX := int(math.Abs(float64(x)))
	absY := int(math.Abs(float64(y)))

	return absX + absY
}

func main() {
	lines := readFile("input.txt")

	distance := getManhattanDistance(lines)

	fmt.Printf("The manhattan distance from start to end is %d\n", distance)
}
