package config

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/campus_course?charset=utf8mb4&parseTime=True&loc=Local"

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("MySQL连接失败：", err)
	}

	DB = database

	fmt.Println("MySQL数据库连接成功")
}
