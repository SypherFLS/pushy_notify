package router

import (
	"pushy/internal/api/handlers"

	"github.com/gin-gonic/gin"
)

func NewRouter(h *handlers.Handler) *gin.Engine {
	router := gin.New()

	router.POST("/sub")
	router.GET("/subs")
	router.DELETE("/unsub")

	return router
}
