package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// var dsn = "tanu:tanu123@/go_test_models?charset=utf8mb4"
	// var d, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	d, err := gorm.Open("mysql", "root:Tanu@1906@/bookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	db = d
}

func GetDB() *gorm.DB {
	return db
}
