package product

import (
	"fiap-fast-food-ms-producao/adapter/database"
	"fiap-fast-food-ms-producao/domain/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreatePedido(c *gin.Context) {

	var req dto.ProductionOrderDTO
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	dbClient, _ := c.Get("db_client")
	mongoClient := dbClient.(database.DatabaseManger)
	collectionName := "production-order"

	productionOrder, _ := mongoClient.Create(collectionName, map[string]interface{}{
		"status": req.Status,
	})
	if productionOrder == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product Order Not Found"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}
