package main

import (
	"fmt"
	"net/http"
	"youtube_video_sharing_app/config"
	"youtube_video_sharing_app/internal/handlers"
	"youtube_video_sharing_app/internal/middleware"
	"youtube_video_sharing_app/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	config := config.LoadConfig()
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
	)
	db := database.ConnectDatabase(dsn)

	r := gin.Default()
	api := r.Group("/api")
	api.POST("/register", handlers.Register(db))
	api.POST("/login", handlers.Login(db))

	protected := r.Group("/api").Use(middleware.AuthMiddleware())
	protected.GET("/protected", func(c *gin.Context) {
		claims, _ := c.Get("claims")
		c.JSON(http.StatusOK, gin.H{"message": "You are authorized", "claims": claims})
	})

	r.Run(":8080")
}
