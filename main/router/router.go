package router

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/database"
	"fiap-fast-food-ms-producao/main/router/api/product"
	"fiap-fast-food-ms-producao/main/router/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(context context_manager.ContextManager, dbManager database.DatabaseManger, productionUpdateChannel chan<- []byte) *gin.Engine {
	router := gin.New()
	router.Use(middleware.SharedContextMiddleware(&context, dbManager, productionUpdateChannel))

	router.PUT("/pedido/:id/:status", product.UpdateStatusPedido)
	router.POST("/pedido", product.CreatePedido)

	return router
}
