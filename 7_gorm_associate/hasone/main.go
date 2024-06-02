package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// 初始化数据库连接
var db *gorm.DB

func init() {
	var err error
	//连接数据库
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}
}

type User struct {
	CreditCard creditCard
	gorm.Model
}

type creditCard struct {
	UserID int
	Number string
	gorm.Model
}

func main() {
	//需要先创建users表，再创建cresdicard表
	err := db.AutoMigrate(&User{}, &creditCard{})
	if err != nil {
		log.Fatalf("表迁移失败: %v", err)
		return
	}

	//创建记录
	c := creditCard{
		Number: "1234565",
	}
	u := User{
		CreditCard: c,
	}
	db.Create(&u)

	//查询记录
	var us User
	db.Model(&User{}).Preload("CreditCard").First(&us)
	fmt.Println(us)
}
