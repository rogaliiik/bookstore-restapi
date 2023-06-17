package config

import (
	"log"

	"github.com/joho/godotenv"

	"github.com/rogaliiik/bookstore/pkg/utils"
)

var Conf *Config

// Init config
func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	dbPassword := utils.GetEnvVar("POSTGRES_PASSWORD")
	dbName := utils.GetEnvVar("DB_NAME")
	dbUser := utils.GetEnvVar("DB_USER")
	dbHost := utils.GetEnvVar("DB_HOST")
	dbPort := utils.GetEnvVar("DB_PORT")

	dbConf := NewDBConfig(dbPassword, dbName, dbUser, dbHost, dbPort)

	port := utils.GetEnvVar("PORT")

	Conf = NewConfig(dbConf, port)
}

type Config struct {
	*DBConfig
	Port string
}

func NewConfig(dbConf *DBConfig, port string) *Config {
	return &Config{
		dbConf,
		port,
	}
}

type DBConfig struct {
	dbPassword string
	dbName     string
	dbUser     string
	dbHost     string
	dbPort     string
}

func NewDBConfig(password string, name string, user string, host string, port string) *DBConfig {
	return &DBConfig{
		password,
		name,
		user,
		host,
		port,
	}
}
