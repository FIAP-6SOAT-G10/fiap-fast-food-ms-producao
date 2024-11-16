package worker

import (
	"encoding/json"
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/database"
	"fiap-fast-food-ms-producao/domain/models"
)

func ProductionOrderConsumer(ctx context_manager.ContextManager, dbManager database.DatabaseManger, ch <-chan map[string]interface{}) {
	collectionName := "production-order"
	for message := range ch {
		var messageMap any
		if messageProductOrder, ok := message["Message"].(string); ok {
			if err := json.Unmarshal([]byte(messageProductOrder), &messageMap); err != nil {
				messageMap = messageProductOrder
			}
		}
		message["Message"] = messageMap
		message["status"] = models.Pending
		dbManager.Create(collectionName, message)
	}
}
