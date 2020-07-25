package main

import (
	"encoding/json"
	"fmt"
)

// 结构体与json

// struct 转换json
// json 转化 struct

type person struct {
	Name string `json:"name" db:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "刘洋",
		Age:  26,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed ,err:%v", err)
		return
	}
	fmt.Printf("%v\n", string(b))
	// 反序列化
	str := `{"name":"文豪", "age":24}`
	var p2 person
	// 传指针是为了能在json.Unmarshal内部修改p2的值
	json.Unmarshal([]byte(str), &p2)
	fmt.Printf("%v\n", p2)
}
