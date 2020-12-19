package main

import "fmt"

func sayNumber() {

}

func playGame(startingNumbers []int, endTurn int) int {
	lastMove := make(map[int]int)
	prevSaidNumber := 0
	age, exists := 0, false

	for i := 0; i < endTurn; i++ {
		// say starting number, save which turn it was said
		if i < len(startingNumbers) {
			startingNumber := startingNumbers[i]
			lastMove[startingNumber] = i
			prevSaidNumber = startingNumber
			continue
		}

		if exists {
			prevSaidNumber = i - age - 1 // -1 since we're 1 turn further than previous number
		} else {
			prevSaidNumber = 0
		}

		age, exists = lastMove[prevSaidNumber]
		// fmt.Printf("Age since number %d was said: %d\n", prevSaidNumber, age)
		// fmt.Printf("Saying number %d!\n", age)

		lastMove[prevSaidNumber] = i
		// fmt.Printf("Turn %d - Last move map: %#v\n", i, lastMove)
	}

	return prevSaidNumber
}

func main() {
	input := []int{0, 1, 5, 10, 3, 12, 19}
	// input := []int{0, 1, 2, 3, 4}
	endTurnPart1 := 2020

	part1Ans := playGame(input, endTurnPart1)

	fmt.Printf("The %d number said was %d\n", endTurnPart1, part1Ans)
}
