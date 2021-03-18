package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// fmttst()
	stringtst()
	// Mathtst()
}

func fmttst() {
	var n = 100
	fmt.Printf("%T\n", n) //int
	fmt.Printf("%v\n", n) //100
	fmt.Printf("%b\n", n) //1100100
	fmt.Printf("%o\n", n) //144
	fmt.Printf("%x\n", n) //64
	var s = "果果"
	fmt.Printf("%s\n", s)  //果果
	fmt.Printf("%#v\n", s) //“果果”

}

func stringtst() {
	//转义符  \\
	var path = "d:go\\src\\05fmt"
	//转义符 \ 添加 """
	var path1 = "\"d:go\\src\\05fmt\""

	fmt.Println(path)  //d:go\src\05fmt
	fmt.Println(path1) //"d:go\src\05fmt"

	//字符串相关操作

	name := "果双成"
	addres := "哈尔滨大庆"
	//fmt.Sprintf  拼串
	s1 := fmt.Sprintf("%s%s", name, addres) //果双成哈尔滨大庆
	fmt.Println(s1)

	//strings用法
	//分割
	ret := strings.Split(path, "\\") //[d:go src 05fmt]
	fmt.Println(ret)
	//包含
	fmt.Println(strings.Contains(s1, "gg")) //false

	//join
	fmt.Println(strings.Join(ret, "+")) //d:go+src+05fmt

	//转大写
	fmt.Println(strings.ToUpper(path)) //D:GO\SRC\05FMT
	//转小写
	fmt.Println(strings.ToLower(path))

	//%c:而C中定义了一些字母前加"\"来表示常见的那些不能显示的ASCII字符，如\0,\t,\n等，就称为转义字符，因为后面的字符，都不是它本来的ASCII字符意思了。
	for i, c := range addres {
		fmt.Printf("%v%c\n", i, c)
	}

	for _, v := range addres {
		fmt.Printf("%c\n", v) //_ 表示匿名变量 防止不使用时报错
	}

	//字符串与字符
	str1 := "果果"
	str2 := '果'

	fmt.Printf("s1：%T s2：%T\n", str1, str2)
	fmt.Printf("s1：%v s2：%v\n", str1, str2)

	//修改字符串
	str3 := []rune(str1) //把字符强制转换成rune切片
	str3[1] = '成'

	fmt.Println("rune切片 ", str3)
	fmt.Println("rune转化为字符串 ", string(str3))

	//类型转换
	str4 := 10 //int
	// var str5 float64
	str5 := float64(str4)
	fmt.Printf("type:%T value:%v\n", str5, str5) //type:float64 value:10

}

//Mathtst getmaxint32
func Mathtst() {
	s1 := math.MaxInt32 //2147483647

	fmt.Println(s1)

}

//Convts ces
func Convts() {
	age := 18
	//数字转字符串
	ret1 := strconv.Itoa(age)
	fmt.Println(ret1)

	//字符串转数字
	ret2, error := strconv.Atoi("100")
	if error == nil {
		fmt.Println("Atoi", ret2)
	}
}
