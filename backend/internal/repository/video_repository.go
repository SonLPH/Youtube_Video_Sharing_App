package repository

import (
	"youtube_video_sharing_app/internal/models"

	"gorm.io/gorm"
)

func CreateVideo(db *gorm.DB, video *models.Video) error {
	return db.Create(video).Error
}
