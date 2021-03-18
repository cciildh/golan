package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsqio/go-nsq"
)

//MyHandler 消费者类型
type MyHandler struct {
	Title string
}

//HandleMessage 需要实现的处理消息方法
func (my *MyHandler) HandleMessage(msg *nsq.Message) (err error) {
	fmt.Printf("%s recv from %v , msg:%v\n", my.Title, msg.NSQDAddress, string(msg.Body))
	return nil
}

//initConsumer 初始化消费者
func initConsumer(topic string, channel string, address string) (err error) {
	config := nsq.NewConfig()
	config.LookupdPollInterval = 15 * time.Second
	consumer, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		fmt.Println("create consumer err ", err)
		return err
	}
	handler := &MyHandler{
		Title: "果果",
	}

	consumer.AddHandler(handler) //将'果果'添加到consumer

	//连接nsq 通过Lookupd查询
	if err := consumer.ConnectToNSQD(address); err != nil {
		return err
	}
	return nil
}
func main() {
	err := initConsumer("guoguo", "first", "192.168.205.188:4150")
	if err != nil {
		fmt.Println("initconsumer err", err)
		return
	}

	c := make(chan os.Signal)        //定义 一个信号通道
	signal.Notify(c, syscall.SIGINT) //转发键盘终端信号到c
	<-c                              //阻塞 （ctrl+c 时退出）
}
