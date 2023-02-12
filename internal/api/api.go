package api

import (
	"net/http"

	"github.com/TutorialEdge/notification-service/internal/notification"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router              *gin.Engine
	notificationService notification.Notifier
}

func New(notifier notification.Notifier) *Handler {
	handler := &Handler{
		notificationService: notifier,
	}
	handler.setupRoutes()
	return handler
}

func (h *Handler) setupRoutes() {
	h.router = gin.Default()
	h.router.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
	})

	h.router.GET("/api/v1/subscribers", h.GetSubscribers)
	h.router.GET("/api/v1/subscriber/:subscriberid", h.GetSubscriber)

	h.router.POST("/api/v1/list", h.CreateList)

	h.router.POST("/api/v1/notification", h.CreateNotification)
	h.router.POST("/api/v1/notification/:notificationid", h.SendNotification)

	h.router.POST("/api/v1/subscribe", h.Subscribe)
	h.router.GET("/api/v1/unsubscribe", h.Unsubscribe)
}

func (h *Handler) Serve() error {
	if err := h.router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
