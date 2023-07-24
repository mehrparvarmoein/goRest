package models

import (
	"rest_api/config"

	"gorm.io/gorm"
)

type Post struct {
    gorm.Model
    Title   string `json:"title"`
    Content string `gorm:"type:text" json:"content"`
    UserID  uint
}

func (post *Post) Save() (*Post, error) {
    err := config.Database.Create(&post).Error
    if err != nil {
        return &Post{}, err
    }
    return post, nil
}