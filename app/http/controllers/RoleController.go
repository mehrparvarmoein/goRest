package controllers

import (
	"net/http"
	"regexp"
	"rest_api/packages"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/harranali/authority"
)

type RoleInput struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

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

	roles, err := packages.Rbac.GetAllRoles()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": roles})
}

// func ShowRoles(context *gin.Context) {
// 	id := context.Param("id")

// 	roles, err := packages.Rbac.GetAllRoles()

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	context.JSON(http.StatusOK, gin.H{"data": roles})
// }

func DeleteRoles(context *gin.Context) {
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

	roles, err := packages.Rbac.GetAllRoles()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": roles})
}
