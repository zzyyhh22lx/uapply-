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

	// lmx的远程数据库
	//dsn := "root:123456@(114.116.89.5:3306)/BMS?charset=utf8mb4&parseTime=True&loc=Local"

	// 其他测试用的数据库
	// 王伟的远程数据库1
	dsn := "root:123456@(localhost:3306)/uapply?charset=utf8mb4&parseTime=True&loc=Local"
	// 王伟的远程数据库2
	// dsn := "bmstest:6j5AkLSjFcMibxA4@(shrewdboy.top:3306)/bmstest?charset=utf8mb4&parseTime=True&loc=Local"

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("mysql connect error")
	}
}
