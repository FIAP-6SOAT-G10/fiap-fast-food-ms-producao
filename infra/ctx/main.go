package ctx

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/infra/db"
	"fmt"
	"log"
	"os"
	"strings"
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

func configContext(ctx *contextManager) error {
	envFileLocation := os.Getenv("ENV_FILE_LOCATION")
	if envFileLocation == "" {
		log.Fatal("ENV_FILE_LOCATION is not set")
	}
	fmt.Printf("%s", envFileLocation)
	viper.SetConfigFile(envFileLocation)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	for _, key := range viper.AllKeys() {
		fmt.Printf("%s -> %s\n", key, viper.Get(key))
		ctx.Set(strings.ToUpper(key), viper.Get(key))
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
		log.Fatalf("Error creating mongo client: %v", err)
	}
	ctx.Set("mongo_client", mongoClient)
	return &ctx
}
