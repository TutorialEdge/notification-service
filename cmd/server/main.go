package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/TutorialEdge/ctxlog"
	"github.com/TutorialEdge/notification-service/internal/api"
	"github.com/TutorialEdge/notification-service/internal/email"
	"github.com/TutorialEdge/notification-service/internal/list"
	"github.com/TutorialEdge/notification-service/internal/notification"
	"github.com/TutorialEdge/notification-service/internal/store"
	"github.com/TutorialEdge/notification-service/internal/subscriber"
)

func Run() error {
	ctx := context.Background()
	log := ctxlog.New(
		ctxlog.WithJSONFormat(),
	)
	log.Info(ctx, "starting notification service")

	connectionString := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_TABLE"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("SSL_MODE"),
	)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Error(ctx, err.Error())
		return fmt.Errorf("could not connect to database: %w", err)
	}

	err = store.MigrateDB(db)
	if err != nil {
		log.Error(ctx, err.Error())
		return err
	}

	serviceStore := store.New(db)
	emailService := email.New(log)

	listService := list.New(
		serviceStore,
		log,
	)
	subscriberService := subscriber.New(
		serviceStore,
		log,
	)
	notificationService := notification.New(
		serviceStore,
		emailService,
		subscriberService,
		log,
	)
	notificationAPI := api.New(
		*notificationService,
		*subscriberService,
		*listService,
		log,
	)
	if err := notificationAPI.Serve(); err != nil {
		log.Error(ctx, err.Error())
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		panic(err.Error())
	}
}
