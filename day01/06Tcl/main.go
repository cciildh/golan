package main

import (
	"fmt"
)

func main() {
	forTs()
}

func ifelseTs() {
	style := "帅的惊动党中央"
	name := "ldh"
	if name == "ldh" {
		fmt.Printf("%v%v\n", name, style) //ldh帅的惊动党中央
	} else if name == "guo" {
		fmt.Println("高富帅")
	} else { //注意此处不能换行
		fmt.Printf("继续努力")
	}

	if age := 19; age >= 18 {
		fmt.Println("已经成年")
	}
	//age此时已经销毁

	switch name {
	case "ldh":
		fmt.Printf("%v%v\n", name, style)
		break
	case "guo":
		fmt.Println("高富帅")
	default:
		fmt.Printf("继续努力")
	}

}

func forTs() {
	//基本用法
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", i)
	}
	//while语句形式
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum) //16
	// 这样写也可以，更像 While 语句形式
	// for sum1 := 1; sum1 <= 10; {
	// 	sum1++
	// 	fmt.Println("果果") //循环输出10次
	// }
	//无限循环:
	// for sum := 1; ; {
	// 	sum++
	// 	fmt.Println(sum)
	// 	//延时2秒执行
	// 	time.Sleep(time.Second * 2) //1000*5
	// }

	//for range 循环  遍历数据
	// name := "ldh果果"
	// for i, v := range name {

	// 	fmt.Printf("index:%d value:%c\n", i, v)
	// }

	//练习陈发口诀
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {

			fmt.Printf("%d*%d=%v\t", i, j, (i * j))
		}
		fmt.Println()
	}
}
