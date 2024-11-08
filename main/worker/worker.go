package worker

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func InitWorker(ctx context_manager.ContextManager) {

	queueUrl := ctx.Get("aws_sqs_url")
	aws_access_key_id := ctx.Get("aws_access_key_id")
	aws_secret_access_key := ctx.Get("aws_secret_access_key")
	aws_session_token := ctx.Get("aws_session_token")
	aws_region := ctx.Get("aws_region")

	sess := session.Must(session.NewSession(
		&aws.Config{
			Endpoint:    aws.String(""),
			Region:      aws.String(aws_region.(string)),
			Credentials: credentials.NewStaticCredentials(aws_access_key_id.(string), aws_secret_access_key.(string), aws_session_token.(string)),
		}))
	sqsClient := sqs.New(sess)

	fmt.Print("Started Worker")

	for {
		messages, err := receiveMessages(ctx, sqsClient)
		if err != nil {
			log.Printf("Error receiving messages: %v", err)
			continue
		}
		for _, msg := range messages {
			log.Printf("Received message: %s", *msg.Body)

			_, err := sqsClient.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(queueUrl.(string)),
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("Error deleting message: %v", err)
				continue
			}

		}
	}
}

func receiveMessages(ctx context_manager.ContextManager, sqsClient *sqs.SQS) ([]*sqs.Message, error) {
	queueUrl := ctx.Get("aws_sqs_url")
	out, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueUrl.(string)),
		MaxNumberOfMessages: aws.Int64(10),
		WaitTimeSeconds:     aws.Int64(10),
		VisibilityTimeout:   aws.Int64(10),
	})
	if err != nil {
		return nil, err
	}
	return out.Messages, nil
}
