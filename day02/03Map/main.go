package main

import "fmt"

func main() {
	//创建map集合
	var ctiysMap map[string]string
	//必须初始化使用make申请map 对象的内存
	ctiysMap = make(map[string]string)
	//map 插入 key -value
	ctiysMap["wuhan"] = "武汉"
	ctiysMap["shanghai"] = "上海"
	ctiysMap["gs"] = "广水"

	//使用健值输出
	fmt.Println("gs键值：", ctiysMap["gs"]) //gs键值： 广水

	//遍历
	for key := range ctiysMap {
		fmt.Println("gs键值：", ctiysMap[key])
	}
	//查看元素是否存在
	Gs, ok := ctiysMap["gs"]
	fmt.Println("ok:", ok) //ok: true
	if ok {
		fmt.Println("gs:", Gs) //gs: 广水
	}

	//删除元素
	delete(ctiysMap, "shanghai")
	for _, v := range ctiysMap {
		fmt.Printf("%v\t", v) //武汉    广水
	}

}
