package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// EXAMPLE 1

/*
 * Cond is used to wait or announce occurance of an event.
 * Cond has lock associated with it, before waiting or modifying the cond, lock must be acquired.
 */

var isNinjaReady bool

func gearUpForMission() {
	cond := sync.NewCond(&sync.Mutex{})
	go gearUpNinja(cond)

	cond.L.Lock()
	defer cond.L.Unlock()
	for !isNinjaReady {
		fmt.Println("Waiting for ninja to get ready")
		cond.Wait()
	}
	fmt.Println("Leaving for mission.")
}

func gearUpNinja(cond *sync.Cond) {
	arbitrarySleep()
	fmt.Println("Ninja is now ready.")
	isNinjaReady = true
	cond.Signal()
}

func arbitrarySleep() {
	sleepSeconds := rand.Intn(5)
	time.Sleep(time.Second * time.Duration(sleepSeconds))
}

// EXAMPLE 2

// calls ninjas to be on standby
// boardcasts the ninjas to start the mission
func gearUpAllNinjasForMission() {
	fmt.Println("Calling ninjas to be on standby.")

	cond := sync.NewCond(&sync.Mutex{})
	var waitForCompletion sync.WaitGroup
	var waitForStandby sync.WaitGroup
	for i := 0; i < 5; i++ {
		waitForCompletion.Add(1)
		waitForStandby.Add(1)
		go standbyNinja(i, cond, &waitForCompletion, &waitForStandby)
	}

	// wait until we get standby status from all routines.
	waitForStandby.Wait()

	// send signal to all routines holding lock on condition at once.
	cond.Broadcast()

	// wait until we get completion status from all routines.
	waitForCompletion.Wait()

	fmt.Println("All ninjas started the mission.")
}

// single ninja to be on standby until the mission is started
/*
 It may look like that the goroutine waiting(rec.cond.Wait()) is holding the lock whole time(rec.Lock()),
 but its not, Internally cond.Wait() unlocks it and it locks it again only when it wakes up by other go routine.

 https://kaviraj.me/understanding-condition-variable-in-go/
*/
func standbyNinja(ninja int, cond *sync.Cond, waitForCompletion *sync.WaitGroup, waitForStandby *sync.WaitGroup) {
	fmt.Println("Ninja", ninja, "on standby.")

	cond.L.Lock()
	defer func() {
		cond.L.Unlock()
		waitForCompletion.Done()
	}()

	waitForStandby.Done()
	cond.Wait()

	fmt.Println("Ninja", ninja, "started mission.")
}

func main() {
	// gearUpForMission()
	gearUpAllNinjasForMission()
}
