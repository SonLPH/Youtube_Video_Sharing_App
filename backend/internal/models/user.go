package models

import "time"

type LoginUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type User struct {
	ID         uint      `json:"id" gorm:"primary_key;autoIncrement"`
	Username   string    `json:"username" gorm:"unique"`
	Password   string    `json:"password"`
	Videos     []Video   `json:"videos" gorm:"foreignKey:UserID"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdateTime time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}
