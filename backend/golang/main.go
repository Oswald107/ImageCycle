package main

import (
	// "encoding/json"
	"log"
	"my-go-api/handler"
	"my-go-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	redisService := service.NewRedisService("localhost:6379", "", 0)
	imageService := service.NewImageService(*redisService)

	handler := handler.NewImageHandler(*imageService)

	router := gin.Default()
	router.GET("/random", handler.GetRandomImage)

	router.Run("localhost:8080")

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
