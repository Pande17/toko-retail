package repository

import (
	"fmt"
	"log"
	"os"
	"projek/toko-retail/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlDB struct{
		DB *gorm.DB
}


var Mysql MysqlDB

func OpenDB() (*gorm.DB, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

		fmt.Println("Connection String:", connString)

	mysqlConn, err := gorm.Open(mysql.Open(connString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	Mysql = MysqlDB{
		DB: mysqlConn,
	}

	err = autoMigrate(mysqlConn)
	if err != nil {
		log.Fatal(err)
	}

	return mysqlConn, err
}

func autoMigrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&model.Model{},
		&model.Barang{},
		&model.Penjualan{},
		&model.Diskon{},
		&model.Histori{},
	)

	if err != nil {
		return err
	}

	return nil
}