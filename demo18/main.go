package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/jinzhu/gorm"
)

// User struc
type User struct {
	gorm.Model   //内嵌gorm.model
	Name         string
	Age          sql.NullInt64 //零值类型
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(120);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` //设置会员号(member number)唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  //设置num为自增类型
	Adderess     string  `gorm:"index:addr"`      //给address字段创建名为addr索引
	IgnoreMe     int     `gorm:"-"`               //忽略本字段
}

// Hasid test
type Hasid struct {
	ID   string // 名为`ID`的字段会默认作为表的主键
	Name string
}

// Noid test
type Noid struct {
	Name string
}

//有id字段时默认id字段为主键，
//没有id字段时默认没有主键，主键需要自行指定

// Animal 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

func main() {

	db, err := gorm.Open("mysql", "test:test@tcp(10.3.6.194:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	//设置默认表名前缀
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return "prefix_" + defaultTableName
	}
	//db.AutoMigrate(&User{}, &Hasid{}, &Noid{})
	var u User
	db.First(&u)
	fmt.Printf("u:%v\n", u)
}
