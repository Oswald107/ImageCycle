package handler

import (
	"log/slog"
	"net/http"

	"my-go-api/service"

	"my-go-api/utility"

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

func (h *ImageHandler) GetRandomImage(c *gin.Context) {
	val := h.Service.GetRandomImage()
	c.Data(http.StatusOK, "image/jpeg", val)
}

func (h *ImageHandler) RefreshCache() {
	h.Service.RefreshCache()
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
