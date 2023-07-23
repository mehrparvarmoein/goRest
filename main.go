package main

import (
	"fmt"
	"log"
	"rest_api/config"
	"rest_api/controllers"
	"rest_api/middleware"
	"rest_api/models"
	"rest_api/packages"

	"github.com/gin-gonic/gin"
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
    config.Database.AutoMigrate(&models.Entry{})
}

func serveApplication() {
    router := gin.Default()

    publicRoutes := router.Group("/auth")
    publicRoutes.POST("/register", controllers.Register)
    publicRoutes.POST("/login", controllers.Login)

    protectedRoutes := router.Group("/api")
    protectedRoutes.Use(middleware.JWTAuthMiddleware())
    protectedRoutes.Use(middleware.Permission("a|b"))
    protectedRoutes.POST("/entry", controllers.AddEntry)
    protectedRoutes.GET("/entry", controllers.GetAllEntries)

    router.Run(":8000")
    fmt.Println("Server running on port 8000")
}