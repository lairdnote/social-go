package model

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// 数据库链接单例
var DB *gorm.DB

// 初始化mysql连接
func Database(connString string) {
	db, err := gorm.Open("postgres", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err.Error())
	}
	db.DB().SetMaxIdleConns(50)
	//打开
	db.DB().SetMaxOpenConns(100)
	//超时
	db.DB().SetConnMaxLifetime(time.Second * 30)

	DB = db


	migration()
}

func migration() {
	// 自动迁移模式
	err := DB.AutoMigrate(&User{})
	if err != nil {
		fmt.Print(err)
	}
}
