package ctx

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/infra/db"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
)

type contextManager struct {
	envs map[string]any
	mu   sync.RWMutex
}

func (c *contextManager) Set(key string, value any) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.envs[key] = value
}

func (c *contextManager) Get(key string) any {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.envs[key]
}

func (c *contextManager) PassContext(obj *gin.Context) {
	for key, value := range c.envs {
		obj.Set(key, value)
	}
}

func configContext(ctx *contextManager) error {
	for _, env := range os.Environ() {
		values := strings.Split(env, "=")
		ctx.Set(values[0], values[1])
	}
	return nil
}

func NewContextManager() context_manager.ContextManager {
	ctx := contextManager{
		envs: make(map[string]any),
	}
	configContext(&ctx)
	mongoClient, err := db.NewDatabaseManager(&ctx)
	if err != nil {
		log.Fatalf("Error creating mongo client")
	}
	ctx.Set("mongo_client", mongoClient)
	return &ctx
}
