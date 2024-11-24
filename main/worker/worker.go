package worker

import (
	"encoding/json"
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/worker_manager"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type BrokerMessageWorkerSQS struct {
	client      *sqs.SQS
	ctx         context_manager.ContextManager
	queueUrl    string
	messageChan chan<- map[string]interface{}
}

func (b *BrokerMessageWorkerSQS) Consume() {
	for {
		messages, err := b.receiveMessages(b.queueUrl)
		if err != nil {
			log.Printf("Error receiving messages: %v", err)
			break
		}
		for _, msg := range messages {
			var result map[string]interface{}
			// Parse the JSON string into the map
			if err := json.Unmarshal([]byte(*msg.Body), &result); err != nil {
				fmt.Println("Error parsing JSON:", err)
				return
			}

			b.Produce(result)

			_, err := b.client.DeleteMessage(&sqs.DeleteMessageInput{
				QueueUrl:      aws.String(b.queueUrl),
				ReceiptHandle: msg.ReceiptHandle,
			})
			if err != nil {
				log.Printf("Error deleting message: %v", err)
				continue
			}

		}
	}
}

func (b *BrokerMessageWorkerSQS) Produce(message map[string]interface{}) {
	b.messageChan <- message
}

func (b *BrokerMessageWorkerSQS) receiveMessages(queueUrl string) ([]*sqs.Message, error) {
	out, err := b.client.ReceiveMessage(&sqs.ReceiveMessageInput{
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

func InitWorker(ctx context_manager.ContextManager, ch chan<- map[string]interface{}) (worker_manager.BrokerMessageConsumer, error) {

	queueUrl := ctx.Get("aws_production_payment_sqs_url")
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

	brokerMessage := BrokerMessageWorkerSQS{
		client:      sqsClient,
		queueUrl:    queueUrl.(string),
		ctx:         ctx,
		messageChan: ch,
	}

	go brokerMessage.Consume()
	return &brokerMessage, nil
}
