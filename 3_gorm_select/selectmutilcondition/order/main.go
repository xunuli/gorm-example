package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
	"time"
)

// 定义模型
type User struct {
	ID       int    `gorm:"AUTO_INCREMENT;primaryKey"`
	Name     string `gorm:"default:hahaha"`
	Age      int64  `gorm:"default:18"`
	Birthday time.Time
}

func main() {
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}
	var user []User
	//排序查询
	fmt.Println("*************第一次查询******************")
	db.Order("age desc, name").Find(&user)
	//select * from users order by age desc, name
	fmt.Println(user)

	fmt.Println("*************第二次查询*********************")
	db.Order("age desc").Order("id").Find(&user)
	//select * from users order by age desc, name
	fmt.Println(user)

	//
	fmt.Println("*************第三次查询*********************")
	db.Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{[]int{1, 2, 3}}, WithoutParentheses: true},
	}).Find(&User{})
	// SELECT * FROM users ORDER BY FIELD(id,1,2,3)
	fmt.Println(user)

}
