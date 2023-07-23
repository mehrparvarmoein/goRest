package controllers

import (
    "rest_api/helper"
    "rest_api/models"
    "github.com/gin-gonic/gin"
    "net/http"
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


// packages.Auth.CreatePermission(authority.Permission{
    //     Name: "Permission a",
    //     Slug: "permission-a",
    // })
    
    // packages.Auth.AssignPermissionsToRole("role-a", []string{
    //     "permission-a",
    // })

    // packages.Auth.CreateRole(authority.Role{
    //     Name: "Role 2",
    //     Slug: "role-2",
    // })
    // packages.Auth.AssignRoleToUser(user.ID, "role-1")

    // ok, err := packages.Auth.CheckUserPermission(user.ID, "a")
    // if err != nil {
    //     context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    //     return
    // }

    // if ok {
    //     context.JSON(http.StatusOK, gin.H{"data": user.Entries, "userId": user.ID})
    // }else{
    //     context.JSON(http.StatusForbidden, gin.H{"message": "access denied"})
    // }