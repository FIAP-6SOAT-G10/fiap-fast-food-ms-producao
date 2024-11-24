package router_test

import (
	"bytes"
	"encoding/json"
	"fiap-fast-food-ms-producao/infra/ctx"
	"fiap-fast-food-ms-producao/infra/db"
	"fiap-fast-food-ms-producao/main/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitRouter(t *testing.T) {
	ctx := ctx.NewContextManager()
	dbManager, _ := db.NewDatabaseManager(ctx)
	productionUpdateChannel := make(chan []byte)

	// Initialize the router
	r := router.InitRouter(ctx, dbManager, productionUpdateChannel)
	go func() {
		<-productionUpdateChannel
	}()
	// Define test cases
	tests := []struct {
		name       string
		method     string
		url        string
		body       interface{}
		statusCode int
	}{
		{
			name:       "Test UpdateStatusPedido",
			method:     http.MethodPut,
			url:        "/pedido/6737ea2d22a3d3067de89bbf/InProgress",
			body:       nil,
			statusCode: http.StatusOK,
		},
		{
			name:       "Test CreatePedido",
			method:     http.MethodPost,
			url:        "/pedido",
			body:       map[string]interface{}{"status": "done"},
			statusCode: http.StatusCreated,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var reqBody []byte
			if test.body != nil {
				var err error
				reqBody, err = json.Marshal(test.body)
				if err != nil {
					t.Fatalf("Failed to marshal body: %v", err)
				}
			}

			req := httptest.NewRequest(test.method, test.url, bytes.NewReader(reqBody))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()

			// Serve the request
			r.ServeHTTP(w, req)

			// Assert response status code
			assert.Equal(t, test.statusCode, w.Code)
		})
	}
}
