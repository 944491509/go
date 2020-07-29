package main

import "fmt"

// 数组
// 存放元素的容器
// 必须指定存放的元素的类型和容量(长度)
// 数组长度是数组类型的一部分
func main() {
	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]

	fmt.Printf("a1:%T  a2:%T\n", a1, a2)

	// 数组初始化
	// 如果不初始化,默认元素都是零值(布尔值:false,整形和浮点型都是0,字符串:"")
	fmt.Println(a1, a2)
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)

	a3 := [9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a3)
	// 初始化方式2
	//根据初始值自动推算数组长度是多少
	a4 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 1, 2, 3}
	fmt.Println(a4)
	// 初始化方式3  根据索引来初始化
	a5 := [5]int{1: 2, 4: 3}
	fmt.Println(a5)

	// 数组的遍历
	citys := [...]string{"北京", "上海", "广州"}
	// 1根据索引遍历
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	// 2 for range遍历
	for i, v := range citys {
		fmt.Println(i, v)
	}

	// 多维数组 [[1,2],[3,4],[5,6]]
	var a11 [3][2]int
	a11 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)

	// 多维数组的遍历
	for _, v1 := range a11 {
		fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Println(v2)
		}
	}
	// 数组是值类型
	b1 := [3]int{1, 2, 3} // b1 [1,2,3]
	b2 := b1
	b2[0] = 100 // b2: [100,2,3]
	fmt.Println(b1, b2)

}
