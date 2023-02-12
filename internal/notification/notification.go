package notification

import (
	"context"
	"errors"

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
}

func New(store Store, emailer Emailer, subService SubscriberService) *Notifier {
	return &Notifier{
		store:      store,
		SubService: subService,
		Emailer:    emailer,
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
	// fetch the subscriber from the database
	sub, err := n.SubService.GetSubscriber(ctx, opts.SubscriberID)
	if err != nil {
		return err
	}

	// only send notifiactions if the user is in a subscribed state
	if !sub.Subscribed {
		return errors.New("subscriber is unsubscribed, not sending notification")
	}

	err = n.Emailer.SendEmail(ctx, "")
	if err != nil {
		return err
	}

	return errors.New("not implemented")
}

// Notification - models a notification in the database
type Notification struct {
	NotificationID string
	TemplateHTML   string
}

// CreateNotification - creates a new notification type in the database which can then
// be sent to subscribers.
func (n *Notifier) CreateNotification(ctx context.Context, not Notification) error {

	return errors.New("not implemented")
}

// UpdateNotification - allows for the updating of the HTML template for a given notification
func (n *Notifier) UpdateNotification(ctx context.Context, not Notification) error {
	return errors.New("not implemented")
}

// DeleteNotification - does what it says on the tin.
func (n *Notifier) DeleteNotification(ctx context.Context, notificationID string) error {
	return errors.New("not implemented")
}
