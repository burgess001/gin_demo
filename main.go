package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	b, _ := ioutil.ReadFile("./hello.txt")

	//_,_ = fmt.Fprintln(w, "<h1>hello Golang</h1>")
	_, _ = fmt.Fprintln(w, string(b))
}
func main() {
	http.HandleFunc("/hello", sayhello)

	err := http.ListenAndServe("127.0.0.1:8000", nil)
	if err != nil {
		fmt.Printf("http serve failed ,err:%v", err)
	}
	//http.ListenAndServe(addr string, handler http.Handler)
	//http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
}
