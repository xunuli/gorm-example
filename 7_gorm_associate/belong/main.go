package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

// 初始化数据库连接
var db *gorm.DB

func init() {
	var err error
	//连接数据库
	dsn := "root:123456@tcp(192.168.198.130:3306)/gormtestdb?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database!")
	}
}

type User struct {
	Name string `gorm:"comment:用户名"`
	//默认情况，CompanyID被隐含地用来在User和company之间创建一个外键联系
	CompanyID int     `gorm:"comment:公司ID"`
	Company   Company `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"` //添加外键约束
	gorm.Model
}

//	type User struct {
//		Name         string  `gorm:"comment:用户名"`
//		CompanyRefer int     `gorm:"comment:公司ID"`
//		Company      Company `gorm:"foreignKey:CompanyRefer;"`
//		gorm.Model
//	}
type Company struct {
	ID   int    `gorm:"primarykey;auto_incrment;comment:公司ID"`
	Name string `gorm:"comment:公司名字"`
}

func main() {
	//自动迁移，用于创建表，由于user表里面有company表结构，所以只需要迁移一个user表就行
	err := db.AutoMigrate(&User{})
	if err != nil {
		log.Fatalf("用户表迁移失败: %v", err)
		return
	}

	//创建记录
	com := Company{
		ID:   2,
		Name: "tencent",
	}
	user := User{
		Name:    "xuji",
		Company: com,
	}
	db.Create(&user)

	//查询记录
	var uu []User
	db.Model(&User{}).First(&uu)
	fmt.Println(uu)
	db.Model(&User{}).Preload("Company").First(&uu)
	fmt.Println(uu)

	//查找关联
	db.Where("id = ?", 1).Take(&uu)
	fmt.Println(user)
	var cc Company
	//// `user` 是源模型，它的主键不能为空
	//// 关系的字段名是 `Company`
	//// 如果匹配了上面两个要求，会开始关联模式，否则会返回错误
	db.Model(&user).Association("Company").Find(&cc)
	fmt.Println(cc)

	//删除关联
	var uu1 User
	db.Where("id = ?", 1).First(&uu1)
	fmt.Println(uu1)
	db.Model(&User{}).Association("Company").Delete(&Company{ID: 1})

	//修改关联
	var uu2 User
	db.Where("id = ?", 1).First(&uu2)
	fmt.Println(uu1)
	db.Model(&User{}).Association("Company").Replace(&Company{ID: 2})

}

func (u *User) BeforeSave(tx *gorm.DB) error {
	u.Name = u.Name + "_" + strconv.Itoa(u.Company.ID)
	return nil
}
