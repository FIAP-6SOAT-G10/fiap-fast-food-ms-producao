package producer

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

func ProductionOrderUpdateProducer(ctx context_manager.ContextManager, ch <-chan []byte) error {
	queueUrl := ctx.Get("aws_production_update_sqs_url")
	aws_region := ctx.Get("aws_region")
	aws_access_key_id := ctx.Get("aws_access_key_id")
	aws_secret_access_key := ctx.Get("aws_secret_access_key")
	aws_session_token := ctx.Get("aws_session_token")

	sess := session.Must(session.NewSession(
		&aws.Config{
			Endpoint:    aws.String(""),
			Region:      aws.String(aws_region.(string)),
			Credentials: credentials.NewStaticCredentials(aws_access_key_id.(string), aws_secret_access_key.(string), aws_session_token.(string)),
		},
	))

	for message := range ch {
		sqsClient := sqs.New(sess)
		_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    aws.String(queueUrl.(string)),
			MessageBody: aws.String(string(message)),
		})

		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}

	}

	return nil
}
