package main

import (
	"fmt"

	"github.com/TutorialEdge/notification-service/internal/api"
)

func Run() error {
	fmt.Println("starting notification service")

	notificationAPI := api.New()
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
