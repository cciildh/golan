package main

import (
	"fmt"
	"time"
)

//GetDays 获取时间差返回天数
func GetDays(start, end string) int64 {
	timeLayout := "2006-01-02"
	loc, _ := time.LoadLocation("Local")
	// 转成时间戳
	startUnix, _ := time.ParseInLocation(timeLayout, start, loc)
	endUnix, _ := time.ParseInLocation(timeLayout, end, loc)
	startTime := startUnix.Unix()
	endTime := endUnix.Unix()
	date := (endTime - startTime) / 86400
	return date
}

func main() {
	date := time.Now()
	fmt.Println("获取年", date.Year())

	fmt.Println("格式化后的当前时间", date.Format("2006-01-02 15:04:05"))

	//定时器
	for i := 0; i < 5; i++ {
		date := time.Now()
		fmt.Println("格式化后的当前时间到毫秒", date.Format("2006-01-02 15:04:05:000"))
		time.Sleep(time.Second * 2)
	}

	//时间修改8天前
	d, _ := time.ParseDuration("-24h")
	date1 := date.Add(8 * d).Format("2006-01-02")
	fmt.Println("时间差8天前", date1)
	//时间差
	fmt.Printf("date1 %v 与 dates%v时间差：%v", date1, date.Format("2006-01-02"), GetDays(date1, date.Format("2006-01-02")))

}
