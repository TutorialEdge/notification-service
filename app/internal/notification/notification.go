package notification

import (
	"context"
	"log"

	"github.com/TutorialEdge/notification-service/internal/database"
	"github.com/TutorialEdge/notification-service/internal/models"
)

type Service struct {
	Store database.Store
}

func New(db database.Store) Service {
	return Service{
		Store: db,
	}
}

func (s *Service) PostNotification(ctx context.Context, noti models.Notification) error {
	log.Println("posting notification")

	return nil
}
