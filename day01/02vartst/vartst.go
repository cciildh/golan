package main

import "fmt"

//批量声明变量
var (
	name string
	age  int
	isOk bool
)

func sum() int {
	return 10
}

func main() {
	var sex string
	sex = "女"
	//变量必须先声明后使用
	fmt.Println(sex)

	name = "果果"
	age = 18
	isOk = true

	fmt.Print(isOk)
	fmt.Printf("name:%s", name) //%s占位符
	fmt.Println(age)
	//匿名变量：定义返回值的函数可以用匿名变量接收防止函数返回值没被使用而报错
_:
	sum()
}
