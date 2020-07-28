package main

import "fmt"

type liuyang struct {
	name string
}

type wenhao struct {
	name string
}

type saipeng struct {
	name string
}

type eat interface {
	eating()
}

func (l liuyang) eating() {
	fmt.Printf("%v吃的土豆粉!\n", l.name)
}

func (w wenhao) eating() {
	fmt.Printf("%v吃的黄焖鸡!\n", w.name)
}

func eater(e eat) {
	e.eating()
}

func main() {
	l := liuyang{name: "刘洋"}
	w := wenhao{name: "马文豪"}
	//s := saipeng{name:"李赛鹏"}
	eater(l)
	eater(w)
	//eater(s)
}
