package main

type waypoint struct {
	x int
	y int
}

func (w *waypoint) goDirection(letter string, value int) {
	switch letter {
	case "N":
		w.y += value
	case "E":
		w.x += value
	case "S":
		w.y -= value
	case "W":
		w.x -= value
	}
}
