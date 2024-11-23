package ctx

import (
	"strings"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestContextManager(t *testing.T) {
	// Initialize the contextManager
	manager := &contextManager{
		envs: make(map[string]any),
		mu:   sync.RWMutex{},
	}

	// Test Set and Get methods
	t.Run("Set and Get", func(t *testing.T) {
		manager.Set("key1", "value1")
		manager.Set("key2", 123)

		assert.Equal(t, "value1", manager.Get("key1"))
		assert.Equal(t, 123, manager.Get("key2"))
		assert.Nil(t, manager.Get("nonexistent"))
	})

	// Test PassContext
	t.Run("PassContext", func(t *testing.T) {
		manager.Set("key1", "value1")
		manager.Set("key2", 123)

		ginContext := &gin.Context{
			Keys: make(map[string]any),
		}

		manager.PassContext(ginContext)

		assert.Equal(t, "value1", ginContext.Keys["key1"])
		assert.Equal(t, 123, ginContext.Keys["key2"])
		assert.Nil(t, ginContext.Keys["nonexistent"])
	})
}

func TestConfigContext(t *testing.T) {
	ctx := &contextManager{
		envs: make(map[string]any),
		mu:   sync.RWMutex{},
	}

	// Mock the environment file content
	envContent := `
TEST_KEY1=value1
TEST_KEY2=value2
`

	// Use viper's SetConfigType and ReadConfig to mock the environment file
	viper.SetConfigType("env") // Specify the type as "env"
	err := viper.ReadConfig(strings.NewReader(envContent))
	assert.NoError(t, err, "should read inline env content without errors")

	// Ensure viper has keys loaded
	assert.Equal(t, "value1", viper.GetString("TEST_KEY1"))
	assert.Equal(t, "value2", viper.GetString("TEST_KEY2"))

	// Call the function, passing the mocked config name
	err = configContext(ctx, ".mock.env") // Use mock name for clarity
	assert.NoError(t, err, "should not return an error")

	// Validate the context has the correct keys and values
	assert.Equal(t, "value1", ctx.Get("test_key1"))
	assert.Equal(t, "value2", ctx.Get("test_key2"))

}

func TestConfigContextFileNotFound(t *testing.T) {
	// Mock contextManager
	ctx := &contextManager{
		envs: make(map[string]any),
		mu:   sync.RWMutex{},
	}

	err := configContext(ctx, "nonexistent.env")

	assert.Error(t, err, "should return an error when file does not exist")
	assert.Equal(t, "no environment variables found in configuration", err.Error())
}
