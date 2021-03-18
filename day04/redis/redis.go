package main

//直连
import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

//Pool 连接池
var Pool redis.Pool

func init() {
	Pool = redis.Pool{
		MaxIdle:     16, //池中的最大空闲连接数
		MaxActive:   32, //在给定时间池分配的最大连接数。
		IdleTimeout: 120,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "192.168.205.178:6379")
		},
	}
}
func main() {

	//network tcp  addres 192.168.205.178
	conn := Pool.Get()
	defer conn.Close() //关闭连接
	_, err := conn.Do("SET", "address", "湖北广水")
	if err != nil {
		fmt.Println("set redis error", err)
	}
	address, err := redis.String(conn.Do("GET", "address"))
	if err != nil {
		fmt.Println("get redis error", err)
	} else {
		fmt.Printf("Got address: %s\n", address)
	}

}
