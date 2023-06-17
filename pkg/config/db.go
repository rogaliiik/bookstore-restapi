package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB instance
var (
	db *gorm.DB
)

// Create connection with Database
func Connect() {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		Conf.dbHost,
		Conf.dbPort,
		Conf.dbUser,
		Conf.dbPassword,
		Conf.dbName,
	)

	d, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	db = d
}

// Call DB instance
func GetDB() *gorm.DB {
	return db
}
