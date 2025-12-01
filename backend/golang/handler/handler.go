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

	key := h.Service.Redis.GetRandomKey()
	val, err := h.Service.Redis.GetBytes(key)
	if err != nil {
		panic(err)
	}
	c.Data(http.StatusOK, "image/jpeg", val)
}
