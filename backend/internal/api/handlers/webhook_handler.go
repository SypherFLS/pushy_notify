package handlers

import (
	"net/http"
	"pushy/internal/api/apiconsts"
	"pushy/internal/api/utils"

	"github.com/gin-gonic/gin"
)

func GitHubWebhookHandler(c *gin.Context) {
	eventTypeRaw, exists := c.Get(apiconsts.GithubEvents)

	if !exists {
		utils.WriteProblem(c, 404, "wrong webhook data", "context key missed")
		return
	}

	deliveryIDRaw, exists := c.Get(apiconsts.GithubDelivery)

	if !exists {
		utils.WriteProblem(c, 404, "wrong webhook data", "context key missed")
		return
	}

	eventType := eventTypeRaw.(string)
	deliveryID := deliveryIDRaw.(int)

	

	bodyBytes, err := c.GetRawData()

	if err != nil {
		utils.WriteProblem(c, 404, "webhook delivery failed", "nil body")
	}

    // 1. Проверить подпись.
    // 2. Проверить deliveryID на дубликат.
    // 3. Создать внутреннее событие.
    // 4. Опубликовать его в RabbitMQ.

    c.Status(http.StatusAccepted)

	_ = eventType
	_ = deliveryID
	_ = bodyBytes
}
