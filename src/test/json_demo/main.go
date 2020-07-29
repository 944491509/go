package main

import (
	"encoding/json"
	"fmt"
)

// json

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"刘洋", "age":26}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p)
}
