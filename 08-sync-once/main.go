package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var missionCompleted bool

func completeNinjaMissionOnlyOnce() {
	var wg sync.WaitGroup
	var once sync.Once
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func(ninja int) {
			if isMissionCompleted() {
				// this is used when we want to call a function only once
				once.Do(func() { markMissionCompleted(ninja) })
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}

func checkMissionCompletionStatus() {
	if missionCompleted {
		fmt.Println("Mission completed!")
	} else {
		fmt.Println("Mission failed!")
	}
}

func markMissionCompleted(ninja int) {
	fmt.Println("Mission completed by ninja", ninja)
	missionCompleted = true
}

// returns true if mission is completed, probability of 1 in 10
func isMissionCompleted() bool {
	rand.Seed(time.Now().UnixNano())
	return 0 == rand.Intn(10)
}

func main() {
	completeNinjaMissionOnlyOnce()
}
