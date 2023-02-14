package list

import (
	"context"

	"github.com/TutorialEdge/ctxlog"
	"github.com/TutorialEdge/notification-service/internal/store"
	"github.com/google/uuid"
)

type Store interface {
	GetList(context.Context, uuid.UUID) (store.List, error)
	CreateList(context.Context, string) (store.List, error)
}

type List struct {
	ListName string
}

type Service struct {
	store Store
	log   *ctxlog.CtxLogger
}

func New(
	store Store,
	log *ctxlog.CtxLogger,
) *Service {
	return &Service{
		store: store,
		log:   log,
	}
}

func (s *Service) GetList(ctx context.Context, listID uuid.UUID) (List, error) {
	s.log.Info(ctx, "getting list")
	l, err := s.store.GetList(ctx, listID)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return List{}, err
	}
	s.log.Info(ctx, "list successfully fetched")
	return List{
		ListName: l.ListName,
	}, nil
}

func (s *Service) CreateList(ctx context.Context, listName string) (List, error) {
	s.log.Info(ctx, "creating a list")
	insertedList, err := s.store.CreateList(ctx, listName)
	if err != nil {
		s.log.Error(ctx, err.Error())
		return List{}, err
	}
	s.log.Info(ctx, "list successfully created")
	return List{
		ListName: insertedList.ListName,
	}, nil
}
