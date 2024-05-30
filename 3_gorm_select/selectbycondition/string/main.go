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

	//定义数据
	var user []User
	//查询全部数据
	//select * from users
	//db.Find(&user)
	//fmt.Println(user)

	//按条件查询
	fmt.Println("********* = 获取第一条 ***********")
	//获取第一条匹配的记录
	//select * from users where name = "didi" limit 1
	db.Where("name = ?", "didi").First(&user)
	fmt.Println(user)
	//获取第一条匹配的记录
	//select * from users where name = "tecent" limit 1
	db.Where("name = ?", "tecent").First(&user)
	fmt.Println(user)

	fmt.Println("********* = 获取所有 ***********")
	//获取匹配的所有记录
	//select * from users where name = "tecent"
	db.Where("name = ?", "tecent").Find(&user)
	fmt.Println(user)

	fmt.Println("********* IN（or） ***********")
	//IN
	//select * from users where age in(18, 22);
	db.Where("age in ?", []int64{18, 22}).Find(&user)
	fmt.Println(user)

	fmt.Println("********* like(模糊查询) ***********")
	//like
	//select * from users where name like %ent%
	db.Where("name like ?", "%ent%").Find(&user)
	fmt.Println(user)

	fmt.Println("********* and(&&) ***********")
	//and
	//select * from users where name = tecent and id >= 7
	db.Where("name = ? and id >= ?", "tecent", 7).Find(&user)
	fmt.Println(user)

	fmt.Println("********* time ***********")
	//time
	//select * from users where Birthday < 2024-05-30 20:06:56.149
	db.Where("birthday < ?", "2024-05-30 20:06:56.149").Find(&user)
	fmt.Println(user)

	fmt.Println("********* between ? and ? ***********")
	//time
	//select * from users where Birthday between 2024-05-30 20:06:56.149 and 2024-05-30 20:06:56.149
	db.Where("birthday between ? and ?", "2024-05-30 20:01:24.770", "2024-05-30 20:06:56.149").Find(&user)
	fmt.Println(user)

}
