package main

import (
	"flag"
	"fmt"
	"time"
)

//flag 获取命令行参数
func main() {
	// 创建一个标志位参数
	name := flag.String("name", "刘洋", "请输入名字")
	age := flag.Int("age", 26, "请输入年龄")
	married := flag.Bool("married", false, "结婚了吗")
	cTime := flag.Duration("cTime", time.Second, "时间")
	flag.Parse()
	fmt.Println(*name)
	fmt.Println(*age)
	fmt.Println(*married)
	fmt.Println(*cTime)
	fmt.Printf("%T\n", *cTime)

	//var name string
	//flag.StringVar(&name, "name", "刘洋", "请输入名字")
	//flag.Parse()
	//fmt.Println(name)

	fmt.Println(flag.Args())  // 返回命令行参数后的其它参数, 以[]string类型
	fmt.Println(flag.NArg())  // 返回命令行参数后的求他参数个数
	fmt.Println(flag.NFlag()) // 返回使用的命令行参数个数
}
