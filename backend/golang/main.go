package main

import (
	// "encoding/json"
	"log/slog"
	"my-go-api/handler"
	"os"

	"my-go-api/factory"
	"my-go-api/utility"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/robfig/cron"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	f, err := os.OpenFile("logs/app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	logger := slog.New(slog.NewTextHandler(f, nil))
	if err != nil {
		logger.Error(err.Error())
	}
	logger.Info("Server running on :8080")

	imageFactory := factory.NewImageFactory(logger)

	handler := handler.NewImageHandler(logger, imageFactory)
	handler.RefreshCache()

	c := cron.New()
	c.AddFunc("* */30 * * * *", handler.RefreshCache)
	c.Start()

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}

	router.Use(cors.New(config))

	ipRateLimiter := utility.NewIPRateLimiter()

	router.GET("/random", handler.RateLimitMiddleware(ipRateLimiter, handler.GetRandomImage))
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	router.Run(":8080")
}
