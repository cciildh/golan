package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//StdUser  用户信息
type StdUser struct {
	ID      int
	Name    string
	Age     int
	SexName string `db:"sexName"`
	Address string
	Phone   string
}

//Db sql链接
var Db *sqlx.DB

func init() {
	dataSource := "ldh:ldh@tcp(192.168.205.136:3306)/firstDB?charset=utf8"
	db, err := sqlx.Connect("mysql", dataSource)
	if err != nil {
		fmt.Println("mysql open error", err)
		return
	}
	db.SetMaxOpenConns(10) //最大连接数
	db.SetMaxIdleConns(5)  //最大空闲数
	Db = db
}
func main() {
	var users []StdUser
	err := Db.Select(&users, "select * from std_emp")

	if err != nil {
		fmt.Println("select err", err)
	}
	fmt.Printf("this is Select res:%v\n", users)

	var user StdUser
	err1 := Db.Get(&user, "SELECT * from std_emp where id=?", 3)
	if err1 != nil {
		fmt.Println("GET error :", err1)
	} else {
		fmt.Printf("this is GET res1:%v", user)
	}
	defer Db.Close()
}
