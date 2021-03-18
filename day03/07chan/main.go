package main

import (
	"fmt"
	"sync"
)

//单方向的channel
func main() {
	const cap = 5 //定义通道容量
	var wg sync.WaitGroup
	//创建一个双向通道
	ch := make(chan int, cap)

	//新开一个协程
	//生产者，生产数字，写入channel
	wg.Add(1)
	go Producer(cap, &wg, ch)
	wg.Wait()
	//消费者，从channel读内容
	Consumer(cap, ch)

}

//Producer 此channel只能写
func Producer(cap int, wg *sync.WaitGroup, in chan<- int) {
	defer wg.Done()
	for i := 0; i < cap; i++ {
		in <- i
	}
	close(in)
}

//Consumer 此channel 只能读
func Consumer(cap int, out <-chan int) {
	for i := 0; i < cap; i++ {
		num := <-out
		fmt.Println("num=", num)
	}
}
