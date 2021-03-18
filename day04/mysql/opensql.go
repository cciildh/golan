package main

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//StdUser  用户信息
type StdUser struct {
	ID      int
	Name    string
	Age     int
	SexName string
	Address string
	Phone   string
}

//Dbcon 打开数据库连接
func main() {
	db, err := sql.Open("mysql", "ldh:ldh@tcp(192.168.205.136:3306)/firstDB?charset=utf8")
	if err != nil {
		fmt.Println("open mysql failed,", err)
		return
	}

	rows, err := db.Query("SELECT * FROM std_emp")
	if err != nil {
		fmt.Println("查询有误。。")
		return
	}
	//创建slice，存入struct，
	datas := make([]StdUser, 0)
	for rows.Next() {
		var id int
		var name string
		var age int
		var sexName string
		var address string
		var phone string
		err := rows.Scan(&id, &name, &age, &sexName, &address, &phone)
		if err != nil {
			fmt.Println("获取失败。。")
		}
		//每读取一行，创建一个user对象，存入datas2中
		user := StdUser{id, name, age, sexName, address, phone}
		datas = append(datas, user)
	}

	rows.Close()
	db.Close()

	// for _, v := range datas {
	// 	fmt.Println(v)
	// }

	tojson, err := json.Marshal(datas)
	if err != nil {
		fmt.Println("json格式化失败")
	}
	fmt.Println(string(tojson))

}
