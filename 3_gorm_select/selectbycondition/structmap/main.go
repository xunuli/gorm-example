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
	//struct
	fmt.Println("********* struct ***********")
	//select * from users where name = "didi" limit 1
	db.Where(&User{Name: "didi"}).First(&user)
	fmt.Println(user)

	//map
	fmt.Println("********* map ***********")
	//select * from users where name = "didi" limit 1
	db.Where(map[string]interface{}{"name": "tecent"}).First(&user)
	fmt.Println(user)

	//主键切片
	fmt.Println("********* 主键切片 ***********")
	//select * from users where id in(7, 17)
	db.Where([]int{7, 17}).Find(&user)
	fmt.Println(user)

}
