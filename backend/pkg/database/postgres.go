package database

import (
	"log"
	"time"
	"youtube_video_sharing_app/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(dsn string) *gorm.DB {
	var database *gorm.DB

	var err error

	for i := 1; i <= 3; i++ {
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if err == nil {
			break
		} else {
			log.Printf("Attempt %d: Failed to initialize dabase. Retrying .... ", i)
			time.Sleep(3 * time.Second)
		}
	}

	err = database.AutoMigrate(&models.User{}, &models.Video{}, &models.Notification{})
	if err != nil {
		log.Fatal(err)
	}

	return database
}
