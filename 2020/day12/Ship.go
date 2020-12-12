package main

import (
	"strconv"
)

type ship struct {
	posX  int
	posY  int
	angle int
}

func getAngleDifference(angle1, angle2 int) (angle int) {
	if angle1 >= 0 {
		angle = (angle2 + angle1) % 360
	} else {
		angle = (angle2 + angle1 + 360) % 360
	}

	return angle
}

func getDirectionMap() (directionMap map[int]string) {
	directionMap = map[int]string{0: "N", 90: "E", 180: "S", 270: "W"}
	return
}

func (s *ship) runInstruction(instruction string) {
	// fmt.Printf("Running instruction %v\n", instruction)
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

func (s *ship) runInstructionWithWaypoint(w *waypoint, instruction string) {
	// fmt.Printf("Running instruction %v\n", instruction)
	operation := string(instruction[0])
	value, _ := strconv.Atoi(instruction[1:])

	switch operation {
	case "N", "E", "S", "W":
		w.goDirection(operation, value)
	case "L":
		s.turnWaypoint(w, -value)
	case "R":
		s.turnWaypoint(w, value)
	case "F":
		s.goForwardWithWaypoint(w, value)
	}
}

func (s *ship) turn(degrees int) {
	newAngle := getAngleDifference(degrees, s.angle)

	s.angle = newAngle
	// fmt.Printf("New ship angle is %d\n", s.angle)
}

func (s *ship) turnWaypoint(w *waypoint, degrees int) {
	newX := w.x
	newY := w.y

	if degrees == -90 || degrees == 270 {
		newX = -w.y
		newY = w.x
	} else if degrees == -180 || degrees == 180 {
		newX = -w.x
		newY = -w.y
	} else if degrees == -270 || degrees == 90 {
		newX = w.y
		newY = -w.x
	}

	w.x = newX
	w.y = newY
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
	directionMap := getDirectionMap()
	direction := directionMap[s.angle]

	s.goDirection(direction, value)
}

func (s *ship) goForwardWithWaypoint(w *waypoint, value int) {
	// fmt.Printf("Ship at %d,%d, waypoint at %d,%d, moving %d times the difference\n", s.posX, s.posY, w.posX, w.posY, value)

	newX := s.posX + w.x*value
	newY := s.posY + w.y*value

	s.posX = newX
	s.posY = newY
	// fmt.Printf("Ship now at %d,%d\n", s.posX, s.posY)
}
