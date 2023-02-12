package subscriber

import (
	"context"
	"errors"

	"github.com/TutorialEdge/notification-service/internal/store"
)

type Subscriber struct {
	SubscriberID string
	Email        string
	Subscribed   bool
}

type Store interface {
	CreateSubscriber(context.Context, string) (store.Subscriber, error)
	Unsubscribe(context.Context, string) error
}

type Service struct {
	Store Store
}

func New(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetSubscriber(ctx context.Context, subID string) (Subscriber, error) {
	return Subscriber{}, errors.New("not implemented")
}

func (s *Service) CreateSubscriber(ctx context.Context, sub Subscriber) (Subscriber, error) {
	newSub, err := s.Store.CreateSubscriber(ctx, sub.Email)
	if err != nil {
		return Subscriber{}, err
	}
	return Subscriber{
		Email: newSub.Email,
	}, nil
}

func (s *Service) Unsubscribe(ctx context.Context, email string) error {
	err := s.Store.Unsubscribe(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
