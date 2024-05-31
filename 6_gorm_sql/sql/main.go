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
type Result struct {
	ID   int
	Name string
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
	db.Model(&User{}).Where("name = ?", "meituan").Find(&user)
	fmt.Println(user)

	var res []Result
	db.Model(&User{}).Raw("select * from users where id <= ?", 5).Scan(&res)
	fmt.Println(res)

	var suu int
	db.Raw("select count(*) from users where name =  ?", "哆啦").Scan(&suu)
	fmt.Println(suu)

	db.Exec("update users set age = ? where id = ?", 999, 1)
	db.Find(&user, 1)
	fmt.Println(user)

	db.Exec("update users set age = ? where id = ?", gorm.Expr("age * ? + ?", 0.1, 1), 1)
	db.Find(&user, 1)
	fmt.Println(user)
}
