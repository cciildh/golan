package main

import "fmt"

//闭包

//匿名函数
var f1 = func(x, y int) {
	fmt.Println("全局匿名函数", x, y)
}

func main() {
	f1(10, 20) //10 20

	//函数内部无法声明带名字的函数
	f2 := func(x, y int) {
		fmt.Println("局部匿名函数", x, y)
	}
	f2(11, 22)

	//如果只是调用一次的函数，还可以简写成立即执行函数
	func(x, y int) {
		fmt.Println("简写调用函数", x, y)
	}(5, 6)

	//闭包
	ret := f4()
	fmt.Printf("闭包函数返回值函数%T\n", ret) //func(int) int

	fmt.Println("调用", ret(100)) //调用 200

	//递归 阶乘
	fmt.Println(Factorial(4)) //24
}

//闭包是一个函数，这个函数包含了它外部作用域的一个变量，将函数作为返回值
func f4() func(int) int {
	var x = 100
	return func(y int) int {
		x += y
		return x
	}
}

//Factorial 阶乘 递归函数
func Factorial(n int) (result int) {
	if n > 0 {
		result = n * Factorial(n-1)
		return
	}
	return 1
}
