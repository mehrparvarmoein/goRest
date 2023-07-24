package middleware

import (
	"fmt"
	"net/http"
	"rest_api/helper"
	"rest_api/packages"
	"strings"

	"github.com/gin-gonic/gin"
)

func Permission(p string) gin.HandlerFunc {
    return func(context *gin.Context) {
        
        user, err := helper.CurrentUser(context)
        if err != nil {
            context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            context.Abort()
            return
        }

        permissions := strings.Split(p,"|")
        for _,p := range permissions {
            fmt.Println(p)
            ok, _ := packages.Rbac.CheckUserPermission(user.ID, p)
            if ok {
                context.Next() 
                return
            }

        }

        context.JSON(http.StatusForbidden, gin.H{"error": "not Authorized!"})
        context.Abort()
    }
}