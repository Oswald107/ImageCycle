package main

import (
	// "encoding/json"
	"log/slog"
	"my-go-api/handler"
	"my-go-api/service"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
)

func main() {
	f, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger := slog.New(slog.NewTextHandler(f, nil))
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Server running on :8080")

	redisService := service.NewRedisService(logger, "localhost:6379", "", 0)
	imageService := service.NewImageService(logger, *redisService)

	handler := handler.NewImageHandler(logger, *imageService)
	handler.RefreshCache()

	c := cron.New()
	c.AddFunc("*/30 * * * * *", handler.RefreshCache)
	c.Start()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	router.GET("/random", handler.GetRandomImage)

	router.Run("localhost:8080")
}
