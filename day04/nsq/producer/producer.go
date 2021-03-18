package main

// 消息生产者 producer
import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/nsqio/go-nsq"
)

//initProducer初始化生产者
func initProducer(addr string) (producer *nsq.Producer, err error) {
	config := nsq.NewConfig()

	producer, err = nsq.NewProducer(addr, config)
	if err != nil {
		fmt.Println("nsq.NewProducer", err)
		return nil, err
	}
	return producer, nil
}

func main() {
	var addr = "192.168.205.188:4150" //使用了tcp监听的端口

	//new NewProducer
	producer, err := initProducer(addr)

	//消息日志
	producer.SetLogger(log.New(ioutil.Discard, "", log.LstdFlags), nsq.LogLevelInfo)
	//ping
	err = producer.Ping()
	if err != nil {
		fmt.Println("producer.Ping", err)
		return
	}
	fmt.Println("producer.Ping", "√")

	//从标准输入读取
	reader := bufio.NewReader(os.Stdin)
	topName := "guoguo" //主题
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read string form stdin faild , err", err)
			continue
		}
		data = strings.TrimSpace(data)
		if strings.ToUpper(data) == "Q" { //输入Q推出
			break
		}
		err = SendMessage(producer, topName, data)
		if err != nil {
			fmt.Println("sendMessage err ", err)
			continue
		}

	}

	defer producer.Stop()

}

//Reader1 组建消息1异步
func Reader1(producer *nsq.Producer) {
	//异步生产消息
	msgCt := 10         // 10个协程
	topName := "guoguo" //主题
	var wg sync.WaitGroup
	wg.Add(msgCt)
	//创建10个groutine
	for i := 0; i < msgCt; i++ {
		msg := "果果" + fmt.Sprintln(i) + ":" + time.Now().Format("2006-01-02 15:04:05") //消息内容
		err := SendMessage(producer, topName, msg)                                     //发送消息
		if err != nil {
			fmt.Println("sendMessage err ", err)
			continue
		}
		defer wg.Done()

		time.Sleep(10 * time.Millisecond) //延时10毫秒
	}

	wg.Wait()

}

//SendMessage 发送消息
//topName 主题
//msg 消息内荣
func SendMessage(producer *nsq.Producer, topName string, msg string) (err error) {
	err = producer.Publish(topName, []byte(msg)) //发送消息

	if err != nil {
		fmt.Println("producer.msg err", err)
		return err
	}
	fmt.Println("producer.Publish", msg, "√")
	return nil
}
