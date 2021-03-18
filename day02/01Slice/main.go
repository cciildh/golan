//切片
package main

import (
	"fmt"
	"sort"
)

func main() {
	//定义切片数组并初始化
	a1 := []int{1, 4, 3, 6, 5}
	fmt.Println(a1)
	//定义一个存放int类型的切片 默认是nil
	var a2 []int
	fmt.Println(a2)        //[]
	fmt.Println(a2 == nil) //true

	//长度和容量
	fmt.Printf("len(a1):%d cap(a1):%d\n", len(a1), cap(a1)) //len(a1):5 cap(a1):5
	fmt.Printf("len(a2):%d cap(a2):%d\n", len(a2), cap(a2)) //len(a2):0 cap(a2):0

	//切片截取
	//打印子切片从索引1(包含) 到索引3(不包含)
	fmt.Println("a1[1:3]==", a1[1:3]) //a1[1:3]== [4 3]

	//打印子切片从索引0默认下限为 0到3不包含索引3
	fmt.Println("a1[:3]", a1[:3]) //[1 4 3]

	//打印子切片[:] 所有
	fmt.Println("a1[:]", a1[:]) //[1 4 3 6 5]

	//排序
	sort.Ints(a1)   //升序
	fmt.Println(a1) //[1 3 4 5 6]

	sort.Sort(sort.Reverse(sort.IntSlice(a1))) //降序
	fmt.Println(a1)                            //[6 5 4 3 1]

	//make函数创建切片 len(a1)表示使用a1的长度,cap(a1)*2表示使用a1的容量*2
	a3 := make([]int, len(a1), (cap(a1) * 2))
	fmt.Printf("len(a3):%d cap(a3):%d\n", len(a3), cap(a3)) //len(a3):5 cap(a3):10
	//拷贝a1切片的内容到a3
	copy(a3, a1)
	fmt.Println("拷贝a1切片的内容到a3==", a3) //拷贝a1切片的内容到a3== [6 5 4 3 1]

	//向切片添加一个元素 append
	a3 = append(a3, 10)
	fmt.Println("向切片添加一个元素:", a3) // [6 5 4 3 1 10]
	//向切片添加多个元素
	a3 = append(a3, 88, 99)
	fmt.Println("向切片添加多个个元素:", a3) // [6 5 4 3 1 10 88 99]

	//删除切片其中index为4的元素[6 5 4 3 1 10 88 99]
	index := 4
	a3 = append(a3[:index], a3[index+1:]...)
	fmt.Println("删除切片其中index为3的元素:", a3) // [6 5 4 3 10 88 99]

}
