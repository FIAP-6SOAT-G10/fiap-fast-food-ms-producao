package ctx

import (
	"errors"
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/infra/db"
	"log"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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

func configContext(ctx *contextManager, envName string) error {
	viper.SetConfigFile(envName)
	if err := viper.ReadInConfig(); err != nil {
		return errors.New("no environment variables found in configuration")
	}

	for _, key := range viper.AllKeys() {
		ctx.Set(key, viper.Get(key))
	}
	return nil
}

func NewContextManager() context_manager.ContextManager {
	ctx := contextManager{
		envs: make(map[string]any),
	}
	configContext(&ctx, "/home/gabs/Documents/projetos/fiap-fast-food-ms-producao/.env")
	mongoClient, err := db.NewDatabaseManager(&ctx)
	if err != nil {
		log.Fatalf("Error creating mongo client")
	}
	ctx.Set("mongo_client", mongoClient)
	return &ctx
}
