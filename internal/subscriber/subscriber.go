package subscriber

import (
	"context"
	"errors"
)

type Subscriber struct {
	SubscriberID string
	Email        string
	Subscribed   bool
}

type Store interface{}

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
