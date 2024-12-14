package product

import (
	"encoding/json"
	"fiap-fast-food-ms-producao/adapter/database"
	"fiap-fast-food-ms-producao/domain/dto"
	"fiap-fast-food-ms-producao/domain/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateStatusPedido(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")

	dbClient, _ := c.Get("db_client")
	mongoClient := dbClient.(database.DatabaseManger)
	collectionName := "production-order"

	objId := id
	productionOrder := mongoClient.ReadOne(collectionName, map[string]interface{}{
		"externalId": objId,
	})
	if productionOrder == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Order Not Found"})
		return
	}

	newStatus, err := models.StatusFromString(status)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Status Not Found"})
		return
	}
	po, ok := productionOrder.(models.ProductionOrder)
	if ok {
		po.Status = int(newStatus)
	}

	query := bson.M{"externalId": objId}
	data := map[string]interface{}{
		"status": newStatus,
	}
	mongoClient.UpdateOne(collectionName, query, data)
	poDTO := dto.ToProductionOrderDTO(&po)

	poBytes, err := StructToBytes(*poDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	productionUpdateChannel, _ := c.Get("production_update_channel")

	productionUpdateChannel.(chan<- []byte) <- poBytes
	fmt.Println("OK")
	c.JSON(http.StatusOK, poDTO)
}

func StructToBytes(dto dto.ProductionOrderDTO) ([]byte, error) {
	// Marshal the struct to JSON
	jsonData, err := json.Marshal(dto)
	if err != nil {
		return nil, err
	}

	return jsonData, nil
}
