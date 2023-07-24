package controllers

import (
	"net/http"
	"rest_api/app/http/validation"
	"rest_api/app/models"
	"rest_api/config"
	"rest_api/helper"

	"github.com/gin-gonic/gin"
)

func GetAllPosts(context *gin.Context) {
	var posts []models.Post
    err := config.Database.Find(&posts).Error

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": posts})
}

func GetUserPosts(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Posts})
}

func AddPost(context *gin.Context) {
	var input validation.PostInput
	
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	post := models.Post{
		Title  : input.Title,
		Content: input.Content,
		UserID : user.ID,
	}

	savedPost, err := post.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedPost})
}



