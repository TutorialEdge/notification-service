package api

import (
	"net/http"

	"github.com/TutorialEdge/notification-service/internal/subscriber"
	"github.com/gin-gonic/gin"
)

// Subscribe - handles a new subscriber request
func (h *Handler) Subscribe(c *gin.Context) {
	var newSubscriber subscriber.Subscriber
	if err := c.ShouldBindJSON(&newSubscriber); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request to subscriber object",
			"error":   err.Error(),
			"status":  http.StatusBadRequest,
		})
	}

	createdSubscriber, err := h.subService.CreateSubscriber(
		c.Request.Context(),
		newSubscriber,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create new subscriber",
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "subscribed",
		"subscriber": createdSubscriber,
	})
}

func (h *Handler) Unsubscribe(c *gin.Context) {
	subID := c.Param("subscriberid")

	if err := h.subService.Unsubscribe(c.Request.Context(), subID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to unsubscribe",
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "unsubscribed"})
}

func (h *Handler) GetSubscriber(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "one subscriber"})
}

func (h *Handler) GetSubscribers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "all subscribers"})
}
