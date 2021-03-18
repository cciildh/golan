//groutine 并发

package main

import (
	"fmt"
	"sync"
	"time"
)

func hello(i int) {
	fmt.Println(i)
}

var wag sync.WaitGroup

func f1(i int) {
	defer wag.Done()
	time.Sleep(time.Second)
	fmt.Println(i)
}

func main() {

	// for i := 1; i < 100; i++ {
	// 	go hello(i) //开启一个单独的groutine 去执行
	// }

	// for i := 1; i < 10; i++ {
	// 	go func(i int) {
	// 		wag.Add(1)
	// 		fmt.Println(i)
	// 	}(i)
	// }

	fmt.Println("main") //主线程

	// 开启99个groutine
	for i := 1; i < 100; i++ {
		wag.Add(1) // 计数器
		go func(i int) {
			defer wag.Done()
			time.Sleep(time.Second)
			fmt.Printf("%v\t", i)
		}(i)
	}
	wag.Wait()        //等待wag的计数器减为0
	fmt.Println("结束") //主线程
	// time.Sleep(time.Second) //延时等待1秒防止主协程先退出导致子协程没有来得及调用

}
