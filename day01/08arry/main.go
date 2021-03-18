package main

import (
	"fmt"
	"sort"
)

func main() {
	//数组的初始化：默认元素都是零值（bool：false,整型浮点型都是0，字符串是："" )
	var a1 [2]int
	//定义时给定初始值
	a2 := [3]int{66, 55, 77}
	//动态初始值
	var citys = [...]string{"武汉", "广水", "哈尔滨"}
	//数组的长度是类型的一部分
	fmt.Printf("a1:%T a2:%T\n", a1, a2) //type :a1:[2]int a2:[3]int
	fmt.Println("初始化值：", a1, a2)        //初始化值： [0 0] [55 66 77]

	//数组遍历
	for _, v := range citys {
		fmt.Printf(" 遍历结果:%v\t", v) //遍历结果:武汉   遍历结果:广水   遍历结果:哈尔滨
	}

	//根据索引来变量
	for i := 0; i < len(citys); i++ {
		fmt.Printf(" 索引遍历：%v", citys[i]) //索引遍历：武汉 索引遍历：广水 索引遍历：哈尔滨
	}

	//2维数组
	var a3 = [3][2]int{
		{1, 2},
		{3, 4},
		{4, 5},
	}

	fmt.Println("二维数组：", a3) //二维数组： [[1 2] [3 4] [4 5]]

	//数组的排序sort 包
	sort.Ints(a2[:]) //升序  a2[:]转换成切片
	fmt.Println(a2)
	sort.Sort(sort.Reverse(sort.IntSlice(a2[:]))) //降序
	fmt.Println(a2)

}
