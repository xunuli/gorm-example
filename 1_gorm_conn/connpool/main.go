package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

func main() {
	//采用gorm连接到mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(192.168.198.130:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}
	fmt.Println(db.Name())
	//获取一个通用数据库对象，sql.DB，然后使用其提供的功能
	sqlDB, err := db.DB()
	//设置空闲连接池中的连接最大数量
	sqlDB.SetMaxIdleConns(10)
	//设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100)
	//设置连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	fmt.Println(sqlDB.Stats())
}
