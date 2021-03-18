package main

import (
	"fmt"
	"time"
)

func work(poolid int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		fmt.Printf("pool:%d   执行 job：%d\n", poolid, j)
		time.Sleep(time.Second)
		fmt.Printf("pool:%d  完成 job:%d\n", poolid, j)
		result <- j
	}

}

func main() {
	jobs := make(chan int, 100)
	result := make(chan int, 100)
	//开启多个groutine
	for i := 1; i < 4; i++ {
		go work(i, jobs, result)
	}

	//7个任务
	for i := 0; i < 5; i++ {
		jobs <- i
	}
	close(jobs) ////不需要再写数据，关闭channel
	//读取
	for i := 0; i < 5; i++ {
		<-result
	}
}
