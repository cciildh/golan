package main

import (
	"encoding/json"
	"fmt"
)

//接口

//EmpService 定义接口
type EmpService interface {
	//方法名 （） 返回参数1，.。。
	GetEmp()
	Save()
}

//EiEmp 定义EiEmp结构体
type EiEmp struct {
	EiName   string
	EiStatus int
	EiRedix  float64
}

//GetEmp EiEmp实现实现接口GetEmp方法
func (ei EiEmp) GetEmp() {
	emp, err := json.Marshal(ei)
	if err != nil {
		return
	}

	fmt.Println("EiEmp:", string(emp))
}

//Save EiEmp实现接口的Save方法
func (ei EiEmp) Save() {
	fmt.Println("eiemp保存成功！")
}

//MiEmp 定义MiEmp结构体
type MiEmp struct {
	MiName   string
	MiStutas int
	MiRedix  float32
}

//GetEmp MiEmp实现实现接口GetEmp方法
//使用指针值接收
func (mi *MiEmp) GetEmp() {
	emp, err := json.Marshal(mi)
	if err != nil {
		return
	}
	fmt.Println("MiEmp:", string(emp))
}

//Save MiEmp实现接口的Save方法
func (mi MiEmp) Save() {
	fmt.Println("miemp保存成功")
}

func main() {

	//接口使用
	var empService EmpService
	empService = EiEmp{
		EiName:   "果果",
		EiRedix:  6500.05,
		EiStatus: 1,
	}
	empService.GetEmp()
	empService.Save()

	//使用地址赋值
	empService = &MiEmp{"小果", 5000, 1}
	empService.GetEmp()
	empService.Save()

}
