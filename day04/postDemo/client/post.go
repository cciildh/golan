package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func main() {
	path := "http://127.0.0.1:9090/login"
	urlStr, _ := url.Parse(path)
	param := url.Values{
		"username": {"果果"},
		"pwd":      {"123"},
	}
	urlStr.RawQuery = param.Encode()
	requst, _ := http.NewRequest("POST", urlStr.String(), nil)

	//cookId 添加cookie
	cookID := &http.Cookie{Name: "userId", Value: "1234"}
	cookName := &http.Cookie{Name: "name", Value: "kongyixueyuan"}
	requst.AddCookie(cookID)
	requst.AddCookie(cookName)

	response, err := http.DefaultClient.Do(requst)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close() //关闭io流
	if response.StatusCode == 200 {
		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
	fmt.Printf("response:%+v\n", response)
	fmt.Printf("response.Header:%+v\n", response.Header)
	fmt.Printf("response.Cookies:%+v\n", response.Cookies())

}
