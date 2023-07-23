package controllers

import (
	"net/http"
	"rest_api/config"
	"rest_api/helper"
	"rest_api/models"
	"rest_api/validation"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input validation.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// var count int64
	// config.Database.Model(&models.User{}).Where("username = ?", input.Username).Count(&count)
	// if count > 0 {
	//     context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "username already exists"})
	//     return
	// }

	var userFind models.User
	exists := config.Database.Where("username = ?", input.Username).First(&userFind)

	if exists.RowsAffected > 0 {
		context.JSON(http.StatusUnprocessableEntity, gin.H{"error": "username already exists"})
		return
	}

	user := models.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := user.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})
}

func Login(context *gin.Context) {
	var input validation.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := models.FindUserByUsername(input.Username)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = user.ValidatePassword(input.Password)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	jwt, err := helper.GenerateJWT(user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"jwt": jwt})
}

// func Welcome(context *gin.Context) {

//     context.JSON(http.StatusOK, gin.H{"jwt": "jwt"})
// }
