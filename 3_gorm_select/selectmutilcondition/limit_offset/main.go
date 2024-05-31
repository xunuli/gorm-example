package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 定义模型
type User struct {
	ID       int    `gorm:"AUTO_INCREMENT;primaryKey"`
	Name     string `gorm:"default:hahaha"`
	Age      int64  `gorm:"default:18"`
	Birthday time.Time
}

func main() {
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}

	var user []User
	fmt.Println("*************第一次查询******************")
	db.Limit(3).Find(&user)
	//select * from users order by age desc, name
	fmt.Println(user)

	// 通过 -1 消除 Limit 条件
	db.Limit(5).Find(&user).Limit(-1).Find(&user)
	// SELECT * FROM users LIMIT 10; (users)
	// SELECT * FROM users; (users)
	fmt.Println(user)

	//select * from users offset 3;
	db.Limit(10).Offset(3).Find(&user)
	fmt.Println(user)
}
