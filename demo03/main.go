package main

import (
	//"github.com/gin-gonic/gin"
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//解析模版
	template.ParseFiles("./hello.tmpl")
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse tmplate file error : %v\n", err)
	}
	u1 := User{
		Name:   "小王子",
		Gender: "男",
		Age:    18,
	}
	m1 := map[string]interface{}{
		"name":   "小王子",
		"gender": "男",
		"age":    18,
	}
	hobbyList := []string{
		"篮球",
		"足球",
		"双色球",
	}
	t.Execute(w, map[string]interface{}{
		"u1":    u1,
		"m1":    m1,
		"hobby": hobbyList,
	})
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf(" http server start failed %v\n", err)
	}
}
