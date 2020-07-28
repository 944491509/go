package main

import (
	"fmt"
	"time"
)

func main() {
	timeLayout := "2006-01-02 15:04:05"

	now := time.Now().Format(timeLayout)
	fmt.Println(now)
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Printf("load loc failed, err:%v  \n", err)
		return
	}
	fmt.Println(loc)
	// 按照指定时区解析时间
	timeObj, err := time.ParseInLocation(timeLayout, "2020-07-28 09:58:00", loc)
	if err != nil {
		fmt.Printf("parse time failed, err:%v\n", err)
		return
	}
	fmt.Println(timeObj)
}
