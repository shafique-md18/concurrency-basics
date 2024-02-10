package main

import (
	"fmt"
)

func captainElectionUsingSelect() {
	ninja1, ninja2 := make(chan string), make(chan string)

	go electCaptain(ninja1, "Ninja 1")
	go electCaptain(ninja2, "Ninja 2")

	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	}
}

func electCaptain(ninja chan string, message string) {
	ninja <- message
}

// The interesting aspect of this code is that the select statement randomly chooses
// between ninja1 and ninja2 when both channels have values available.
// This randomness makes the election roughly fair, as both ninjas have an equal chance
// of being selected in each iteration.
func captainElectionRoughlyFair() {
	ninja1 := make(chan interface{})
	close(ninja1) // closed channels are non-blocking to select and the received value is nil
	ninja2 := make(chan interface{})
	close(ninja2)

	var ninja1Count, ninja2Count int
	// this can be used as a polling mechanism
	for i := 0; i < 1000; i++ {
		select {
		// if both channels are closed hence select will
		// randomly choose either of the channel and these channels will return nil
		case <-ninja1:
			ninja1Count++
		case <-ninja2:
			ninja2Count++
		default:
			fmt.Println("Neither")
		}
	}

	fmt.Printf("Ninja 1: %d, Ninja 2: %d\n", ninja1Count, ninja2Count)
}

func main() {
	captainElectionRoughlyFair()
}
