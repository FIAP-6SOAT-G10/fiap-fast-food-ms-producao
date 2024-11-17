package product

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func NewProductionUpdateChannel() chan []byte {
	productionUpdateChannel := make(chan []byte)
	return productionUpdateChannel
}
func ReadChannel(ch chan []byte) {
	<-ch
}

func TestUpdateStatusPedido(t *testing.T) {
	productionUpdateChannel := NewProductionUpdateChannel()
	go ReadChannel(productionUpdateChannel)
	router := gin.New()
	router.Use(SharedContextMiddlewareMock(NewMockDatabase(), productionUpdateChannel))
	router.PUT("/pedido/:id/:status", UpdateStatusPedido)
	url := "/pedido/6737ea2d22a3d3067de89bbf/InProgress"
	reqFound, _ := http.NewRequest("PUT", url, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, reqFound)
	assert.Equal(t, http.StatusOK, w.Code)

}
