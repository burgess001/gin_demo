package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	//路由组
	videoGroup := r.Group("/video")
	{
		videoGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/index",
			})
		})
		videoGroup.GET("/xxx", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/xxxx",
			})
		})
		videoGroup.GET("/oooo", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "/video/oooo",
			})
		})
	}
	r.Run(":8000")
}
