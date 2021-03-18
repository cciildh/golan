package main

import (
	"encoding/json"
	"fmt"
)

//Empinfo 基本信e息
type empinfo struct {
	Name string
	Age  int
	Pwd  string
}

func main() {
	emp := empinfo{
		Name: "果果",
		Age:  19,
		Pwd:  "123",
	}

	empjs, err := json.Marshal(emp)
	if err != nil {
		return
	}

	fmt.Println(string(empjs))
}
