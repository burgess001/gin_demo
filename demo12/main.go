package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Useinfo 用户
type Useinfo struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

func main() {
	r := gin.Default()
	r.GET("/user", func(c *gin.Context) {
		// name := c.Query("name")
		// password := c.Query("password")
		// u := Useinfo{
		// 	username: name,
		// 	password: password,
		// }
		// fmt.Printf("%#v\n", u)

		var u Useinfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})

	r.POST("/form", func(c *gin.Context) {
		var u Useinfo
		err := c.ShouldBind(&u)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		} else {
			fmt.Printf("%#v\n", u)
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	r.Run(":8000")
}
