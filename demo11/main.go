package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	//get uri param
	r := gin.Default()
	r.GET("/user/:name/:age", func(c *gin.Context) {
		//获取路径参数
		name := c.Param("name")
		age := c.Param("age")
		c.JSON(http.StatusOK, gin.H{
			"name": name,
			"age":  age,
		})
	})
	r.GET("/blog/:year/:mounth", func(c *gin.Context) {
		year := c.Param("year")
		mounth := c.Param("mounth")
		c.JSON(http.StatusAccepted, gin.H{
			"year":   year,
			"mounth": mounth,
		})
	})
	r.Run(":8000")
}
