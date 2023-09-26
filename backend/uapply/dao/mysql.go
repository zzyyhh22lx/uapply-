package dao

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDb() *gorm.DB {
	return db
}

var (
	db  *gorm.DB
	err error
)

func InitMysql() {

	// 连接数据库
	dsn := "root:123456@(localhost:3306)/uapply?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql connect error")
	}
}
