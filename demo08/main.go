package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/json", func(c *gin.Context) {
		//map
		// data := map[string]interface{}{
		// 	"name": "kkk",
		// 	"age":  "20",
		// }
		// c.JSON(http.StatusOK, data)
		//gin.H
		// data := gin.H{"name": "kkk", "age": 10}
		// c.JSON(http.StatusOK, data)
		//结构体

	})

	type msg struct {
		Name string `json:"name"`
		Age  int
	}
	r.GET("/json2", func(c *gin.Context) {
		data := msg{
			Name: "哇",
			Age:  18,
		}
		c.JSON(http.StatusAccepted, data)
	})
	r.Run(":8000")
}
