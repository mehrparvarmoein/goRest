package controllers

import (
	"net/http"
	"rest_api/packages"

	"github.com/gin-gonic/gin"
	"github.com/harranali/authority"
)

func IndexRoles(context *gin.Context) {

	roles, err := packages.Rbac.GetAllRoles()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": roles})

}

func StoreRoles(context *gin.Context) {

	err := packages.Rbac.CreateRole(authority.Role{
		Name: "Role 1",
		Slug: "role-1",
	})

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"data": "role created"})
}
