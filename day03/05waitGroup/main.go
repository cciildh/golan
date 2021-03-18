package main

import (
	"fmt"
	"sync"
)

// func waitTst(i int, wg *sync.WaitGroup, ch chan int) {
// 	defer wg.Done()
// 	time.Sleep(time.Second) //延时1秒
// 	fmt.Printf("%v %v\t", ch, i)

// }

// func a() {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("A", i)
// 	}
// }

// func b() {
// 	defer wg.Done()
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("B", i)
// 	}
// }

// var wg sync.WaitGroup

func main() {
	fmt.Println("开始")
	var wg sync.WaitGroup

	// var ch chan int
	// ch = make(chan int, 10) //带缓冲区的通道

	ch1 := make(chan int) //不带缓冲区的通道

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go Inchan1(i, &wg, ch1)
	}
	wg.Wait()

	//读取
	for i := 0; i < 10; i++ {
		num := <-ch1 //读取管道中内容，没有内容前，阻塞
		fmt.Println("num =", num)
	}
	fmt.Println("结束")
}

//Inchan1 写入
func Inchan1(i int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	fmt.Println("子协程：", i)
	ch <- i

}
