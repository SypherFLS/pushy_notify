package handlers

import (
	"net/http"
	"pushy/internal/dto"
	"pushy/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) SubscribeHandler(c *gin.Context) {
	var path dto.URL

	if err := c.ShouldBindJSON(&path); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

}
