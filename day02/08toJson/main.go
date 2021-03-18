package main

import (
	"encoding/json"
	"fmt"
)

//结构体与JSON
//1、序列化：把Go语言中的结构体变量 ---》json格式的字符串
//2、反序列化： json格式字符串 ---》go语言中能够识别的结构体变量

//EmpInfo 人员信息
type EmpInfo struct {
	Name   string `json:"name"`
	Age    int
	Addres []string
}

func main() {
	emp := EmpInfo{
		Name:   "李稻花",
		Age:    18,
		Addres: []string{"广水", "武汉"},
	}
	bey, error := json.Marshal(emp)
	if error != nil {
		fmt.Println("Marshal to json失败", error)
		return
	}
	fmt.Println(string(bey))
}
