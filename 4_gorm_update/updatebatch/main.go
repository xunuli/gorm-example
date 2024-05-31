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
	db.Model(&User{}).Where("id = ?", "1").Find(&user)
	fmt.Println(user)

	//根据struct更新
	db.Model(&User{}).Where("name = ?", "guoqi").Updates(User{Name: "大雄", Age: 25})
	db.Table("users").Where("age = ?", "25").Scan(&user)
	fmt.Println(user)

	//根据map更新
	db.Table("users").Where("id in ?", []int{1, 3, 5}).Updates(map[string]interface{}{"name": "哆啦", "age": "27"})
	db.Table("users").Where("age = ?", "27").Scan(&user)
	fmt.Println(user)

	//获取记录数
	result := db.Table("users").Where("id in ?", []int{2, 4, 6}).Updates(map[string]interface{}{"name": "努力", "age": "28"})
	db.Table("users").Where("age = ?", "28").Scan(&user)
	fmt.Println(user)
	fmt.Println(result.RowsAffected)

}
