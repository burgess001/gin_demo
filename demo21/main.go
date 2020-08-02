package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User struct
type User struct {
	gorm.Model
	Name   string
	Age    int64
	Active bool
}

func main() {
	db, err := gorm.Open("mysql", "test:test@tcp(10.3.6.194:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	//create
	// u1 := User{Name: "koujw", Age: 18, Active: true}
	// db.Create(&u1)
	// u2 := User{Name: "jinzhu", Age: 18, Active: true}
	// db.Create(&u2)

	//query
	var user User
	db.First(&user)
	user.Name = "koujw"
	user.Age = 22
	db.Debug().Save(&user)
	db.Debug().Model(&user).Update("name", "zhangsan")
	db.Model(&user).Where("active = ?", true).Update("name", "hello")
	m1 := map[string]interface{}{
		"name":   "sanzhang",
		"age":    99,
		"active": false,
	}
	//db.Debug().Model(&user).Updates(m1)
	db.Model(&user).Select("age").Updates(m1) //只更新age字段
}
