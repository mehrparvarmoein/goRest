package controllers

import (
	"net/http"
	"rest_api/packages"

	"github.com/gin-gonic/gin"
	"github.com/harranali/authority"
)

type PermissionInput struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func IndexPermissions(context *gin.Context) {

	permissions, err := packages.Rbac.GetAllPermissions()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": permissions})

}

func StorePermissions(context *gin.Context) {
	var input RoleInput

	// Bind the JSON body to the struct
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	name := input.Name
	if name == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, name is required"})
		return
	}

	err := packages.Rbac.CreateRole(authority.Role{
		Name: name,
		Slug: toSlug(name),
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permissions, err := packages.Rbac.GetAllPermissions()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": permissions})
}

// func ShowPermissions(context *gin.Context) {
// 	id := context.Param("id")

// 	permissions, err := packages.Rbac.GetAllPermissions()

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	context.JSON(http.StatusOK, gin.H{"data": permissions})
// }

func DeletePermissions(context *gin.Context) {
	var input RoleInput

	// Bind the JSON body to the struct
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	slug := input.Slug
	if slug == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input, slug is required"})
		return
	}

	err := packages.Rbac.DeleteRole(slug)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	permissions, err := packages.Rbac.GetAllPermissions()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": permissions})
}
