package main

import (
	//"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
	}
	msg := "index"
	t.Execute(w, msg)
}

func home(w http.ResponseWriter, r *http.Request) {
	name := "home"
	t, err := template.ParseFiles("home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
	}
	t.Execute(w, name)
}

func index2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/index.tmpl")
	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
	}
	msg := "index2kkk"
	t.Execute(w, msg)

}

func home2(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./templates/base.tmpl", "./templates/home.tmpl")
	if err != nil {
		fmt.Printf("parse template failed err:%v\n", err)
	}
	msg := "蓝色的"
	t.Execute(w, msg)
}
func main() {
	http.HandleFunc("/index", index)
	http.HandleFunc("/home", home)
	http.HandleFunc("/index2", index2)
	http.HandleFunc("/home2", home2)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf(" http server start failed %v\n", err)
	}
}
