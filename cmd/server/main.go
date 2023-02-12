package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"

	"github.com/TutorialEdge/notification-service/internal/api"
	"github.com/TutorialEdge/notification-service/internal/email"
	"github.com/TutorialEdge/notification-service/internal/list"
	"github.com/TutorialEdge/notification-service/internal/notification"
	"github.com/TutorialEdge/notification-service/internal/store"
	"github.com/TutorialEdge/notification-service/internal/subscriber"
)

func Run() error {
	fmt.Println("starting notification service")

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
		return fmt.Errorf("could not connect to database: %w", err)
	}

	err = store.MigrateDB(db)
	if err != nil {
		return err
	}

	serviceStore := store.New(db)
	emailService := email.New()

	listService := list.New()
	subscriberService := subscriber.New(serviceStore)
	notificationService := notification.New(serviceStore, emailService, subscriberService)
	notificationAPI := api.New(*notificationService, *subscriberService, *listService)
	if err := notificationAPI.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		panic(err.Error())
	}
}
