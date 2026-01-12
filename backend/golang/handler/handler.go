package handler

import (
	"log/slog"
	"net/http"

	"my-go-api/service"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	logger  *slog.Logger
	Service service.ImageService
}

func NewImageHandler(logger *slog.Logger, s service.ImageService) *ImageHandler {
	return &ImageHandler{
		logger:  logger.With("handler"),
		Service: s,
	}
}

// func CORSMiddleware() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
// 		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
// 		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

// 		if c.Request.Method == "OPTIONS" {
// 			c.AbortWithStatus(204)
// 			return
// 		}

// 		c.Next()
// 	}
// }

func (h *ImageHandler) GetRandomImage(c *gin.Context) {
	val := h.Service.GetRandomImage()
	c.Data(http.StatusOK, "image/jpeg", val)
}

func (h *ImageHandler) RefreshCache() {
	h.Service.RefreshCache()
}
