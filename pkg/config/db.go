package config

import (
	"fmt"
	"github.com/rogaliiik/bookstore/pkg/utils"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB instance
var (
	db *gorm.DB
)

// Create connection with Database
func Connect() {
	dbPassword := utils.GetEnvVar("POSTGRES_PASSWORD")
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		Conf.DB.Host,
		Conf.DB.Port,
		Conf.DB.User,
		dbPassword,
		Conf.DB.DBName,
		Conf.DB.Sslmode,
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
