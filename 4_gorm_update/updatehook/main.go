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

	var user User
	db.Model(&User{}).Where("id = ?", "6").Find(&user)
	fmt.Println(user)

	//更新
	db.Table("users").Where("id = ?", "6").Updates(map[string]interface{}{"name": "我一定行", "age": 777})
	db.Table("users").Where("id = ?", "6").Scan(&user)
	fmt.Println(user)

	//只有model的时候才会执行钩子函数
	db.Model(&User{}).Where("id = ?", "6").Update("age", "787")
	db.Table("users").Where("id = ?", "6").Scan(&user)
	fmt.Println(user)
}

// 更新hook
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	fmt.Println("数据库正在更新！")
	return nil
}
