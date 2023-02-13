package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) SendNotification(c *gin.Context) {
	h.log.Info(c.Request.Context(), "sending a notification")
	c.JSON(http.StatusOK, gin.H{"message": "notification sent"})
}

func (h *Handler) CreateNotification(c *gin.Context) {
	h.log.Info(c.Request.Context(), "creating a notification")
	c.JSON(http.StatusOK, gin.H{"message": "notification created"})
}

func (h *Handler) UpdateNotification(c *gin.Context) {
	h.log.Info(c.Request.Context(), "updating a notification")
	c.JSON(http.StatusOK, gin.H{"message": "notification created"})
}

func (h *Handler) GetNotification(c *gin.Context) {
	h.log.Info(c.Request.Context(), "get a notification")
	c.JSON(http.StatusOK, gin.H{"message": "notification retrieved"})
}
