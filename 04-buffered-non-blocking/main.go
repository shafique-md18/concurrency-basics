package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ninjaStarTargetPractice_Single() {
	channel := make(chan string)
	go throwingNinjaStar(channel)
	fmt.Println(<-channel)
}

func throwingNinjaStar(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	score := rand.Intn(10)
	channel <- fmt.Sprintf("You scored: %d", score)
}

func ninjaStarTargetPractice_Iterative() {
	channel := make(chan string)
	numRounds := 3
	go throwingNinjaStar_Iterative(channel, 3)
	for i := 0; i < numRounds; i++ {
		// blocking operation - waits until there is something in the channel
		// because channel size is 0
		fmt.Println(<-channel)
	}
}

func throwingNinjaStar_Iterative(channel chan string, numRounds int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		// blocking operation - waits until the channel is consumed
		channel <- fmt.Sprintf("You scored: %d", score)
	}
}

func ninjaStarTargetPractice_BufferedNonBlocking() {
	channel := make(chan string, 3)
	numRounds := 3
	go throwingNinjaStar_BufferedNonBlocking(channel, 3)
	for i := 0; i < numRounds; i++ {
		// non-blocking operation - only waits when channel is empty
		// because channel size is 3
		fmt.Println(<-channel)
	}
}

func throwingNinjaStar_BufferedNonBlocking(channel chan string, numRounds int) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		// non-blocking operation - only waits when channel is full
		channel <- fmt.Sprintf("You scored: %d", score)
	}
}

func ninjaStarTargetPractice_BufferedNonBlockingWithClose() {
	channel := make(chan string, 3)
	go throwingNinjaStar_BufferedNonBlockingWithClose(channel)
	for {
		message, open := <-channel
		if !open {
			break
		}
		fmt.Println(message)
	}
}

func throwingNinjaStar_BufferedNonBlockingWithClose(channel chan string) {
	rand.Seed(time.Now().UnixNano())
	numRounds := 3
	for i := 0; i < numRounds; i++ {
		score := rand.Intn(10)
		channel <- fmt.Sprintf("You scored: %d", score)
	}
	close(channel)
}

func main() {
	ninjaStarTargetPractice_BufferedNonBlocking()
}
