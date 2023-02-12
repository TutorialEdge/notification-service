package email

import (
	"context"
	"errors"
)

// The package that handles actually sending the emails downstream
type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) SendEmail(ctx context.Context, html string) error {
	return errors.New("not implemented")
}
