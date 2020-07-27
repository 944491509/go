package main

import "fmt"

type cat struct {
}

type dog struct {
}

type person struct {
}

type speaker interface {
	speak() //
}

func (c cat) speak() {
	fmt.Println("喵喵喵~")
}

func (d dog) speak() {
	fmt.Println("汪汪汪~")
}

func (p person) speak() {
	fmt.Println("啊啊啊~")
}

func da(x speaker) {
	// 接受一个参数,传什么进来那我就打什么
	x.speak()
}

func main() {
	var c cat
	var d dog
	var p person
	da(c)
	da(d)
	da(p)

}
