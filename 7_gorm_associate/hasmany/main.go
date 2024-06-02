package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 初始化数据库连接
var db *gorm.DB

func init() {
	var err error
	//连接数据库
	//dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}
}

type User struct {
	gorm.Model
	Creditcards []Creditcards
}

type Creditcards struct {
	gorm.Model
	Number string
	UserID uint
}

func main() {
	//自动迁移，建表
	db.AutoMigrate(&User{}, &Creditcards{})

	c1 := Creditcards{
		Number: "123456",
	}
	c2 := Creditcards{
		Number: "654321",
	}

	u := User{
		Creditcards: []Creditcards{c1, c2},
	}

	db.Create(&u)

	var user User
	db.Model(&User{}).Preload("Creditcards").Find(&user)
	fmt.Println(user)

}
