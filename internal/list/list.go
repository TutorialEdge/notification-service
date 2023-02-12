package list

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type List struct {
	ListName string
}

type Service struct{}

func New() *Service {
	return &Service{}
}

func (s *Service) GetList(ctx context.Context, listID uuid.UUID) (List, error) {
	return List{}, errors.New("not implemented")
}

func (s *Service) CreateList(ctx context.Context, newList List) (List, error) {
	return List{}, errors.New("not implemented")
}
