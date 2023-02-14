package subscriber

import (
	"context"

	"github.com/TutorialEdge/ctxlog"
	"github.com/TutorialEdge/notification-service/internal/store"
	"github.com/google/uuid"
)

type Subscriber struct {
	SubscriberID string
	Email        string
	Subscribed   bool
	CreatedAt    string
	UpdatedAt    string
	DeletedAt    string
}

type Store interface {
	CreateSubscriber(context.Context, string) (store.Subscriber, error)
	GetSubscribers(context.Context, int32) ([]store.Subscriber, error)
	GetSubscriber(context.Context, uuid.UUID) (store.Subscriber, error)
	Unsubscribe(context.Context, string) error
}

type Service struct {
	Store Store
	log   *ctxlog.CtxLogger
}

func New(
	store Store,
	log *ctxlog.CtxLogger,
) *Service {
	return &Service{
		log:   log,
		Store: store,
	}
}

type PaginatedResponse struct {
	Page  int32
	Limit int32
	Data  []Subscriber
}

func (s *Service) GetSubscribers(ctx context.Context, limit, page int32) (PaginatedResponse, error) {
	s.log.Info(ctx, "getting subscribers")
	subs, err := s.Store.GetSubscribers(ctx, limit)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return PaginatedResponse{}, err
	}

	var normalizedSubs []Subscriber
	for _, sub := range subs {
		normalizedSubs = append(normalizedSubs, Subscriber{
			SubscriberID: sub.SubscriberID.String(),
			Email:        sub.Email,
			Subscribed:   sub.IsSubscribed.Bool,
			CreatedAt:    sub.CreatedAt.Time.String(),
			UpdatedAt:    sub.UpdatedAt.Time.String(),
			DeletedAt:    sub.DeletedAt.Time.String(),
		})
	}
	s.log.Info(ctx, "returning subscribers")
	return PaginatedResponse{
		Page:  page,
		Limit: limit,
		Data:  normalizedSubs,
	}, nil
}

func (s *Service) GetSubscriber(ctx context.Context, subID string) (Subscriber, error) {
	s.log.Info(ctx, "getting single subscriber")
	sub, err := s.Store.GetSubscriber(ctx, uuid.MustParse(subID))
	if err != nil {
		s.log.Error(ctx, err.Error())
		return Subscriber{}, err
	}

	s.log.Info(ctx, "returning single subscriber")
	return Subscriber{
		SubscriberID: sub.SubscriberID.String(),
		Email:        sub.Email,
		Subscribed:   sub.IsSubscribed.Bool,
		CreatedAt:    sub.CreatedAt.Time.String(),
		UpdatedAt:    sub.UpdatedAt.Time.String(),
		DeletedAt:    sub.DeletedAt.Time.String(),
	}, nil
}

func (s *Service) CreateSubscriber(ctx context.Context, sub Subscriber) (Subscriber, error) {
	s.log.Info(ctx, "creating a single subscriber")
	newSub, err := s.Store.CreateSubscriber(ctx, sub.Email)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return Subscriber{}, err
	}

	s.log.Info(ctx, "new subscriber created")
	return Subscriber{
		SubscriberID: newSub.SubscriberID.String(),
		Email:        newSub.Email,
		Subscribed:   newSub.IsSubscribed.Bool,
		CreatedAt:    newSub.CreatedAt.Time.String(),
		UpdatedAt:    newSub.UpdatedAt.Time.String(),
		DeletedAt:    newSub.DeletedAt.Time.String(),
	}, nil
}

func (s *Service) Unsubscribe(ctx context.Context, email string) error {
	s.log.Info(ctx, "unsubscribing subscriber")
	err := s.Store.Unsubscribe(ctx, email)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	s.log.Info(ctx, "user unsubscribed")
	return nil
}
