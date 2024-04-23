package models

import (
	connection "project/database"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null"`
	Email    string `gorm:"unique;not null"`
	FullName string
	Password string `gorm:"not null"`
}

func AutoMigrate() {
	if err := connection.DB.AutoMigrate(&User{}); err != nil {
		panic(err)
	}
}
