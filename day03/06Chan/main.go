package main

import (
	"fmt"
	"sync"
)

//不带缓冲区的

func main() {
	//创建一个协程容器
	var wg sync.WaitGroup
	//创建一个无有缓存的channel
	ch := make(chan int, 3)
	//len(ch)缓冲区剩余数据个数，cap(ch)缓冲区大小
	fmt.Printf("len(ch)=%d,cap(ch)=%d\n", len(ch), cap(ch))

	wg.Add(1) //
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			fmt.Println("子协程:i", i)
			ch <- i //写入管道
		}
	}()
	wg.Wait()
	
	for i := 0; i < 3; i++ {
		num := <-ch //读取管道中内容，没有内容前，阻塞
		fmt.Println("num =", num)
	}
	fmt.Println("结束")

}
