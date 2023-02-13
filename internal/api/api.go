package api

import (
	"context"
	"net/http"

	"github.com/TutorialEdge/ctxlog"
	"github.com/TutorialEdge/notification-service/internal/list"
	"github.com/TutorialEdge/notification-service/internal/notification"
	"github.com/TutorialEdge/notification-service/internal/subscriber"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router              *gin.Engine
	notificationService notification.Notifier
	subService          subscriber.Service
	listService         list.Service
	log                 *ctxlog.CtxLogger
}

func New(
	notifier notification.Notifier,
	subService subscriber.Service,
	listService list.Service,
	log *ctxlog.CtxLogger,
) *Handler {
	log.Info(context.Background(), "hello world")
	handler := &Handler{
		notificationService: notifier,
		subService:          subService,
		listService:         listService,
		log:                 log,
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
	h.router.GET("/api/v1/notification/:notificationid", h.GetNotification)
	h.router.PUT("/api/v1/notification/:notificationid", h.UpdateNotification)

	h.router.POST("/api/v1/subscribe", h.Subscribe)
	h.router.GET("/api/v1/unsubscribe", h.Unsubscribe)
}

func (h *Handler) Serve() error {
	if err := h.router.Run(":8080"); err != nil {
		return err
	}
	return nil
}
