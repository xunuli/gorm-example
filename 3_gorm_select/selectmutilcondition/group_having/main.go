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

type result struct {
	Name  string
	Total int
}

func main() {
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}
	var res []result
	db.Model(&User{}).Select("name, count(age) as total").Where("name like ?", "xu%").Group("name").Find(&res)
	//select name, count(age) as total from users where name like xu% group by name
	fmt.Println(res)

	db.Model(&User{}).Select("name, count(age) as total").Group("name").Find(&res)
	//select name, count(age) as total from users group by name
	fmt.Println(res)

	db.Model(&User{}).Select("name, count(age) as total").Group("name").Having("count(age)>3").Find(&res)
	//select name, count(age) as total from users group by name having count(age)>3
	fmt.Println(res)
}
