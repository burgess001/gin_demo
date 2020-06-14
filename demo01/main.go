package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello golang!",
	})
	//c.HTML(code=200, name string, obj interface{})
}

func main() {
	r := gin.Default() //返回默认的路由引擎
	r.GET("/hello", sayHello)

	r.GET("/book", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"method": "GET",
		})
	})
	r.POST("/book", func(context *gin.Context) {
		context.JSON(http.StatusCreated, gin.H{
			"method": "POST",
		})
	})
	r.PUT("/book", func(context *gin.Context) {
		context.JSON(http.StatusAccepted, gin.H{
			"method": "PUT",
		})
	})
	r.DELETE("/book", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"method": "DELETE",
		})
	})
	r.Run(":8000")
}
