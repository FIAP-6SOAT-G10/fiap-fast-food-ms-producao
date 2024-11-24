package worker

import (
	"fiap-fast-food-ms-producao/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestInitWorker(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockCtx := mocks.NewMockContextManager(ctrl)

	mockCtx.EXPECT().Get("aws_production_payment_sqs_url").Return("default_value").Times(1)
	mockCtx.EXPECT().Get("aws_access_key_id").Return("default_value").Times(1)
	mockCtx.EXPECT().Get("aws_secret_access_key").Return("default_value").Times(1)
	mockCtx.EXPECT().Get("aws_session_token").Return("default_value").Times(1)
	mockCtx.EXPECT().Get("aws_region").Return("default_value").Times(1)

	ch := make(chan map[string]interface{}, 10)
	worker, err := InitWorker(mockCtx, ch)

	assert.NotNil(t, worker)
	assert.Nil(t, err)
}

func TestPubSubChannelWorker(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockCtx := mocks.NewMockContextManager(ctrl)
	mockSqs := mocks.NewMockSQSAPI(ctrl)
	ch := make(chan map[string]interface{}, 10)

	broker := &BrokerMessageWorkerSQS{
		client:      mockSqs,
		queueUrl:    "queue-test",
		ctx:         mockCtx,
		messageChan: ch,
	}

	testMessage := map[string]interface{}{
		"nome": "Test",
	}

	broker.Produce(testMessage)

	select {
	case receivedMessage := <-ch:
		if receivedMessage["nome"] != testMessage["nome"] {
			t.Errorf("Expected message: %v, got: %v", testMessage, receivedMessage)
		}
	default:
		t.Errorf("No message received in the channel")
	}

}
