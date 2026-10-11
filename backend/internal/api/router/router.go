package router

import (
	"pushy/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handlers.Handler) *gin.Engine{
	router := gin.New()


	return router
}