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
	//连接数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}
	var user []User
	//根据字段查询
	fmt.Println("********查询1************")
	db.Select("name", "age").Find(&user)
	fmt.Println(user)
	fmt.Println("********查询2************")
	db.Select([]string{"name", "age"}).Find(&user)
	fmt.Println(user)
	fmt.Println("********查询3************")
	result, err := db.Table("users").Select("COALESCE(age)").Rows()
	fmt.Println(result)
}
