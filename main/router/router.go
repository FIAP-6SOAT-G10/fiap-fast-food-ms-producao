package router

import (
	"fiap-fast-food-ms-producao/main/router/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()
	router.PUT("/pedido/:id/:status", api.UpdateStatusPedido)
	return router
}
