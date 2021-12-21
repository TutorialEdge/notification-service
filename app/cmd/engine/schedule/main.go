package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func Handler(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	log.Println("Scheduling Notifications to Go Out")
	// q := os.Getenv("queueUrl")
	q := "https://sqs.us-east-1.amazonaws.com/853957954650/test-app-notificationsQueue"
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
	})
	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
	}
	svc := sqs.New(sess)
	_, err = svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl: &q,
		MessageAttributes: map[string]*sqs.MessageAttributeValue{
			"email": &sqs.MessageAttributeValue{
				DataType:    aws.String("String"),
				StringValue: aws.String("me@elliotf.dev"),
			},
		},
		MessageBody: aws.String("Hello World"),
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Println("Successfully published to SQS")
	return events.APIGatewayProxyResponse{
		Body:       "scheduling notifications to go out",
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
