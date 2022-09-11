package common

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	User "m0nk3y/pocScanner/model"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	host := "localhost"
	port := "3306"
	database := "pocWebSite"
	username := "root"
	password := "qsjdmwdc1b"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset)

	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect database, err: " + err.Error())
	}
	err1 := db.AutoMigrate(&User.User{}) // 自动生成表
	if err1 != nil {
		panic("failed to connect database, err: " + err1.Error())
	}
	DB = db
	return db
}

func GetDB() *gorm.DB {
	return DB
}
