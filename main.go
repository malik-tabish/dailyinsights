package main

import (
	"log"
	"project/config"
	connection "project/database"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	config.Init()
	connection.InitDB()
	// defer connection.CloseDB()
	//models.AutoMigrate()
}

func main() {

	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.GET("/", handler)

	log.Println("Server is running at http://localhost:" + config.Environment.AppPort)
	log.Fatal(server.Run(":" + config.Environment.AppPort))
}

func handler(c *gin.Context) {
	c.String(200, "Welcome to the server!")
}
