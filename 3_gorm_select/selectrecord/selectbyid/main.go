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
	dsn := "root:123456@tcp(127.0.0.1:3306)/study?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("conn failed: %v", err)
		return
	}
	//通用数据接口
	sqlDB, err := db.DB()
	fmt.Println(db.Name())
	fmt.Println(sqlDB.Stats())

	//自动迁移，建表
	err = db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("迁移失败：%v", err)
	}

	//创建数据
	user := []User{
		{Name: "tecent", Age: 18, Birthday: time.Now()},
		{Name: "didi", Age: 19, Birthday: time.Now()},
		{Name: "alibaba", Age: 20, Birthday: time.Now()},
		{Name: "byteda", Age: 21, Birthday: time.Now()},
		{Name: "baidu", Age: 22, Birthday: time.Now()},
		{Name: "meituan", Age: 23, Birthday: time.Now()},
		{Name: "guoqi", Age: 24, Birthday: time.Now()},
		{Name: "yanjiusuo", Age: 25, Birthday: time.Now()},
	}
	result := db.CreateInBatches(&user, 8)
	//查询记录
	fmt.Println(result)
	fmt.Println(result.Error)
	fmt.Println(result.RowsAffected) //返回影响行数

	//查询
	var user1 []User
	//select * from user1  where id in(1, 2, 3, 4, 5)
	db.Find(&user1, []int{1, 2, 3, 4, 5})
	fmt.Println(user1)

	//Select * From users where id = 2;
	db.First(&user1, 2)
	fmt.Println(user1)
	//Select * From users where id = 3;
	db.First(&user1, "3")
	fmt.Println(user1)
}
