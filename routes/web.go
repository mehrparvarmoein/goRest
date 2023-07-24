package routes

import (
	"rest_api/app/http/controllers"
	"rest_api/app/http/middleware"
	
	"github.com/gin-gonic/gin"
)

func Web() {
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
}
