package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

type User struct {
	ID       int    `gorm:"AUTO_INCREMENT;primaryKey"`
	Name     string `gorm:"default:hahaha"`
	Age      int64  `gorm:"default:18"`
	Birthday time.Time
}
type Result struct {
	Name string
	Age  int
}

func main() {
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}

	var user []User
	db.Distinct("name", "age").Order("age, name").Find(&user)
	//select name, age from users order by age, name
	fmt.Println(user)

	//scan
	var res []Result
	db.Table("users").Select("name", "age").Where("name = ?", "xuji").Scan(&res)
	fmt.Println(res)
	//原生sql
	db.Raw("select name, age from users where name = ?", "lujie").Scan(&res)
	fmt.Println(res)
}
