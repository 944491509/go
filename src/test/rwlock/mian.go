package main

import (
	"fmt"
	"sync"
	"time"
)

// rwlock

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()
	lock.Lock()
	//rwlock.RLock()
	fmt.Println(x)
	time.Sleep(time.Millisecond)
	lock.Unlock()
	//rwlock.RUnlock()
}

func write() {
	defer wg.Done()
	lock.Lock()
	//rwlock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 5)
	lock.Unlock()
	//rwlock.Unlock()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	time.Sleep(time.Second)
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()

	fmt.Println(time.Now().Sub(start))
}
