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

	mockCtx.EXPECT().Get("AWS_PRODUCTION_PAYMENT_SQS_URL").Return("default_value").Times(1)
	mockCtx.EXPECT().Get("AWS_ACCESS_KEY_ID").Return("default_value").Times(1)
	mockCtx.EXPECT().Get("AWS_SECRET_ACCESS_KEY").Return("default_value").Times(1)
	mockCtx.EXPECT().Get("AWS_SESSION_TOKEN").Return("default_value").Times(1)
	mockCtx.EXPECT().Get("AWS_REGION").Return("default_value").Times(1)

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
