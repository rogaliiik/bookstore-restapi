package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB instance
var (
	db *gorm.DB
)

// Create connection with Database
func Connect() {
	db_password := os.Getenv("DB_PASSWORD")

	d, err := gorm.Open("mysql", "root:"+db_password+"@/gobookstore?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// Call DB instance
func GetDB() *gorm.DB {
	return db
}
