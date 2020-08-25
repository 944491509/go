package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	//lock.Lock()
	//x = x + 1
	//lock.Unlock()   // 一般写法 添加锁
	atomic.AddInt64(&x, 1)
	wg.Done()
}

func main() {
	start := time.Now()
	wg.Add(100000)
	for i := 0; i < 100000; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(x)
	fmt.Println(time.Now().Sub(start))
}
