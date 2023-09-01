package controllers

import (
	"net/http"
	"regexp"
	"rest_api/packages"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pooriaghaedi/authority"
)

func toSlug(input string) string {
	// Convert to lowercase
	slug := strings.ToLower(input)

	// Replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// Remove any unwanted characters
	reg, err := regexp.Compile("[^a-z0-9-]+")
	if err != nil {
		return slug // If there's an error with the regex, return the slug as is
	}
	slug = reg.ReplaceAllString(slug, "")

	return slug
}

func IndexRoles(context *gin.Context) {

	roles, err := packages.Rbac.GetAllRoles()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": roles})

}

func StoreRoles(context *gin.Context) {
	var RoleInput struct {
		Name string `json:"name" binding:"required"`
	}

	// Bind the JSON body to the struct
	if err := context.ShouldBindJSON(&RoleInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := packages.Rbac.CreateRole(authority.Role{
		Name: RoleInput.Name,
		Slug: toSlug(RoleInput.Name),
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "role created successfully"})
}

func DeleteRoles(context *gin.Context) {
	var RoleInput struct {
		Slug string `json:"slug" binding:"required"`
	}

	// Bind the JSON body to the struct
	if err := context.ShouldBindJSON(&RoleInput); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := packages.Rbac.DeleteRole(RoleInput.Slug)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": "role deleted successfully"})
}

func AssignRoleToUser(context *gin.Context) {
	var requestBody struct {
		Role   string `json:"role" binding:"required"`
		UserID uint   `json:"userid" binding:"required"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := packages.Rbac.AssignRoleToUser(requestBody.UserID, requestBody.Role)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": "role assigned successfully"})
}

func RevokeUserRole(context *gin.Context) {
	var requestBody struct {
		Role   string `json:"role" binding:"required"`
		UserID string `json:"userid" binding:"required"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := packages.Rbac.RevokeUserRole(requestBody.UserID, requestBody.Role)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": "role revoked successfully"})
}

func RevokeRolePermission(context *gin.Context) {
	var requestBody struct {
		Role       string `json:"role" binding:"required"`
		Permission string `json:"permission" binding:"required"`
	}

	if err := context.ShouldBindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := packages.Rbac.RevokeRolePermission(requestBody.Role, requestBody.Permission)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": "permission revoked successfully"})
}
