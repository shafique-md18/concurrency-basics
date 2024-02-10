package main

import (
	"fmt"
	"time"
)

func attackEvilNinjasWithSleep() {
	startTime := time.Now()
	defer func() {
		fmt.Printf("It took %v for attacking evil ninjas\n", time.Since(startTime))
	}()

	evilNinjas := []string{"Andy", "Tom", "Jack", "Nielson"}

	for idx := range evilNinjas {
		go attack(evilNinjas[idx])
	}

	// without sleep none of the attacks complete
	time.Sleep(time.Second)
}

func attack(target string) {
	time.Sleep(time.Second)
	fmt.Println("Threw ninja starts at " + target)
}

func main() {
	attackEvilNinjasWithSleep()
}
