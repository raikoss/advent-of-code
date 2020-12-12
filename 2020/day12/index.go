package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
		line := scanner.Text()

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func navigateShip(lines []string) ship {
	ship := ship{0, 0, 90} // facing east initially

	for _, line := range lines {
		ship.runInstruction(line)
		// fmt.Printf("New position: %d,%d\n", ship.posX, ship.posY)
	}

	return ship
}

func navigateShipWithWaypoint(lines []string) ship {
	ship := ship{0, 0, 0} // doesn't matter which way it faces initially
	waypoint := waypoint{10, 1}

	for _, line := range lines {
		ship.runInstructionWithWaypoint(&waypoint, line)
	}

	return ship
}

func getManhattanDistance(s ship) int {
	absX := int(math.Abs(float64(s.posX)))
	absY := int(math.Abs(float64(s.posY)))

	return absX + absY
}

func main() {
	lines := readFile("input.txt")
	// lines := []string{
	// 	"F10",
	// 	"N3",
	// 	"F7",
	// 	"R90",
	// 	"F11",
	// }

	navigatedShip := navigateShip(lines)
	distance := getManhattanDistance(navigatedShip)

	fmt.Printf("The manhattan distance from start to end is %d\n", distance)

	shipWithWaypoint := navigateShipWithWaypoint(lines)
	distanceWithWaypoint := getManhattanDistance(shipWithWaypoint)

	fmt.Printf("The distance using the waypoint is %d\n", distanceWithWaypoint)
}
