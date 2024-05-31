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
	db.Model(&User{}).Where("name = ?", "meituan").Find(&user)
	fmt.Println(user)

	//删除记录
	db.Model(&User{}).Where("id = ?", "14").Delete(&user)
	db.Model(&User{}).Where("name = ?", "meituan").Find(&user)
	fmt.Println(user)
}

func (u *User) BeforeDelete(tx *gorm.DB) error {
	fmt.Println("正在删除！")
	return nil
}
