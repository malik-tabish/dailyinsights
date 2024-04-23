package connection

import (
	"fmt"
	"log"

	"project/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Environment.DBUsername, config.Environment.DBPassword, config.Environment.DBHost, config.Environment.DBPort, config.Environment.DBName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error while connecting to database:	%v", err)
	}
	log.Println("Database Connected Successfully")
}

func CloseDB() {
	mysqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Error getting DB: %v", err)
	}
	if err := mysqlDB.Close(); err != nil {
		log.Fatalf("Error closing database: %v", err)
	}
}
