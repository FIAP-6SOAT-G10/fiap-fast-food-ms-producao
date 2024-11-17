package product

import (
	"bytes"
	"encoding/json"
	"fiap-fast-food-ms-producao/adapter/database"
	"fiap-fast-food-ms-producao/domain/dto"
	"fiap-fast-food-ms-producao/domain/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type databaseManagerMock struct{}

func (d *databaseManagerMock) Create(collection string, data map[string]interface{}) (any, error) {
	model := models.ProductionOrder{
		ID:     primitive.NewObjectID(),
		Status: 1,
	}
	return model, nil
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

func SharedContextMiddlewareMock(dbMock database.DatabaseManger) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("db_client", dbMock)
		c.Next()
	}
}

func TestCreatePedido(t *testing.T) {
	router := gin.New()
	router.Use(SharedContextMiddlewareMock(NewMockDatabase()))
	router.POST("/pedido", CreatePedido)
	body := dto.ProductionOrderDTO{
		Status: "Pending",
	}
	jsonValue, _ := json.Marshal(body)
	reqFound, _ := http.NewRequest("POST", "/pedido", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusCreated, w.Code)
}
