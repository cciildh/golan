package main

//cluster 集群简单用法
import (
	"fmt"
	"time"

	"github.com/gitstliu/go-redis-cluster"
)

// type Options struct {
//     StartNodes	    []string	    // Startup nodes

//     ConnTimeout	    time.Duration   // Connection timeout
//     ReadTimeout	    time.Duration   // Read timeout
//     WriteTimeout    time.Duration   // Write timeout

//     KeepAlive	    int		    // Maximum keep alive connecion in each node
//     AliveTime	    time.Duration   // Keep alive timeout
// }

func main() {
	fmt.Println("")
	cluster, err := redis.NewCluster(
		&redis.Options{
			StartNodes:   []string{"192.168.205.178:7001", "192.168.205.178:7002", "192.168.205.178:7003", "192.168.205.178:7004", "192.168.205.178:7005", "192.168.205.178:7006"},
			ConnTimeout:  time.Millisecond * 50, //设置连接超时时间 50 毫秒
			ReadTimeout:  time.Millisecond * 50,
			WriteTimeout: time.Millisecond * 50,
			KeepAlive:    16,
			AliveTime:    time.Second * 60,
		})
	if err != nil {
		fmt.Println("redis conn errr", err)
		return
	}
	defer cluster.Close()

	cluster.Do("set", "name", "果果到此一游")

	name, _ := redis.String(cluster.Do("get", "name"))
	fmt.Printf("get redis cluster name :%s", name)

}
