package producer

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
)

func ProductionOrderUpdateProducer(ctx context_manager.ContextManager, ch <-chan []byte, sqsClient sqsiface.SQSAPI) error {
	queueUrl := ctx.Get("AWS_PRODUCTION_UPDATE_SQS_URL").(string)

	for message := range ch {
		_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
			QueueUrl:    aws.String(queueUrl),
			MessageBody: aws.String(string(message)),
		})
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			continue
		}
	}

	return nil
}
