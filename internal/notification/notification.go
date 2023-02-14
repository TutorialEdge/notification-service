package notification

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/TutorialEdge/ctxlog"
	"github.com/TutorialEdge/notification-service/internal/store"
	"github.com/TutorialEdge/notification-service/internal/subscriber"
	"github.com/google/uuid"
)

// Emailer - handles the low-level implementation details
// for sending emails.
type Emailer interface {
	SendEmail(context.Context, string) error
}

// SubscriberService - fetches the subscribers.
type SubscriberService interface {
	GetSubscriber(context.Context, string) (subscriber.Subscriber, error)
}

type Store interface {
	CreateNotification(context.Context, store.CreateNotificationParams) (store.Notification, error)
	DeleteNotification(context.Context, uuid.UUID) error
	GetNotification(context.Context, uuid.UUID) (store.Notification, error)
}

// Notifier - the notifier service
type Notifier struct {
	store      Store
	Emailer    Emailer
	SubService SubscriberService
	log        *ctxlog.CtxLogger
}

func New(
	store Store,
	emailer Emailer,
	subService SubscriberService,
	log *ctxlog.CtxLogger,
) *Notifier {
	return &Notifier{
		store:      store,
		SubService: subService,
		Emailer:    emailer,
		log:        log,
	}
}

// NotificationOpts - all of the necessary info needed to send
// a notification to a subscriber.
type NotificationOpts struct {
	NotificationID string
	SubscriberID   string
	Meta           map[string]any
}

// SendNotification - sends a notification to a user.
func (n *Notifier) SendNotification(ctx context.Context, opts NotificationOpts) error {
	n.log.Info(ctx, "sending a notification")
	// fetch the subscriber from the database
	sub, err := n.SubService.GetSubscriber(ctx, opts.SubscriberID)
	if err != nil {
		n.log.Error(ctx, fmt.Sprintf("error fetching subscriber: %s", err.Error()))
		return err
	}

	// only send notifiactions if the user is in a subscribed state
	if !sub.Subscribed {
		n.log.Error(ctx, "user is not subscribed - not sending notification")
		return errors.New("subscriber is unsubscribed, not sending notification")
	}

	err = n.Emailer.SendEmail(ctx, "")
	if err != nil {
		n.log.Error(ctx, err.Error())
		return err
	}

	n.log.Info(ctx, "notification sent successfully")
	return nil
}

// Notification - models a notification in the database
type Notification struct {
	NotificationName string
	NotificationID   string
	TemplateHTML     string
	CreatedAt        string
	UpdatedAt        string
}

// CreateNotification - creates a new notification type in the database which can then
// be sent to subscribers.
func (n *Notifier) CreateNotification(ctx context.Context, not Notification) (Notification, error) {
	n.log.Info(ctx, "creating a new notification")

	newNot, err := n.store.CreateNotification(ctx, store.CreateNotificationParams{
		NotificationName: sql.NullString{String: not.NotificationName, Valid: true},
		Html:             sql.NullString{String: not.TemplateHTML, Valid: true},
	})
	if err != nil {
		n.log.Error(ctx, err.Error())
		return Notification{}, err
	}

	n.log.Info(ctx, "new notification created")
	return Notification{
		NotificationName: newNot.NotificationName.String,
		TemplateHTML:     newNot.Html.String,
		CreatedAt:        newNot.CreatedAt.Time.String(),
		UpdatedAt:        newNot.UpdatedAt.Time.String(),
	}, nil
}

// UpdateNotification - allows for the updating of the HTML template for a given notification
func (n *Notifier) UpdateNotification(ctx context.Context, not Notification) error {
	n.log.Info(ctx, "updating notification")

	n.log.Info(ctx, "notification successfully updated")
	return errors.New("not implemented")
}

// DeleteNotification - does what it says on the tin.
func (n *Notifier) DeleteNotification(ctx context.Context, notificationID string) error {
	n.log.Info(ctx, "deleting notification")

	err := n.store.DeleteNotification(ctx, uuid.MustParse(notificationID))
	if err != nil {
		n.log.Error(ctx, err.Error())
		return err
	}

	n.log.Info(ctx, "notification successfully deleted")
	return errors.New("not implemented")
}
