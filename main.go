package main

import (
	"fmt"
	"log"
	"rest_api/app/models"
	"rest_api/config"
	"rest_api/packages"
	"rest_api/routes"

	"github.com/joho/godotenv"
)


func main() {
    loadEnv()
    loadDatabase()
    packages.InitAuthority()
    serveApplication()

}

func loadEnv() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func loadDatabase() {
    config.Connect()
    config.Database.AutoMigrate(&models.User{})
    config.Database.AutoMigrate(&models.Post{})
}

func serveApplication() {
    routes.Web()
    fmt.Println("Server running on port 8000")
}