package api

import (
	"fiap-fast-food-ms-producao/adapter/database"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStatusPedido(c *gin.Context) {
	id := c.Param("id")
	status := c.Param("status")
	dbClient, _ := c.Get("db_client")
	mongoClient := dbClient.(database.DatabaseManger)
	result, err := mongoClient.Create("production-order", map[string]interface{}{
		"name": "Gabs",
	})
	if err != nil {
		log.Fatalf("Error while try to create data inside MongoDB")
	}
	fmt.Println(result)
	c.JSON(http.StatusOK, gin.H{"id": id, "status": status})
}
