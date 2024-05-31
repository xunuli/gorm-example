package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 定义数据库模型
type User struct {
	ID       int `gorm:"primarykey;auto_incrment"`
	Name     string
	Age      int64
	Birthday time.Time
}

func main() {
	//dsn
	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("连接失败: %v", err)
		return
	}

	var user User
	db.Model(&User{}).Where("id = ?", "1").Find(&user)
	fmt.Println(user)

	//更新多列
	db.Model(&User{}).Select("name", "age").Where("id = ?", 6).Updates(map[string]interface{}{"name": "haha", "age": 999})
	db.Model(&User{}).Where("id = ?", "6").Scan(&user)
	fmt.Println(user)
}
