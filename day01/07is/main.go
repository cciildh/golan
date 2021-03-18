package main

import "fmt"

var (
	a = 5
	b = 2
)

func main() {
	fmt.Println("+运算：", a+b) //7
	fmt.Println("-运算：", a-b) //3
	fmt.Println("*运算：", a*b) //10
	fmt.Println("/运算：", a/b) //取整2
	fmt.Println("%运算：", a%b) //余数1
}
