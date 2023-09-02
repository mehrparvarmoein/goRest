package models

import (
	"fmt"
	"log"
	"math/rand"
	"rest_api/config"
	"rest_api/packages"
	"time"

	"github.com/pooriaghaedi/authority"
	"gorm.io/gorm"
)

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789" +
	"!@#$%^&*()-_+=<>,."

func generateRandomPassword(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[r.Intn(len(charset))]
	}
	return string(b)
}

func CreateSuperAdmin() {
	superAdminUsername := "superadmin"
	rawPassword := generateRandomPassword(12)

	// Check if superadmin already exists
	var user User
	if err := config.Database.Where("username = ?", superAdminUsername).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {

			superAdmin := User{
				Username: superAdminUsername,
				Password: rawPassword,
			}

			_, err := superAdmin.Save()
			if err != nil {
				log.Fatalf("Failed to create superadmin: %v", err)
			}

			log.Println("Superadmin user created, Password: ", rawPassword)
			CreatePermission(superAdmin)

		} else {
			// Some other error occurred
			log.Fatalf("Failed to check superadmin existence: %v", err)
		}
	} else {
		log.Println("Superadmin user already exists!")

	}
}

func CreatePermission(user User) {
	fmt.Println(user.ID)
	err := packages.Rbac.CreatePermission(authority.Permission{
		Name: "Superadmin",
		Slug: "superadmin",
	})

	if err != nil {
		log.Fatalf("Failed to create Permission: %v", err)

	}

	err = packages.Rbac.CreateRole(authority.Role{
		Name: "sa",
		Slug: "sa",
	})

	if err != nil {
		log.Fatalf("Failed to create role: %v", err)
	}

	err = packages.Rbac.AssignPermissionsToRole("sa", []string{
		"superadmin",
	})

	if err != nil {
		log.Fatalf("Failed to assign permission to role: %v", err)
	}

	err = packages.Rbac.AssignRoleToUser(user.ID, "sa")
	if err != nil {
		log.Fatalf("Failed to assign role to user: %v", err)
	}
}
