package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

// 定义模型
type UserInfo struct {
	ID       int    `gorm:"AUTO_INCREMENT;primaryKey"`
	Name     string `gorm:"default:hahaha"`
	Age      int64  `gorm:"default:18"`
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
	err = db.AutoMigrate(&UserInfo{})
	if err != nil {
		log.Fatalf("迁移失败：%v", err)
	}

	user := UserInfo{
		Name:     "didi",
		Age:      18,
		Birthday: time.Now(),
	}
	result := db.Create(&user)
	//INSERT INTO `user_infos` (`name`,`age`, `brithday`) VALUES ("didi_xuji", 18)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected) //返回影响行数
	user1 := UserInfo{
		Name:     "didi",
		Age:      18,
		Birthday: time.Now(),
	}
	result1 := db.Session(&gorm.Session{SkipHooks: true}).Create(&user1)
	//INSERT INTO `user_infos` (`name`,`age`, `brithday`) VALUES ("didi", 18)
	fmt.Println(result1.Error)
	fmt.Println(result1.RowsAffected) //返回影响行数
}

func (u *UserInfo) BeforeSave(tx *gorm.DB) error {
	u.Name += "_xuji"
	return nil
}
