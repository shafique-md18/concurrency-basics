package main

import (
	"fmt"
	"time"
)

func attackEvilNinjaWithSimpleChannel() {
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
	attackEvilNinjaWithSimpleChannel()
}
