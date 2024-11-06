package worker

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func InitWorker() {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Endpoint: aws.String(""),
			Region:   aws.String("us-east-1"),
		}))

	sqsClient := sqs.New(sess)

	fmt.Print("Started Worker")

	for {
		messages, err := receiveMessages(sqsClient)
		if err != nil {
			log.Printf("Error receiving messages: %v", err)
			continue
		}
		for _, msg := range messages {
			log.Printf("Received message: %s", *msg.Body)
		}
	}
}

func receiveMessages(sqsClient *sqs.SQS) ([]*sqs.Message, error) {
	queueUrl := "https://sqs.us-east-1.amazonaws.com/730335514370/production-order-updates-queue"
	out, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(10),
		VisibilityTimeout:   aws.Int64(10),
	})
	if err != nil {
		return nil, err
	}
	return out.Messages, nil
}
