package handler

import (
	"log/slog"
	"net/http"

	"my-go-api/factory"
	"my-go-api/utility"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	logger       *slog.Logger
	imageFactory *factory.ImageFactory
}

func NewImageHandler(logger *slog.Logger, imageFactory *factory.ImageFactory) *ImageHandler {
	return &ImageHandler{
		logger:       logger.With("handler"),
		imageFactory: imageFactory,
	}
}

func (h *ImageHandler) GetRandomImage(c *gin.Context) {
	imageName := h.imageFactory.GetRandomImageName()
	imageContent := h.imageFactory.GetImage(imageName)
	c.Data(http.StatusOK, "image/jpeg", imageContent)
}

func (h *ImageHandler) RefreshCache() {
	h.imageFactory.RefreshCache()
}

func (h *ImageHandler) RateLimitMiddleware(ipRateLimiter *utility.IPRateLimiter, next gin.HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip := c.ClientIP()
		limiter := ipRateLimiter.GetLimiter(ip)
		if limiter.Allow() {
			next(c)
		} else {
			h.logger.Error("Rate Limit Exceeded", http.StatusTooManyRequests)
		}
	}
}
