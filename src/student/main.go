package main

import (
	"fmt"
	"os"
)

var (
	allStudent map[int64]*student // 变量声明
)

type student struct {
	id   int64
	name string
}

/**
student类型的构造函数
*/
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	// 打印所有的学生
	for k, v := range allStudent {
		fmt.Printf("学号:%d 姓名:%s\n", k, v.name)
	}
}

func addStudent() {
	// 向allStudent中添加下学生
	// 创建学生
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生的学号:")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名:")
	fmt.Scanln(&name)
	// 调用构造函数
	newStu := newStudent(id, name)
	allStudent[id] = newStu
}

func deleteStudent() {
	// 1. 请输入要删除学生的学号
	var (
		deleteID int64
	)
	fmt.Print("请输入要删除学生的学号:")
	fmt.Scanln(&deleteID)
	// 去allStudent这个map中删除学号对应的键值对
	delete(allStudent, deleteID)
}

func main() {
	allStudent = make(map[int64]*student, 48) // 初始化(开辟内存空间)
	for {
		// 打印菜单
		fmt.Println("欢迎进入学生系统!")
		// 等待用户选择
		fmt.Println(`
			1.查看所有学生
			2.新增学生
			3.删除学生
			4.退出
		`)
		fmt.Print("请输入你的选项:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d这个选项\n", choice)
		// 执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) // 退出
		default:
			fmt.Println("输入错误!")
		}
	}
}
