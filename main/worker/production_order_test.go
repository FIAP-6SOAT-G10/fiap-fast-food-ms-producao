package worker

import (
	"fiap-fast-food-ms-producao/adapter/database"
	"testing"

	"github.com/stretchr/testify/assert"
)

type databaseManagerMock struct {
	CreateCalled bool
}

func (d *databaseManagerMock) Create(collection string, data map[string]interface{}) (any, error) {
	d.CreateCalled = true
	return nil, nil
}

func (d *databaseManagerMock) ReadOne(collection string, query map[string]interface{}) any {
	return nil
}

func (d *databaseManagerMock) UpdateOne(collection string, query any, data map[string]interface{}) (any, error) {
	return nil, nil
}

func (d *databaseManagerMock) Disconnect() error {
	return nil
}

func NewMockDatabase() database.DatabaseManger {
	return &databaseManagerMock{}
}

func TestProductionOrderConsumer(t *testing.T) {
	productionOrderConsumerChannel := make(chan map[string]interface{})
	go func() {
		productionOrderConsumerChannel <- map[string]interface{}{
			"Type":             "Notification",
			"MessageId":        "149689f0-8fc0-51d3-ae02-d653f9a8d227",
			"TopicArn":         "arn:aws:sns:us-east-1:317575625452:updates-payments-topic",
			"Message":          "{\"internalPaymentId\":\"27c5d210-0880-4a3b-b2b5-c4f891843588\",\"externalPaymentId\":92422075280,\"externalId\":\"12345678900\",\"payer\":\"email@email.com.br\",\"paymentAmount\":12.40,\"paymentDate\":\"2024-11-04T22:25:38.9838217\",\"paymentRequestDate\":\"2024-11-04T21:21:15.100804\",\"paymentStatus\":\"paid\",\"paymentMethod\":\"account_money\",\"paymentType\":\"account_money\"}",
			"Timestamp":        "2024-11-05T01:26:00.883Z",
			"SignatureVersion": "1",
		}
	}()
	db := NewMockDatabase()
	mockDb, _ := db.(*databaseManagerMock)
	ProductionOrderConsumer(db, productionOrderConsumerChannel)
	assert.Equal(t, mockDb.CreateCalled, true)
}
