package main

import "fmt"

// 地址
type address struct {
	province string
	city     string
}

// 职员
type staff struct {
	name    string
	age     int
	address // 匿名嵌套
}

// 公司
type company struct {
	name string
	addr address
}

func main() {
	staff := staff{
		name: "刘洋",
		age:  26,
		address: address{
			province: "山西省",
			city:     "运城市",
		},
	}
	fmt.Println(staff)
	fmt.Println(staff.name, staff.address.city)
	// 现在自己结构体找这个字段,住不到就去匿名结构体中查找
	fmt.Println(staff.city) // 匿名嵌套
}
