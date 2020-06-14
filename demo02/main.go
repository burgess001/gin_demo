package main

import (
	//"github.com/gin-gonic/gin"
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("parse temp failed %v\n", err)
	}
	err = t.Execute(w, "xiaowangzi")
	if err != nil {
		fmt.Println("failed!")
	}
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf(" http server start failed %v\n", err)
	}
}
