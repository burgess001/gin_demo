package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		//get query string
		//name := c.Query("query")
		//name := c.DefaultQuery("query", "xxxx")
		name, ok := c.GetQuery("query")
		if !ok {
			name = "xxx"
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	})

	r.Run(":8000")
}
