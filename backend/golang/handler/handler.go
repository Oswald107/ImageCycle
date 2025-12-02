package handler

import (
	"net/http"

	"my-go-api/service"

	"github.com/gin-gonic/gin"
)

type ImageHandler struct {
	Service service.ImageService
}

func NewImageHandler(s service.ImageService) *ImageHandler {
	return &ImageHandler{Service: s}
}

func (h *ImageHandler) GetRandomImage(c *gin.Context) {
	val := h.Service.GetRandomImage()
	c.Data(http.StatusOK, "image/jpeg", val)
}

func (h *ImageHandler) RefreshCache() {
	h.Service.RefreshCache()
}
