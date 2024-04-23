package config

import (
	"log"
	"os"
)

type Config struct {
	AppPort    string
	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string
}

var Environment Config

func Init() {
	log.Println("Load from .env ")
	Environment = Config{
		AppPort:    os.Getenv("PORT"),
		DBHost:     os.Getenv("DBHOST"),
		DBPort:     os.Getenv("DBPORT"),
		DBUsername: os.Getenv("DBUSERNAME"),
		DBPassword: os.Getenv("DBPASSWORD"),
		DBName:     os.Getenv("DBNAME"),
	}
}
