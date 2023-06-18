package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

var Conf Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Print(err)
	}
	configFile, err := os.ReadFile(wd + "/configs/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(configFile, &Conf)
	if err != nil {
		log.Fatal(err)
	}
}

type Config struct {
	Port string   `yaml:"port"`
	DB   DBConfig `yaml:"db"`
}

type DBConfig struct {
	User    string `yaml:"username"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	DBName  string `yaml:"dbname"`
	Sslmode string `yaml:"sslmode"`
}
