package api

import (
	"net/http"

	"github.com/TutorialEdge/ctxlog"
	"github.com/TutorialEdge/notification-service/internal/subscriber"
	"github.com/gin-gonic/gin"
)

// Subscribe - handles a new subscriber request
func (h *Handler) Subscribe(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.Info(ctx, "request to create new subscriber")
	var newSubscriber subscriber.Subscriber
	if err := c.ShouldBindJSON(&newSubscriber); err != nil {
		h.log.Error(c.Request.Context(), err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "failed to bind request to subscriber object",
			"error":   err.Error(),
			"status":  http.StatusBadRequest,
		})
		return
	}

	ctx = ctxlog.WithFields(ctx, ctxlog.Fields{
		"newSubscriber": newSubscriber,
	})

	createdSubscriber, err := h.subService.CreateSubscriber(ctx, newSubscriber)
	if err != nil {
		h.log.Error(ctx, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create new subscriber",
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
		})
		return
	}

	h.log.Info(ctx, "new subscriber created")
	c.JSON(http.StatusOK, gin.H{
		"message":    "subscribed",
		"subscriber": createdSubscriber,
	})
}

func (h *Handler) Unsubscribe(c *gin.Context) {
	ctx := c.Request.Context()
	h.log.Info(ctx, "request to unsubscriber")
	subID := c.Param("subscriberid")

	if err := h.subService.Unsubscribe(ctx, subID); err != nil {
		h.log.Error(ctx, err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to unsubscribe",
			"error":   err.Error(),
			"status":  http.StatusInternalServerError,
		})
	}

	h.log.Info(ctx, "new subscriber created")
	c.JSON(http.StatusOK, gin.H{"message": "unsubscribed"})
}

func (h *Handler) GetSubscriber(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "one subscriber"})
}

func (h *Handler) GetSubscribers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "all subscribers"})
}
