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
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("迁移失败：%v", err)
	}
	//创建多条记录
	fmt.Println("==========================创建单条记录==========================")
	//初始化数据单条记录
	user := User{
		//ID: 1,
		Name:     "xuji",
		Age:      23,
		Birthday: time.Now(),
	}
	//创建记录
	reult := db.Create(&user)
	//查询记录
	fmt.Println(user.ID)
	fmt.Println(reult.Error)
	fmt.Println(reult.RowsAffected) //返回影响行数

	//创建多条记录
	fmt.Println("========================创建多条记录============================")
	//采用切片的方式实现
	user1 := []*User{
		{Name: "lujie", Age: 18, Birthday: time.Now()},
		{Name: "nihao", Age: 19, Birthday: time.Now()},
	}
	result1 := db.Create(user1)
	fmt.Println(result1.Error)
	fmt.Println(result1.RowsAffected)
	//创建多条记录
	fmt.Println("====================根据指定字段创建记录================================")
	//使用指定的字段创建记录
	user2 := User{
		//ID: 1,
		Name:     "didi",
		Age:      26,
		Birthday: time.Now(),
	}
	result2 := db.Select("Name", "Age").Create(&user2)
	//INSERT INTO `users` (`name`,`age`) VALUES ("didi", 26)
	fmt.Println(result2.Error)
	fmt.Println(result2.RowsAffected)
	//omit方法
	user3 := User{
		//ID: 1,
		Name:     "tecent",
		Age:      25,
		Birthday: time.Now(),
	}
	result3 := db.Omit("Name").Create(&user3)
	//INSERT INTO `users` (`age`,`brithtime`) VALUES ("25", time.now())
	fmt.Println(result3.Error)
	fmt.Println(result3.RowsAffected)
	//使用Batchs方法来指定批量插入的批次大小
	fmt.Println("====================CreateInBatches================================")
	//使用CreateInBatches批量创建
	user4 := []User{
		{Name: "didi_1", Age: 1, Birthday: time.Now()},
		{Name: "didi_2", Age: 2, Birthday: time.Now()},
		{Name: "didi_3", Age: 3, Birthday: time.Now()},
	}
	result4 := db.CreateInBatches(&user4, 3)
	fmt.Println(result4.Error)
	fmt.Println(result4.RowsAffected)
}
