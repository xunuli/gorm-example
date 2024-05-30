package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 定义模型
type User struct {
	ID       int `gorm:"AUTO_INCREMENT;primaryKey"`
	Name     string
	Age      int64
	Birthday time.Time
}

func main() {
	//连接数据库
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}
	//自动迁移，建表
	//err = db.AutoMigrate(&User{})
	//if err != nil {
	//	log.Fatalf("迁移失败：%v", err)
	//}
	fmt.Println("**********查询语句****************")
	var user User
	//获取第一条记录，按主键升序
	result := db.First(&user, 1)
	//select * from users order by id limit 1;
	fmt.Println(result.RowsAffected) // 返回找到的记录数
	fmt.Println(result.Error)        // returns error or nil
	fmt.Println(user)
	//获取一条记录，没有指定排序字段
	db.Take(&user)
	//select * from users limit 1;
	fmt.Println(user)
	//获取最后一条记录，主键降序
	db.Last(&user)
	//select * from users order by id desc limit 1;
	fmt.Println(user)

	result11 := map[string]interface{}{}
	db.Model(&User{}).First(&result11)
	fmt.Println(result11)

	result12 := map[string]interface{}{}
	//db.Table("users").First(&result12) 不满足
	db.Table("users").Take(&result12)
	fmt.Println(result12)
}
