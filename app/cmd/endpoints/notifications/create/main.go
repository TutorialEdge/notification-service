package main

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/TutorialEdge/notification-service/internal/database"
	"github.com/TutorialEdge/notification-service/internal/models"
	"github.com/TutorialEdge/notification-service/internal/notification"
)

var db database.Store

func init() {
	var err error
	if db, err = database.SetupDB(); err != nil {
		log.Println("failed to connect to db")
		log.Println(err.Error())
	}
	log.SetFormatter(&log.JSONFormatter{})
}

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	// 1. parse request body
	// 2. store notification in the database

	notificationService := notification.New(db)

	var noti models.Notification
	_ = json.Unmarshal([]byte(request.Body), &noti)
	err := notificationService.PostNotification(context.Background(), noti)
	if err != nil {
		return events.APIGatewayProxyResponse{
			Body:       "Error posting notification",
			StatusCode: 500,
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Notification submitted",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
