package repository

import (
	"youtube_video_sharing_app/internal/models"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	return db.Create(user).Error
}
