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

	//dryrun
	stmt := db.Session(&gorm.Session{DryRun: true}).First(&user, 1).Statement
	ss := stmt.SQL.String()
	fmt.Println(ss, stmt.Vars)

	//
	rows, err := db.Model(&User{}).Where("name = ?", "大雄").Select("name", "age").Rows()
	defer rows.Close()

	for rows.Next() {
		db.ScanRows(rows, &user)
		fmt.Println(user)
	}
}
