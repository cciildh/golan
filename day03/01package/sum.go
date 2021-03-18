package sum

import "fmt"

func init() {
	fmt.Println("我是自动执行的")
}

//Add 计算合计
func Add(x, y int) int {
	return x + y
}
