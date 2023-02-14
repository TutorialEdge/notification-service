package subscriber

import (
	"context"

	"github.com/TutorialEdge/notification-service/internal/store"
	"github.com/google/uuid"
)

type Subscriber struct {
	SubscriberID string
	Email        string
	Subscribed   bool
}

type Store interface {
	CreateSubscriber(context.Context, string) (store.Subscriber, error)
	GetSubscribers(context.Context, int32) ([]store.Subscriber, error)
	GetSubscriber(context.Context, uuid.UUID) (store.Subscriber, error)
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

type PaginatedResponse struct {
	Page  int32
	Limit int32
	Data  []Subscriber
}

func (s *Service) GetSubscribers(ctx context.Context, limit, page int32) (PaginatedResponse, error) {
	subs, err := s.Store.GetSubscribers(ctx, limit)
	if err != nil {
		return PaginatedResponse{}, err
	}

	var normalizedSubs []Subscriber
	for _, sub := range subs {
		normalizedSubs = append(normalizedSubs, Subscriber{
			SubscriberID: sub.SubscriberID.String(),
			Email:        sub.Email,
			Subscribed:   sub.IsSubscribed.Bool,
		})
	}

	return PaginatedResponse{
		Page:  page,
		Limit: limit,
		Data:  normalizedSubs,
	}, nil
}

func (s *Service) GetSubscriber(ctx context.Context, subID string) (Subscriber, error) {
	sub, err := s.Store.GetSubscriber(ctx, uuid.MustParse(subID))
	if err != nil {
		return Subscriber{}, err
	}
	return Subscriber{
		SubscriberID: sub.SubscriberID.String(),
		Email:        sub.Email,
		Subscribed:   sub.IsSubscribed.Bool,
	}, nil
}

func (s *Service) CreateSubscriber(ctx context.Context, sub Subscriber) (Subscriber, error) {
	newSub, err := s.Store.CreateSubscriber(ctx, sub.Email)
	if err != nil {
		return Subscriber{}, err
	}
	return Subscriber{
		SubscriberID: newSub.SubscriberID.String(),
		Email:        newSub.Email,
		Subscribed:   newSub.IsSubscribed.Bool,
	}, nil
}

func (s *Service) Unsubscribe(ctx context.Context, email string) error {
	err := s.Store.Unsubscribe(ctx, email)
	if err != nil {
		return err
	}
	return nil
}
