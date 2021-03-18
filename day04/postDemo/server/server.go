package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type result struct {
	Code int
	Msg  string
	Data []string
}

type empinfo struct {
	Name string
	Pwd  string
	Age  int
}

//Login 登陆服务
func Login(w http.ResponseWriter, r *http.Request) {

	params := r.URL.Query() //获取url 参数
	fmt.Println(params)
	var emp empinfo
	emp.Age = 19
	for key := range params {
		value := params.Get(key)
		// w.Write([]byte(key + ":" + value))
		if key == "username" {
			emp.Name = value
		} else {
			emp.Pwd = value
		}
	}
	fmt.Println(emp)
	bey, _ := json.Marshal(emp)

	fmt.Println(string(bey))
	rest := result{
		Code: 200,
		Msg:  "登陆成功！",
		Data: []string{string(bey)},
	}
	beyRest, _ := json.Marshal(rest)

	fmt.Println(string(beyRest))

	//打印Cookie
	// fmt.Println(r.Cookies())
	// r.ParseForm()

	cookeid, err := r.Cookie("userId")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("cookieid:", cookeid.Name, cookeid.Value)

	w.Write([]byte(beyRest))

}

func main() {
	http.HandleFunc("/login", Login)
	//设置监听的端口
	http.ListenAndServe("0.0.0.0:9090", nil) //全网都可访问
}
