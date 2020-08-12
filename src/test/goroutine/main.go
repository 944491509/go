package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Println("hello", i)
}

// 程序启动之后会创建一个goroutine去执行
func main() {
	for i := 0; i < 100; i++ {
		//go hello(i)  // 开启一个单独的goroutine去执行hello函数(任务)
		//i := i
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main")
	// main 函数结束了后, 由main函数启动的goroutine也都结束了
	time.Sleep(time.Second)
}
