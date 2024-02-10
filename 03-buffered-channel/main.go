package main

import (
	"fmt"
	"time"
)

// channels are FIFO queue.

// this function gives error -> fatal error: all goroutines are asleep - deadlock!
// channels have 0 capacity. If we add something to channel inside a routine, we must have
// some other routine to consume this item!
// Here deadlockExample() is both the producer and consumer hence the deadlock. (Move this code block inside main!)
func deadlockExample() {
	channel := make(chan bool) // increase capacity to remove deadlock -> make(chan bool, 1)
	channel <- true
	fmt.Println(<-channel)
}

func attackEvilNinjaWithBufferedChannels() {
	startTime := time.Now()
	defer func() {
		fmt.Printf("It took %v for attacking evil ninja\n", time.Since(startTime))
	}()

	smokeSignal := make(chan bool)
	go attackWithChannel("Andy", smokeSignal)
	<-smokeSignal // Go takes care of polling this periodically
}

func attackWithChannel(target string, attacked chan bool) {
	fmt.Println("Throwing ninja starts at " + target)
	time.Sleep(time.Second)
	attacked <- true
}

func main() {
	attackEvilNinjaWithBufferedChannels()
}
