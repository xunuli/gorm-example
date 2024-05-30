package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type test1 struct {
	ID   int `gorm:"PRIMARYKEY;AUTO_INCRMENT"`
	Name string
	Hoby string
}

func main() {
	//dsn
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}
	//自动迁移，建表
	err = db.AutoMigrate(&test1{})
	if err != nil {
		log.Fatalf("迁移失败：%v", err)
	}
	//根据map创建单条记录
	result := db.Model(&test1{}).Create(map[string]interface{}{
		"ID": 2, "Name": "didi_2", "Hoby_2": "篮球_2",
	})
	//insert into `test1` (`ID`, `Name`, `Hoby`) value (1, `didi_1`, `篮球`)
	fmt.Println(result)

	//根据map创建多条记录
	result1 := db.Model(&test1{}).Create([]map[string]interface{}{
		{"ID": 3, "Name": "didi_3", "Hoby": "篮球_3"},
		{"ID": 4, "Name": "didi_4", "Hoby": "篮球_4"},
		{"ID": 5, "Name": "didi_5", "Hoby": "篮球_5"},
	})
	fmt.Println(result1)

}
