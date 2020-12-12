package main

type waypoint struct {
	posX int
	posY int
}

// func (w *waypoint) runInstruction(instruction string) {
// 	// fmt.Printf("Running instruction %v\n", instruction)
// 	operation := string(instruction[0])
// 	value, _ := strconv.Atoi(instruction[1:])

// 	switch operation {
// 	case "N", "E", "S", "W":
// 		w.goDirection(operation, value)
// 	case "L":
// 		w.turn(-value)
// 	case "R":
// 		w.turn(value)
// 	case "F":
// 		w.goForward(value)
// 	}
// }

func (w *waypoint) goDirection(letter string, value int) {
	switch letter {
	case "N":
		w.posY += value
	case "E":
		w.posX += value
	case "S":
		w.posY -= value
	case "W":
		w.posX -= value
	}
}
