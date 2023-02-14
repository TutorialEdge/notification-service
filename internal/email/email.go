package email

import (
	"context"
	"os"
	"time"

	"github.com/TutorialEdge/ctxlog"
	"github.com/mailgun/mailgun-go/v3"
)

// The package that handles actually sending the emails downstream
type Service struct {
	mg  *mailgun.MailgunImpl
	log *ctxlog.CtxLogger
}

func New(log *ctxlog.CtxLogger) *Service {
	domain := os.Getenv("MAILGUN_DOMAIN")
	apiKey := os.Getenv("MAILGUN_API_KEY")
	mg := mailgun.NewMailgun(domain, apiKey)
	return &Service{
		log: log,
		mg:  mg,
	}
}

func (s *Service) SendEmail(ctx context.Context, html string) error {
	s.log.Info(ctx, "sending email notification")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	m := s.mg.NewMessage(
		"hello <support@myepitaph.com>",
		"Hello",
		"Testing some Mailgun awesomeness!",
		"YOU@YOUR_DOMAIN_NAME",
	)

	_, id, err := s.mg.Send(ctx, m)
	ctx = ctxlog.WithFields(ctx, ctxlog.Fields{
		"mailgun_id": id,
	})
	if err != nil {
		s.log.Error(ctx, err.Error())
		return err
	}
	s.log.Info(ctx, "email notification sent")
	return nil
}
