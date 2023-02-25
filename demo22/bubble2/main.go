package main

import (
	"gin_demo/demo22/bubble2/controller"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	//DB s
	DB *gorm.DB
)

// Todo st
type Todo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func initMysql() (err error) {
	dsn := "test:test@tcp(10.3.6.194:3306)/bubble?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return err
	}
	print("xxxx")
	return DB.DB().Ping()
}
func main() {
	err := initMysql()
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	//绑定模型
	DB.AutoMigrate(&Todo{})

	r := gin.Default()
	//待办事项
	//添加
	//查看
	//修改
	//删除
	v1group := r.Group("v1")
	{
		//添加
		v1group.POST("/todo", controller.CreateAtodo)
		//查看所有
		v1group.GET("/todo", func(c *gin.Context) {
			var todoList []Todo
			if err := DB.Find(&todoList).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todoList)
			}
		})
		//查看某一个
		v1group.GET("/todo/:id", func(c *gin.Context) {

		})
		//修改某一个
		v1group.PUT("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			var todo Todo
			if err := DB.Where("id = ?", id).First(&todo).Error; err != nil {
				c.JSON(http.StatusAccepted, gin.H{"error": err.Error()})
				return
			}
			c.BindJSON(&todo)
			if err = DB.Save(&todo).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, todo)
			}
		})
		v1group.DELETE("/todo/:id", func(c *gin.Context) {
			id, ok := c.Params.Get("id")
			if !ok {
				c.JSON(http.StatusOK, gin.H{"error": "无效id"})
				return
			}
			if err := DB.Where("id=?", id).Delete(Todo{}).Error; err != nil {
				c.JSON(http.StatusOK, gin.H{"error": err.Error()})
			} else {
				c.JSON(http.StatusOK, gin.H{id: "deleted"})
			}
		})
	}
	r.Run(":8000")
}
