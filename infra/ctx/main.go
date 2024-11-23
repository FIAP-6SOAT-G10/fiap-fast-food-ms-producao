package ctx

import (
	"errors"
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/infra/db"
	"log"
	"reflect"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
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

type MockContextManager struct {
	ctrl     *gomock.Controller
	recorder *MockContextManagerMockRecorder
}

// MockContextManagerMockRecorder is the mock recorder for MockContextManager.
type MockContextManagerMockRecorder struct {
	mock *MockContextManager
}

// NewMockContextManager creates a new mock instance.
func NewMockContextManager(ctrl *gomock.Controller) *MockContextManager {
	mock := &MockContextManager{ctrl: ctrl}
	mock.recorder = &MockContextManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContextManager) EXPECT() *MockContextManagerMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockContextManager) Get(key string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockContextManagerMockRecorder) Get(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockContextManager)(nil).Get), key)
}

// Set mocks base method.
func (m *MockContextManager) Set(key string, value interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Set", key, value)
}

// Set indicates an expected call of Set.
func (mr *MockContextManagerMockRecorder) Set(key, value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockContextManager)(nil).Set), key, value)
}
