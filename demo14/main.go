package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//gin 重定向
	r.GET("index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogou.com")
	})
	r.GET("/a", func(c *gin.Context) {
		//跳转到b对应的路由函数
		c.Request.URL.Path = "/b"
		r.HandleContext(c)

	})
	r.GET("/b", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "b ok",
		})
	})
	r.Run(":8000")
}
