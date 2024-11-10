package context_manager

import "github.com/gin-gonic/gin"

type ContextManager interface {
	Set(key string, value any)
	Get(key string) any
	PassContext(obj *gin.Context)
}
