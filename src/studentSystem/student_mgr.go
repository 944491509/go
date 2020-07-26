package main

import "fmt"

// 有一个物体
// 1.保存了一些数据  -->结构体的字段
// 2.有3个功能  --> 结构体的方法

type student struct {
	id   int64
	name string
}

// 学生的管理者
type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudent() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号：%d 姓名：%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	// 1.根据用户的输入创建一个新的学生
	// 2.把新的学生放到s.allStudent这个map里
	var (
		stuID   int64
		stuName string
	)
	// 获取用户输入
	fmt.Println("请输入学号:")
	fmt.Scanln(&stuID)
	fmt.Println("请输入姓名:")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Print("添加成功\n")
}

// 编辑学生
func (s studentMgr) editStudent() {
	// 1.获取输入的学号
	var stuId int64
	fmt.Println("请输入学号:")
	fmt.Scanln(&stuId)
	// 2.展示该学号对应的信息， 如果没有提示查无此人
	stuObj, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人！")
		return
	}
	fmt.Printf("你要修改的学生信息如下：学号：%d  姓名：%s\n", stuObj.id, stuObj.name)
	// 3.请输入修改后的名称
	fmt.Println("请输入学生要修改的名字:")
	var newName string
	fmt.Scanln(&newName)
	// 4.更新学生的姓名
	stuObj.name = newName
	s.allStudent[stuId] = stuObj // 更新map
	fmt.Print("修改成功\n")
}

// 删除学生
func (s studentMgr) delStudent() {
	// 1. 请输入要删除学生的学号
	var (
		stuId int64
	)
	fmt.Print("请输入要删除学生的学号:")
	fmt.Scanln(&stuId)
	// 2.展示该学号对应的信息， 如果没有提示查无此人
	_, ok := s.allStudent[stuId]
	if !ok {
		fmt.Println("查无此人！")
		return
	}
	// 去allStudent这个map中删除学号对应的键值对
	delete(s.allStudent, stuId)
	fmt.Print("删除成功\n")
}

var (
	allStudent map[int64]student
)
