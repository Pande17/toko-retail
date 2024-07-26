package repository

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	const MYSQL = "root:@tcp(127.0.0.1:3306)/db_tokoretail_2?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := MYSQL

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("can not connect to database nigga")
	}

	fmt.Println("Connected to database nigger...")
}