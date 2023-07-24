package controllers

import (
	"net/http"
	"os"
	"rest_api/app/http/validation"
	"rest_api/app/models"
	"rest_api/config"
	"rest_api/helper"

	"github.com/gin-gonic/gin"
)

func Register(context *gin.Context) {
	var input validation.AuthenticationInput

	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

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

	context.JSON(http.StatusOK, gin.H{"user" : user , "jwt": jwt ,  "ttl" : os.Getenv("TOKEN_TTL")})
}

