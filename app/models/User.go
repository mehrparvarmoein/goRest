package models

import (
	"html"
	"log"
	"math/rand"
	"rest_api/config"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;not null;" json:"-"`
	Posts    []Post
}

func (user *User) Save() (*User, error) {
	err := config.Database.Create(&user).Error
	if err != nil {
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(passwordHash)
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func (user *User) ValidatePassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
}

func FindUserByUsername(username string) (User, error) {
	var user User
	err := config.Database.Where("username=?", username).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func FindUserById(id uint) (User, error) {
	var user User
	err := config.Database.Preload("Posts").Where("ID=?", id).Find(&user).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
	"0123456789" +
	"!@#$%^&*()-_+=<>,."

func generateRandomPassword(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
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

		} else {
			// Some other error occurred
			log.Fatalf("Failed to check superadmin existence: %v", err)
		}
	} else {
		log.Println("Superadmin user already exists!")
	}
}
