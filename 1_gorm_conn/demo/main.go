package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// User
type User struct {
	Name   string
	Age    uint
	Gender string
	Hobby  string
	gorm.Model
}

func main() {
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}

	//自动迁移
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("迁移失败：%v", err)
	}
	//初始化
	u1 := User{
		Name:   "xuji",
		Age:    18,
		Gender: "男",
		Hobby:  "篮球",
	}
	u2 := User{
		Name:   "lujie",
		Age:    25,
		Gender: "男",
		Hobby:  "游泳",
	}

	//创建记录，插入数据
	db.Create(&u1)
	db.Create(&u2)

	//Select 查询语句
	var user User
	//根据整型主键ID查找id=1的第一条记录
	db.First(&user, 1)
	fmt.Println(user)
	//查找name为lujie的第一条记录
	var user1 User
	db.Find(&user1, "name = ?", "lujie")
	fmt.Println(user1)

	//Update 更新语句 -将xuji的年龄改为22
	db.Model(&user1).Update("Age", "23")
	fmt.Println(user1)
	//update 更新多个字段
	db.Model(&user1).Updates(User{Name: "lulu", Age: 520})
	fmt.Println(user1)

	//Delete删除语句--删除name为”lullu的记录
	db.Delete(&user, 1)
}
