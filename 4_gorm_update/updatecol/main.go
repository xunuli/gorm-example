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

	var user []User
	db.Where("name = ?", "tecent").Order("id").Find(&user)
	fmt.Println(user)

	//条件更新
	db.Model(&User{}).Where("name = ? and id <= ?", "didi", "9").Update("age", "17").Find(&user)
	//update user set age = 17 where name = "tecent", id <= 10
	fmt.Println(user)
	db.Table("users").Where("name = ? and id > ?", "baidu", "5").Update("age", "17").Scan(&user)
	//db.Table("users").Where("id > ?", "5").Update("age", "17").Find(&user)  出错
	//update user set age = 17 where name = "tecent", id <= 10
	fmt.Println(user)
}
