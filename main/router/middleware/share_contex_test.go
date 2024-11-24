package middleware

import (
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/database"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type databaseManagerMock struct{}

func (d *databaseManagerMock) Create(collection string, data map[string]interface{}) (any, error) {
	return nil, nil
}

func (d *databaseManagerMock) ReadOne(collection string, query map[string]interface{}) any {
	return nil
}

func (d *databaseManagerMock) UpdateOne(collection string, query any, data map[string]interface{}) (any, error) {
	return nil, nil
}

func (d *databaseManagerMock) Disconnect() error {
	return nil
}

func NewMockDatabase() database.DatabaseManger {
	return &databaseManagerMock{}
}

type contextManagerMock struct{}

func (c *contextManagerMock) Set(key string, value any) {}

func (c *contextManagerMock) Get(key string) any {
	return nil
}

func (c *contextManagerMock) PassContext(obj *gin.Context) {}

func NewContextManager() context_manager.ContextManager {
	return &contextManagerMock{}
}

func TestSharedContextMiddleware(t *testing.T) {
	// Mock dependencies
	mockContextManager := NewContextManager()
	mockDBManager := NewMockDatabase()
	mockProductionUpdateChannel := make(chan<- []byte)

	// Create a Gin router with the middleware
	router := gin.New()
	router.Use(SharedContextMiddleware(&mockContextManager, mockDBManager, mockProductionUpdateChannel))

	// Define a test handler to inspect the middleware's effect
	router.GET("/test", func(c *gin.Context) {
		// Check if the context has the expected values
		dbClient, exists := c.Get("db_client")
		assert.True(t, exists, "db_client should be set in the context")
		assert.Equal(t, mockDBManager, dbClient, "db_client value is incorrect")

		productionChannel, exists := c.Get("production_update_channel")
		assert.True(t, exists, "production_update_channel should be set in the context")
		assert.Equal(t, mockProductionUpdateChannel, productionChannel, "production_update_channel value is incorrect")

		// Indicate success
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Perform a test request
	w := performRequest(router, "GET", "/test")

	// Assert the response
	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), `"status":"ok"`)
}

// Helper function to perform a request
func performRequest(r *gin.Engine, method, path string) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	req := &http.Request{
		Method: method,
		URL:    &url.URL{Path: path},
	}
	r.ServeHTTP(w, req)
	return w
}
