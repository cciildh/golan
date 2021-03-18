package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
func defualt(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("hllo word"))
}
func main() {

	//设置路由
	http.HandleFunc("/GetEmp", GetEmp)
	http.HandleFunc("/GetListEmp", GetListEmp)

	http.HandleFunc("/", defualt)

	//设置监听端口
	http.ListenAndServe("192.168.205.177:80", nil)
	defer Db.Close()
}

//GetEmp 获取人员信息
func GetEmp(w http.ResponseWriter, r *http.Request) {

	parms := r.URL.Query()
	userID := parms.Get("ID") //获取人员ID参数
	var user StdUser

	err := Db.Get(&user, "SELECT * from std_emp where id=?", userID)
	if err != nil {
		fmt.Println("db.get err ", err)
		return
	}
	//json 转换
	bey, err := json.Marshal(user)
	if err != nil {
		fmt.Println("objct to json err", err)
		return
	}
	//response
	w.Write(bey)

}

//GetListEmp 获取所有人员
func GetListEmp(w http.ResponseWriter, r *http.Request) {
	var users []StdUser
	err := Db.Select(&users, "select * from std_emp")
	if err != nil {
		fmt.Println("select err", err)
		return
	}

	//json 转换
	bey, err := json.Marshal(users)
	if err != nil {
		fmt.Println("objct to json err", err)
		return
	}
	//response
	w.Write(bey)

}
