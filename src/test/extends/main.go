package main

import "fmt"

type animal struct {
}

type dog struct {
	animal
	feet int
}

// 狗在叫
func (d dog) wang() {
	fmt.Println("%s在叫:汪汪汪~", d.name)
}
