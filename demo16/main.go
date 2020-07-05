package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//Userinfo userinfo
type Userinfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Info string `json:"info"`
}

//indexHandler
func indexHandler(c *gin.Context) {
	fmt.Println("index ...")
	userinfo := Userinfo{
		Name: "koujw",
		Age:  10,
		Info: "index",
	}
	c.JSON(http.StatusAccepted, userinfo)
}

//m1 定义一个中间件m1,统计耗时
func m1(c *gin.Context) {
	fmt.Println("m1 in ...")
	//计时
	start := time.Now()
	c.Next() //调用后续的处理函数
	//c.Abort() //阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost %v\n", cost)
	fmt.Println("m1 out ...")
}

func m2(c *gin.Context) {
	fmt.Println("m2 in ...")
	c.Next()
	fmt.Println("m2 out ...")
}

func authMiddleware(doCheck bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("in auth ...")
		if doCheck {
			//做认证逻辑
		}
	}
}
func main() {
	r := gin.Default()
	//为全局注册中间件
	r.Use(m1, m2, authMiddleware(true))
	r.GET("/index", indexHandler)
	r.GET("shop", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "shop",
		})
	})
	r.GET("/user", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "user",
		})
	})
	//路由组注册中间件方法1
	xxGroup := r.Group("/xxx", authMiddleware(true))
	{
		xxGroup.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"msg": "/xxx/index",
			})
		})
	}
	//路由组注册中间件方法2
	xx2Group := r.Group("/xx2")
	xx2Group.Use(authMiddleware(true))
	{
		xx2Group.GET("/index", func(c *gin.Context) {
			c.JSON(http.StatusCreated, gin.H{
				"msg": "/xx2/index",
			})
		})
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "page not found",
		})
	})
	r.Run(":8000")
}
