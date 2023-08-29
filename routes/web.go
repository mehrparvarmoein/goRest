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

	adminRoutes := router.Group("/admin")
	adminRoutes.Use(middleware.JWTAuthMiddleware())
	adminRoutes.Use(middleware.Permission("superadmin"))
	adminRoutes.GET("/users", controllers.GetAllUsers)
	adminRoutes.GET("/roles", controllers.IndexRoles)
	adminRoutes.POST("/roles", controllers.StoreRoles)
	// adminRoutes.GET("/roles/:id", controllers.ShowRoles)
	// adminRoutes.PUT("/roles/:id", controllers.UpdateRoles)
	adminRoutes.DELETE("/roles", controllers.DeleteRoles)
	adminRoutes.GET("/permissions", controllers.IndexPermissions)
	adminRoutes.POST("/permissions", controllers.StorePermissions)
	adminRoutes.DELETE("/permissions", controllers.DeletePermissions)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.Use(middleware.Permission("superadmin"))
	protectedRoutes.POST("/posts", controllers.AddPost)
	protectedRoutes.GET("/posts", controllers.GetAllPosts)
	protectedRoutes.GET("/user-posts", controllers.GetUserPosts)

	router.Run(":8000")
}
