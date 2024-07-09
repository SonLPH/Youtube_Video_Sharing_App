package models

import "time"

type Notification struct {
	ID        uint      `json:"id" gorm:"primary_key,autoIncrement"`
	UserID    uint      `json:"user_id"`
	User      User      `gorm:"foreignKey:UserID;references:ID"`
	VideoID   uint      `json:"video_id"`
	Video     Video     `gorm:"foreignKey:VideoID;references:ID"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
