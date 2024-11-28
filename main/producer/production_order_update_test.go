package producer

import (
	"fiap-fast-food-ms-producao/mocks"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"github.com/golang/mock/gomock"
)

type MockSQSClient struct {
	sqsiface.SQSAPI
	sendMessageFunc func(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error)
}

func (m *MockSQSClient) SendMessage(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
	if m.sendMessageFunc != nil {
		return m.sendMessageFunc(input)
	}
	return &sqs.SendMessageOutput{}, nil
}

func TestProductionOrderUpdateProducer(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCtx := mocks.NewMockContextManager(ctrl)
	mockCtx.EXPECT().Get("AWS_PRODUCTION_UPDATE_SQS_URL").Return("https://example-queue-url").Times(1)

	// Mock SQS Client
	mockSQSClient := &MockSQSClient{
		sendMessageFunc: func(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {
			if aws.StringValue(input.QueueUrl) != "https://example-queue-url" {
				t.Errorf("Unexpected QueueUrl: %v", aws.StringValue(input.QueueUrl))
			}
			if aws.StringValue(input.MessageBody) != "test message" {
				t.Errorf("Unexpected MessageBody: %v", aws.StringValue(input.MessageBody))
			}
			return &sqs.SendMessageOutput{}, nil
		},
	}

	// Create a channel for messages
	messageChannel := make(chan []byte, 1)
	messageChannel <- []byte("test message")
	close(messageChannel)

	// Call the function
	err := ProductionOrderUpdateProducer(mockCtx, messageChannel, mockSQSClient)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
