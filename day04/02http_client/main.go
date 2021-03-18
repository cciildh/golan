package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {

	//组建url请求地址
	path := "http://127.0.0.1:9090/login"
	urlobj, _ := url.Parse(path)
	//组件参数
	data := url.Values{}
	data.Set("name", "果双城")
	data.Set("password", "123456")
	querstr := data.Encode() //url encode格式化之后的url
	fmt.Println("格式化之后的url:", querstr)
	urlobj.RawQuery = querstr
	req, _ := http.NewRequest("GET", urlobj.String(), nil)
	fmt.Println("url:", urlobj.String())
	//创建client对象
	response, err := http.DefaultClient.Do(req)
	// clien := http.Client{}
	//调用get方法，发送请求
	// response, err := clien.Get(urlobj.String())
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}

	fmt.Printf("response.Body：%+v\n", response.Body)
	fmt.Printf("response.Header：%+v\n", response.Header)
	fmt.Printf("response.StatusCode：%+v\n", response.StatusCode)
	fmt.Printf("response.Status：%+v\n", response.Status)
	fmt.Printf("response.Request：%+v\n", response.Request)
	fmt.Printf("response.Cookies：%+v\n", response.Cookies())

}
