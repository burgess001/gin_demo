package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//User struct
type User struct {
	Name string `gorm:"primary_key"`
	Age  int32
}

//Userinfo struct
type Userinfo struct {
	gorm.Model
	Name string
	Age  int32
}

func main() {
	db, err := gorm.Open("mysql", "test:test@tcp(10.3.6.194:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{}, &Userinfo{})
	//create
	// u1 := User{Name: "yangchaoyue", Age: 20}
	// u2 := User{Name: "caixukun", Age: 30}
	// db.Create(&u1)
	// db.Create(&u2)
	// u3 := Userinfo{Name: "yangchaoyue", Age: 20}
	// u4 := Userinfo{Name: "caixukun", Age: 30}
	// db.Create(&u3)
	// db.Create(&u4)
	//query
	var user User
	db.First(&user) //first
	fmt.Printf("user:%#v\n", user)
	var users []User
	db.Where("name IN(?)", []string{"yangchaoyue", "caixukun"}).Find(&users)
	fmt.Printf("user:%#v\n", users)

	var userinfos []Userinfo
	db.Debug().Where([]int64{2, 3, 4}).Find(&userinfos)
	fmt.Printf("userinfos:%#v", userinfos)
}
