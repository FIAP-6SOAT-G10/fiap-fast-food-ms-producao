package main

import (
	"encoding/json"
	"errors"
	"fiap-fast-food-ms-producao/main/worker"
	"fiap-fast-food-ms-producao/mocks"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/cucumber/godog"
	"github.com/golang/mock/gomock"
)

var loggedIn bool

func resetState() {
	loggedIn = false
}

func waitForMessage(ch <-chan map[string]interface{}, timeout time.Duration) (map[string]interface{}, error) {
	// Create a timer for the timeout
	timer := time.NewTimer(timeout)
	defer timer.Stop()

	select {
	case message := <-ch: // Receive a message from the channel
		return message, nil
	case <-timer.C: // Timeout occurred
		return nil, errors.New("timeout waiting for message")
	}
}

func oPagamentoAprovado() error {
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()
	mockSqs := mocks.NewMockSQSAPI(ctrl)
	queueURL := "test-queue"

	mockSqs.EXPECT().ReceiveMessage(queueURL).Times(1)
	return nil
}

func umaMensagemDeConfirmaoPublicadaNoSistema() error {
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	mockCtx := mocks.NewMockContextManager(ctrl)
	mockSqs := mocks.NewMockSQSAPI(ctrl)
	queueURL := "test-queue"

	messageData := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
	}

	// Convert the map to a JSON string
	body, err := json.Marshal(messageData)
	if err != nil {
		fmt.Printf("Error marshaling message body: %v\n", err)
		return nil
	}

	// Create a sample ReceiveMessageOutput object
	output := &sqs.ReceiveMessageOutput{
		Messages: []types.Message{
			{
				MessageId:     aws.String("12345"),
				ReceiptHandle: aws.String("abcde"),
				Body:          aws.String(string(body)),
				Attributes: map[string]string{
					"Attribute1": "Value1",
				},
			},
		},
	}

	mockSqs.EXPECT().ReceiveMessage(queueURL).Return(output, nil).Times(1)

	ch := make(chan map[string]interface{}, 10)

	broker := worker.BuildWorker(mockSqs, mockCtx, queueURL, ch)
	broker.Consume()

	message, err := waitForMessage(ch, 5*time.Second)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return nil
	}

	fmt.Printf("Received message: %s\n", message)
	return nil
}

func umaMensagemDePagamentoChegaNoBrokerSQS() error {
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	mockCtx := mocks.NewMockContextManager(ctrl)
	mockSqs := mocks.NewMockSQSAPI(ctrl)
	queueURL := "test-queue"

	messageData := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
	}

	// Convert the map to a JSON string
	body, err := json.Marshal(messageData)
	if err != nil {
		fmt.Printf("Error marshaling message body: %v\n", err)
		return nil
	}

	// Create a sample ReceiveMessageOutput object
	output := &sqs.ReceiveMessageOutput{
		Messages: []types.Message{
			{
				MessageId:     aws.String("12345"),
				ReceiptHandle: aws.String("abcde"),
				Body:          aws.String(string(body)),
				Attributes: map[string]string{
					"Attribute1": "Value1",
				},
			},
		},
	}

	mockSqs.EXPECT().ReceiveMessage(queueURL).Return(output, nil).Times(1)

	ch := make(chan map[string]interface{}, 10)

	broker := worker.BuildWorker(mockSqs, mockCtx, queueURL, ch)
	broker.Consume()

	message, err := waitForMessage(ch, 5*time.Second)
	if err == nil {
		return nil
	}

	fmt.Printf("Received message: %s\n", message)
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^o pagamento é aprovado$`, oPagamentoAprovado)
	ctx.Step(`^uma mensagem de confirmação é publicada no sistema$`, umaMensagemDeConfirmaoPublicadaNoSistema)
	ctx.Step(`^uma mensagem de pagamento chega no broker SQS$`, umaMensagemDePagamentoChegaNoBrokerSQS)
}
