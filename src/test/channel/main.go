package main

import (
	"fmt"
	"sync"
)

var b chan int

func noBufChannel() {
	b = make(chan int) // 不带缓冲区通道的初始化
	go func() {
		x := <-b
		fmt.Println("后台goroutine从通道中取到了", x)
	}()
	b <- 10 // 卡主了
	fmt.Println("10发送到了通道b中了...")
}

func bufChannel() {
	fmt.Println(b)
	b = make(chan int, 16) // 带缓冲区的通道的初始化
	b <- 10
	fmt.Println("10发送到无缓冲区通道了..")
	b <- 20
	fmt.Println("20发送到无缓冲区通道了..")
	x := <-b
	fmt.Println("从通道中取到了", x)
	close(b)
}

var wg sync.WaitGroup
var once sync.Once

func main() {
	//noBufChannel()
	//bufChannel()
	a := make(chan int, 100)
	b := make(chan int, 100)
	wg.Add(3)
	go f1(a)
	go f2(a, b)
	go f2(a, b)
	wg.Wait()
	for ret := range b {
		fmt.Println(ret)
	}
}

// 启动一个goroutine,生成100个数发送到ch1
// 启动一个goroutine,从ch1中取值,计算其平方放到ch2
// 在main中从ch2取值打印出来

func f1(ch1 chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		ch1 <- i
	}
	close(ch1)
}

// 限制通道类型
func f2(ch1 <-chan int, ch2 chan<- int) {
	defer wg.Done()
	for {
		x, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- x * x
	}
	//close(ch2)
	once.Do(func() {
		close(ch2)
	}) // 确保某个操作值执行一次
}
