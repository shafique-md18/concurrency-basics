package main

import (
	"fmt"
	"strconv"
	"sync"
)

func attackNinjasInGroup() {
	var beeper sync.WaitGroup
	for i := 0; i < 1000; i++ {
		beeper.Add(1) // this should not be within go routine as that will run parallely
		// and we want the count to be in sync
		go attackNinja(strconv.Itoa(i), &beeper)
	}
	beeper.Wait()
	fmt.Println("All ninjas attacked!")
}

func attackNinja(target string, beeper *sync.WaitGroup) {
	fmt.Println("Attacked ninja ", target)
	beeper.Done()
}

func main() {
	attackNinjasInGroup()
}
