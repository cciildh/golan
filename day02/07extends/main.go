package main

import "fmt"

//Animal 动物
type Animal struct {
	name string
	id   int
}

//Cat 猫
type Cat struct {
	Animal
}

//Call 给Animal实现一个物叫的方法
//类似于GO 接口实现
func (a Animal) Call() {
	fmt.Printf("%v 发出了叫声\n", a.name)
}

//CatCall 给Cat实现一个喵喵
func (c Cat) CatCall() {
	fmt.Printf("%v号%v 喵喵喵！！\n", c.id, c.name)
}
func main() {
	c1 := Cat{
		Animal{"哆啦A梦", 1},
	}
	fmt.Println(c1)
	c1.Call()
	c1.CatCall()
}
