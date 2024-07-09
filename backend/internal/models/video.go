package models

import "time"

type Video struct {
	ID          uint      `json:"id" gorm:"primary_key;autoIncrement"`
	URL         string    `json:"url" gorm:"unique"`
	Title       string    `json:"title"`
	UserID      uint      `json:"user_id"`
	User        User      `gorm:"foreignKey:UserID;references:ID"`
	Description string    `json:"description"`
	Like        int64     `json:"like"`
	Dislike     int64     `json:"dislike"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}
