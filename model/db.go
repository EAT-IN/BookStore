package model

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func InitDB() {
	DB, err = gorm.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	// 关闭表明复数
	DB.SingularTable(true)
	// 开启数据库调试
	DB.LogMode(true)
	//设置最大闲置数量
	DB.DB().SetMaxIdleConns(5)
	// 设置最大连接数
	DB.DB().SetMaxOpenConns(10)
	// 自动建表
	DB.AutoMigrate(&User{}, &Book{}) // 自动创建表

}
