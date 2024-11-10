package middleware

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/database"

	"github.com/gin-gonic/gin"
)

func SharedContextMiddleware(context *context_manager.ContextManager, dbManager database.DatabaseManger) gin.HandlerFunc {
	return func(c *gin.Context) {
		(*context).PassContext(c)
		c.Set("db_client", dbManager)
		c.Next()
	}
}
