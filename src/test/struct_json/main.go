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

// 沟造（结构体变量的）函数， 返回值是对应的结构体类型
func newPerson(n string, i int) (p person) {
	p = person{
		Name: n,
		Age:  i,
	}
	return p
}

// 方法
// 接收者使用对应类型的首字母小写
// 指定来接收者之后，只有接受者才能调用这个方法
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是%s\n", p.Name, str)
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

	p3 := person{"张朋", 26}
	p4 := newPerson("老曹", 27)
	fmt.Println(p3, p4)
	p1.dream("学好go语言")
	p2.dream("学好java")
}