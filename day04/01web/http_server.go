package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	str := "李稻花到此一游"
	res.Write([]byte(str))

}

func hello1(w http.ResponseWriter, r *http.Request) {
	//需要先编译不然会找不到文件
	f, err := ioutil.ReadFile("./server.html")
	if err != nil {
		w.Write([]byte(fmt.Sprintf("%v", err)))
	}
	w.Write(f)
}

//Login 登陆
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	fmt.Println("URL:", r.URL)
	fmt.Println(r.URL.Query()) //自动帮我们识别url 中的 query parm

	fmt.Println(ioutil.ReadAll(r.Body))

	//输出
	// var parm map[string]string

	parms := r.URL.Query()
	for key := range parms {
		value := parms.Get(key)
		fmt.Printf("%v: %v ", key, value)
		w.Write([]byte(key + ":" + value))
	}
	w.Write([]byte("登陆成功！"))
}

func main() {
	http.HandleFunc("/hello", hello) // //设置访问的路由
	http.HandleFunc("/hello1", hello1)

	http.HandleFunc("/login", Login) //登陆路由
	// http.ListenAndServe("127.0.0.1:9090", nil)//本地ip
	// http.ListenAndServe("192.168.205.177:9090", nil)//本机ip
	//设置监听的端口
	http.ListenAndServe("0.0.0.0:9090", nil) //全网都可访问
}
