package main

import "fmt"

//指针
func main() {

	a := 18
	//&根据变量取地址
	b := &a
	fmt.Println("&取a变量内存地址", b) //0xc0000100d0
	//*：根据变量去职
	fmt.Println("*使用指针访问值", *b) //18

	c := *b
	fmt.Println("使用指针访问值:", c) //18
	//a/b/c内存地址不同
	fmt.Printf("a:%v b:%v c:%v \n", &a, &b, &c) //a:0xc0000a0078 b:0xc0000ca018 c:0xc0000a00b0

	//声明指针变量
	var sor *int
	sor = &a
	fmt.Printf("a地址：%v sor地址：%v\n", &a, sor) //a地址：0xc0000100d0 sor地址：0xc0000100d0
	fmt.Println("使用指针访问值:", *sor)            //18

	//空指针
	var ss *int
	fmt.Println("初始声明指针变量", ss) //<nil>
	//空指针判断：
	fmt.Println(ss == nil) //true

	//new函数申请一个内存地址 防止空指针nil
	var bb = new(int)
	bb = &a
	fmt.Printf("a地址：%v bb地址：%v 值：%v", &a, bb, *bb) //a地址：0xc0000a0078 bb地址：0xc0000a0078 值：18

}
