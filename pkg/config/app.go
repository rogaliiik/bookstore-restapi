package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// loads values from .env into the system
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

// DB instance
var (
	db *gorm.DB
)

// Create connection with Database
func Connect() {

	db_password, exists := os.LookupEnv("DB_PASSWORD")

	if !exists {
		panic("db_password is not in .env")
	}

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
