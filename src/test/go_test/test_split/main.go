package main

import (
	"fmt"
	"test/go_test/split_string"
)

func main() {
	ret := split_string.Split("abcdbcdesbcdr", "b")
	fmt.Printf("%#v\n", ret)
}
