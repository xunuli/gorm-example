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
	db.Model(&User{}).Where("id in ?", []int{1, 2, 3}).Find(&user)
	fmt.Println(user)

	//删除单条记录
	db.Where("id = ?", "2").Delete(&user)
	//delete from users where id = 2
	db.Model(&User{}).Where("id in ?", []int{1, 2, 3}).Find(&user)
	fmt.Println(user)

	//删除多条记录
	db.Model(&User{}).Where("name = ?", "yanjiusuo").Find(&user)
	fmt.Println(user)
	db.Model(&User{}).Where("name like ?", "yanjiu%").Delete(&user)
	db.Model(&User{}).Where("name = ?", "guoqi").Find(&user)
	fmt.Println(user)
}
