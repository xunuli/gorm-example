package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 数据库表模型
type User struct {
	ID       int `gorm:"AUTO_INCREMENT;primaryKey"`
	Name     string
	Age      int64
	Birthday time.Time
}

func main() {
	//数据库路径
	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}

	var user []User
	result := db.Find(&user)
	//select * from user
	fmt.Println(user)
	fmt.Println(result)
	fmt.Println(result.RowsAffected)
	fmt.Println(result.Error)
}
