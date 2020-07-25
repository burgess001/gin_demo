package main

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// User struct
type User struct {
	ID int64
	//Name string
	//Name *string `gorm:"default:'xixi'"`
	Name sql.NullString `gorm:"default:'xixi'"`
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "test:test@tcp(10.3.6.194:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(&User{})
	// u := User{
	// 	Name: "",
	// 	Age:  20,
	// }
	// u := User{
	// 	Name: new(string),
	// 	Age:  20,
	// }
	u := User{
		Name: sql.NullString{String: "", Valid: true},
		Age:  20,
	}
	db.Debug().Create(&u)
}
