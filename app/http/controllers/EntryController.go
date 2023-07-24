package controllers

import (
	"net/http"
	"rest_api/app/models"
	"rest_api/helper"

	"github.com/gin-gonic/gin"
)

func AddEntry(context *gin.Context) {
	var input models.Entry
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.UserID = user.ID

	savedEntry, err := input.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"data": savedEntry})
}

func GetAllEntries(context *gin.Context) {
	user, err := helper.CurrentUser(context)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"data": user.Entries})

}

// packages.Rbac.CreatePermission(authority.Permission{
//     Name: "Permission a",
//     Slug: "permission-a",
// })

// packages.Rbac.AssignPermissionsToRole("role-a", []string{
//     "permission-a",
// })

// packages.Rbac.CreateRole(authority.Role{
//     Name: "Role 2",
//     Slug: "role-2",
// })
// packages.Rbac.AssignRoleToUser(user.ID, "role-1")

// ok, err := packages.Rbac.CheckUserPermission(user.ID, "a")
// if err != nil {
//     context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//     return
// }

// if ok {
//     context.JSON(http.StatusOK, gin.H{"data": user.Entries, "userId": user.ID})
// }else{
//     context.JSON(http.StatusForbidden, gin.H{"message": "access denied"})
// }
