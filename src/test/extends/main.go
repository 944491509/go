package main

import "fmt"

// 定义个动物
type animal struct {
	name string
}

// 给动物实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动\n", a.name)
}

// 小狗
type dog struct {
	animal // animal 拥有的方法dog都拥有
	feet   int
}

// 狗在叫
func (d dog) wang() {
	fmt.Printf("%s在叫:汪汪汪~\n", d.name)
}

func main() {
	d1 := dog{
		animal: animal{
			name: "小白",
		},
		feet: 2,
	}
	fmt.Println(d1)
	d1.move()
	d1.wang()
}
