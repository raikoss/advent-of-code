package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// hnnng
func readFile(path string) (int, []string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var buses []string
	var estimate int
	lineIndex := 0

	for scanner.Scan() {
		line := scanner.Text()

		if lineIndex == 0 {
			intLine, err := strconv.Atoi(line)

			if err != nil {
				log.Fatal("Bruh")
			}

			estimate = intLine
		} else if lineIndex == 1 {
			buses = strings.Split(line, ",")
		}

		lineIndex++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return estimate, buses
}

func getLowestBusTime(nearestBusTimes map[int]int) (int, int) {
	lowestTime := 0
	lowestTimeBusID := 0

	for bus, overflow := range nearestBusTimes {
		// first element is lowest in the beginning, other have to have lower overflow
		if lowestTimeBusID == 0 || overflow < lowestTime {
			lowestTimeBusID = bus
			lowestTime = overflow
		}
	}

	return lowestTimeBusID, lowestTime
}

func findNearestTime(estimate int, bus int) int {
	return bus - estimate%bus
}

func findBusWithNearestTime(estimate int, buses []string) (int, int) {
	nearestBusTimes := make(map[int]int)

	for _, bus := range buses {
		busID, err := strconv.Atoi(bus)

		// if 'x', skip it
		if err != nil {
			continue
		}

		overflow := findNearestTime(estimate, busID)

		nearestBusTimes[busID] = overflow
	}

	fmt.Printf("Times by id: %#v\n", nearestBusTimes)

	return getLowestBusTime(nearestBusTimes)
}

func getIDMultipliedWithLowestTime(busID, lowestTime int) int {
	return busID * lowestTime
}

func main() {
	estimate, buses := readFile("input.txt")

	part1Answer := getIDMultipliedWithLowestTime(findBusWithNearestTime(estimate, buses))

	fmt.Printf("The ID times lowest waiting time for bus is %d\n", part1Answer)
}
