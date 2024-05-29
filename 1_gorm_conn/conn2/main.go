package main

import (
	sql2 "database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	//采用gorm连接到mysql
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(192.168.198.130:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	//新建一个sqldb
	sqldb, err := sql2.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("conn1 failed: %v", err)
		return
	}
	fmt.Println(sqldb.Ping())
	fmt.Println(sqldb.Stats())
	//通过已有的数据库来初始化一个*gorm.DB
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqldb,
	}), &gorm.Config{})
	fmt.Println(gormDB.Name())
	fmt.Println(gormDB.Statement)
}
