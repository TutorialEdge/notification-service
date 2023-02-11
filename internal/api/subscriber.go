package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Subscribe(c *gin.Context) {
	// subscribes to a specific list
	c.JSON(http.StatusOK, gin.H{"message": "subscribed"})
}

func (h *Handler) Unsubscribe(c *gin.Context) {
	// unsubscribes a customer from all lists
	c.JSON(http.StatusOK, gin.H{"message": "unsubscribed"})
}

func (h *Handler) GetSubscriber(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "one subscriber"})
}

func (h *Handler) GetSubscribers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "all subscribers"})
}
