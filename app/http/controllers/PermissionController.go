package controllers

import (
	"net/http"
	"rest_api/packages"

	"github.com/gin-gonic/gin"
	"github.com/pooriaghaedi/authority"
)

func IndexPermissions(context *gin.Context) {

	permissions, err := packages.Rbac.GetAllPermissions()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": permissions})

}

func StorePermissions(context *gin.Context) {
	var PermissionInput struct {
		Name string `json:"name" binding:"required"`
	}

	// Bind the JSON body to the struct
	if err := context.ShouldBindJSON(&PermissionInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	name := PermissionInput.Name

	err := packages.Rbac.CreatePermission(authority.Permission{
		Name: name,
		Slug: toSlug(name),
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "Permission stored sucessfully"})
}

func DeletePermissions(context *gin.Context) {
	var PermissionInput struct {
		Slug string `json:"slug" binding:"required"`
	}

	// Bind the JSON body to the struct
	if err := context.ShouldBindJSON(&PermissionInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := packages.Rbac.DeletePermission(PermissionInput.Slug)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "Permission deleted sucessfully"})
}

func AssignPermissionsToRole(context *gin.Context) {
	var requestBody struct {
		Role        string   `json:"role" binding:"required"`
		Permissions []string `json:"permissions" binding:"required"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := packages.Rbac.AssignPermissionsToRole(requestBody.Role, requestBody.Permissions)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "Permissions assigned successfully"})
}
