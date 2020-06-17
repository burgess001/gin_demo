package main

import (
	//"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"text/template"
)

func f1(w http.ResponseWriter, r *http.Request) {
	k := func(name string) (string, error) {
		return name + "年轻又帅气", nil
	}
	//定义模版
	t := template.New("f.tmpl")

	//解析模版
	t.Funcs(template.FuncMap{
		"kua99": k,
	})
	_, err := t.ParseFiles("./f.tmpl")

	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
	}
	//渲染模版
	name := "小王子"
	t.Execute(w, name)
}

func f2(w http.ResponseWriter, r *http.Request) {
	name := "小公主"
	t, err := template.ParseFiles("t.tmpl", "ul.tmpl")
	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
	}
	t.Execute(w, name)
}

func main() {
	http.HandleFunc("/", f1)
	http.HandleFunc("/tmpldemo", f2)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf(" http server start failed %v\n", err)
	}
}
