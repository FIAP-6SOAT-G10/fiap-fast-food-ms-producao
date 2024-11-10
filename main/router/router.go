package router

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/database"
	"fiap-fast-food-ms-producao/main/router/api"
	"fiap-fast-food-ms-producao/main/router/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(context context_manager.ContextManager, dbManager database.DatabaseManger) *gin.Engine {
	router := gin.New()
	router.Use(middleware.SharedContextMiddleware(&context, dbManager))

	router.PUT("/pedido/:id/:status", api.UpdateStatusPedido)

	return router
}
