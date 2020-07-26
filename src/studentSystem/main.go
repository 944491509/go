package main

import (
	"fmt"
	"os"
)

// 学生管理系统
var smr studentMgr // 声明一个全局的变量学生管理对象

func showMenu() {
	fmt.Println("welcome sms!")
	fmt.Println(`
			1.查看所有学生
			2.新增学生
			3.修改学生
			4.删除学生
			5.退出
		`)
}

func main() {
	smr = studentMgr{ // 修改了全局的变量
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Print("请输入序号:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Println("你输入的是:", choice)

		switch choice {
		case 1:
			smr.showStudent()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.delStudent()
		case 5:
			os.Exit(1)
		default:
			fmt.Println("输入错误！")

		}
	}
}
