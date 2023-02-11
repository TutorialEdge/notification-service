package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendNotification(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "notification sent"})
}

func (h *Handler) CreateNotification(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "notification created"})
}
