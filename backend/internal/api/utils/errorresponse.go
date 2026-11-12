package utils

import (
	"pushy/internal/dto"

	"github.com/gin-gonic/gin"
)

func WriteProblem(c *gin.Context, status int, title string, detail string) {
	c.Header("Content-Type", "application/problem+json")

	c.JSON(status, dto.ProblemResponse{
		Type:   "about:blank",
		Title:  title,
		Status: status,
		Detail: detail,
	})
}
