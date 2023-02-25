package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAtodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	if err = DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}
