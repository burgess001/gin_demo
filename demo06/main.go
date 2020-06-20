package main

import (
	//"github.com/gin-gonic/gin"
	"fmt"
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	//定义模版
	//解析模版
	t, err := template.New("index.tmpl").Delims("{[", "]}").ParseFiles("./templates/index.tmpl")

	// t, err := template.ParseFiles("./templates/index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
		return
	}
	//渲染模版
	msg := "小王子"
	t.Execute(w, msg)
}
func xss(w http.ResponseWriter, r *http.Request) {

	//解析模版之前定义一个自定义函数safe
	t, err := template.New("xss.tmpl").Funcs(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	}).ParseFiles("./templates/xss.tmpl")
	// t, err := template.ParseFiles("./templates/xss.tmpl")
	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
		return
	}
	str1 := "<scrpit> alert(112) </script>"
	str2 := "<a href='http://baidu.com'>baidu</a>"
	t.Execute(w, map[string]string{
		"str1": str1,
		"str2": str2,
	})
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/xss", xss)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf(" http server start failed %v\n", err)
	}
}
