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

	//定义uers切片数据，保存查询结果
	var user []User

	fmt.Println("********* 取反 not ***********")
	//select * from users where not name = "yanjiusuo" limit 1
	db.Not("name = ?", "yanjiusuo").First(&user)
	fmt.Println(user)

	fmt.Println("********* 取反 not in ***********")
	//select * from users where name not in("guoqi", "yanjiusuo")
	db.Not(map[string]interface{}{"name": []string{"guoqi", "yanjiusuo"}}).Find(&user)
	fmt.Println(user)

	fmt.Println("********* 取反 struct ***********")
	//select * from users where name not
	db.Not(User{Name: "guoqi"}).Find(&user)
	fmt.Println(user)

	fmt.Println("********* 不在切片中 ***********")
	//select * from users where name not
	db.Not([]int{8, 16}).Find(&user)
	fmt.Println(user)
}
