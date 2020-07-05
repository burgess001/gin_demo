package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//Userinfo userinfo
type Userinfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	db, err := gorm.Open("mysql", "test:test@tcp(10.3.6.194:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	//创建表自动迁移
	db.AutoMigrate(&Userinfo{})

	//创建数据行
	// u1 := Userinfo{
	// 	ID:     0,
	// 	Name:   "koujw",
	// 	Gender: "男",
	// 	Hobby:  "玩",
	// }
	// db.Create(&u1)

	//查询
	var u Userinfo
	db.First(&u)
	fmt.Printf("u:%v\n", u)
	//更新
	db.Model(&u).Update("Gender", "男")
	//删除
	db.Delete(&u)

}
