package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func Handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	log.Println("Processing Incoming Notification")
	for _, message := range sqsEvent.Records {
		log.Printf("%+v\n", message)
	}

	return nil
}

func main() {
	lambda.Start(Handler)
}
