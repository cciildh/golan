package main

import "fmt"

//函数
func main() {
	Sum(3, 4)
	//接收函数返回值
	name, age := returns("果果", 18)
	fmt.Println(name, age) //果果 18

	//函数中使用defer延时关键字
	deferDome()

}

// func function_name( [parameter list] ) [return_types] {
// 	函数体
//  }

//Sum 计算和
func Sum(a, b int) int {
	return a + b
}

//Sum1 return sum
func Sum1(a, b int) (sum int) {
	sum = a + b
	return
}

//Print 打印信息
func Print(name string, age int) {
	fmt.Println(name, age)
}

//函数返回多个值
func returns(name string, age int) (string, int) {
	return name, age
}

func deferDome() {
	fmt.Println("开始了")
	defer fmt.Println("defer") //前面加上defer表示该语句最后执行
	fmt.Println("end")
}
