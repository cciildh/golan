package main

import "fmt"

//结构体

//EmpInfos 定义一个结构体
type EmpInfos struct {
	name   string
	age    int
	addres string
}

//NewEmpInfo 构造函数
func NewEmpInfo(name string, age int, addres string) EmpInfos {
	return EmpInfos{
		name:   name,
		age:    age,
		addres: addres,
	}
}

func main() {
	// var empinfo EmpInfos
	// empinfo.name = "果果"
	// empinfo.age = 18
	// empinfo.addres = "哈尔滨大庆"

	// fmt.Println(empinfo) //{果果 18 哈尔滨大庆}

	// var info1 *EmpInfos
	// info1 = &empinfo
	// fmt.Println(*info1)

	// //参数为结构体的方法
	// f1(*info1) //果果

	//构造函数调用
	newemp := NewEmpInfo("张三", 30, "湖北广水")
	fmt.Printf("%T %v\n", newemp, newemp) //main.EmpInfos {张三 30 湖北广水}

}

func f1(emp EmpInfos) {
	fmt.Println(emp.name)
}
