package main

import (
	"fmt"
	"sync"
	"time"
)

// inconsistent read
var countWithoutMutex int

func incrementWithoutMutex() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go increment(&wg)
	}
	wg.Wait()
	fmt.Println(countWithoutMutex)
}

func increment(wg *sync.WaitGroup) {
	countWithoutMutex++
	wg.Done()
}

var (
	// provides safety to count variable
	lock   sync.Mutex
	rwLock sync.RWMutex
	count  int
)

func incrementWithMutex() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go incrementWithLock(&wg)
	}
	wg.Wait()
	fmt.Println(count)
}

// both lock and wait group should be used
func incrementWithLock(wg *sync.WaitGroup) {
	lock.Lock()
	defer func() {
		lock.Unlock()
		wg.Done()
	}()

	count++
}

// RWMutex - any number of threads can acquire read locks but only one thread can acquire write lock!
// Thread having write lock prevents other threads to acquire read lock
func readAndWrite() {
	var wg sync.WaitGroup
	wg.Add(9)

	go read(&wg)
	go read(&wg)
	go read(&wg)
	go read(&wg)
	go read(&wg)
	go read(&wg)
	// reads and write are mutually exclusive
	go write(&wg)
	go read(&wg)
	go read(&wg)

	wg.Wait()
}

func read(wg *sync.WaitGroup) {
	rwLock.RLock()
	defer func() {
		rwLock.RUnlock()
		wg.Done()
	}()

	fmt.Println("Locking reads")
	time.Sleep(time.Second)
	fmt.Println("Unlocking reads")
}

func write(wg *sync.WaitGroup) {
	rwLock.Lock()
	defer func() {
		rwLock.Unlock()
		wg.Done()
	}()

	fmt.Println("Locking writes")
	time.Sleep(time.Second)
	fmt.Println("Unlocking writes")
}

func main() {
	// incrementWithoutMutex()
	// incrementWithMutex()
	readAndWrite()
}
