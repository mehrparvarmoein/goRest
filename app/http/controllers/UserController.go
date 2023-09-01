package controllers

import (
	"net/http"
	"rest_api/app/models"
	"rest_api/config"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {

	var users []models.User
	// err := config.Database.Find(&users).Error
	err := config.Database.Debug().Preload("Roles").Find(&users).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": users})

}
