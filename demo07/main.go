package main

import (
	"html/template"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// r.LoadHTMLFiles("./templates/index.tmpl", "./templates/users/index.tmpl") //parse template
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// load static file
	r.Static("/xxx", "./static")
	r.LoadHTMLGlob("templates/**/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusAccepted, "index.tmpl", gin.H{
			"title": "www.baidu.com",
		})
	})

	r.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusAccepted, "users/index.tmpl", gin.H{
			"title": "users",
		})
	})

	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusAccepted, "posts/index.tmpl", gin.H{
			"title": "<a href='http://baidu.com'> 百度 </a>",
		})
	})
	r.Run(":8000")
}
