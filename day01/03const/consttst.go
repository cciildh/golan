package main

import (
	"fmt"
)

//常量的定义
const pi = 3.14

//批量常量的定义
const (
	success = "ok"
	code    = 200
)

//表示 n2\n3 默认是和n1一样的值
const (
	n1 = 100
	n2
	n3 //100
)

//iota是gO语言计数器，const没新增一行（计数）就自动增加一次
const (
	m1 = iota //0  const iota出现时重置为0
	m2        //1
	m3        //2
)

//匿名
const (
	a1 = iota //0  const iota出现时重置为0
	a2        //1
	_         //2
	a3        //3
)

//插队
const (
	b1 = iota //0
	b2 = 100  //100const没新增一行（计数）就自动增加一次不是值加一
	b3 = iota //2
	b4        //3
)

//多个常量在一行
const (
	c1, c2 = iota + 1, iota + 2 //c1: 0+1  c2 :0+2
	c3, c4 = iota + 1, iota + 2 // c3: 1+1  c4 : 1+2
)

func main() {
	fmt.Println(pi)
	fmt.Println("success:" + success)
	fmt.Println("code", code)

	fmt.Println("n3", n3)
	fmt.Println("m3", m3)

	fmt.Println("匿名 a3", a3)

	fmt.Println("iota插队 b4", b4)

}
