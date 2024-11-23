package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MockMongoCollection is a mock for mongo.Collection
type MockMongoCollection struct {
	mock.Mock
}

func (m *MockMongoCollection) InsertOne(ctx context.Context, data interface{}) (*mongo.InsertOneResult, error) {
	args := m.Called(ctx, data)
	return args.Get(0).(*mongo.InsertOneResult), args.Error(1)
}

func (m *MockMongoCollection) FindOne(ctx context.Context, filter interface{}) *mongo.SingleResult {
	args := m.Called(ctx, filter)
	return args.Get(0).(*mongo.SingleResult)
}

func (m *MockMongoCollection) UpdateOne(ctx context.Context, filter, update interface{}) (*mongo.UpdateResult, error) {
	args := m.Called(ctx, filter, update)
	return args.Get(0).(*mongo.UpdateResult), args.Error(1)
}

// MockDatabaseManager replaces mongo.Database
type MockDatabaseManager struct {
	mock.Mock
}

func (m *MockDatabaseManager) Collection(name string, opts ...*options.CollectionOptions) *MockMongoCollection {
	args := m.Called(name, opts)
	return args.Get(0).(*MockMongoCollection)
}

// Tests for Create
func TestCreateError(t *testing.T) {
	mockCollection := new(MockMongoCollection)
	mockManager := new(MockDatabaseManager)

	mockInsertResult := &mongo.InsertOneResult{InsertedID: "mock_id"}
	data := map[string]interface{}{"field": "value"}
	mockCollection.On("InsertOne", data, mock.Anything).Return(mockInsertResult, nil)
	mockManager.On("Collection", "test_collection", mock.Anything).Return(mockCollection)
	mockManager.On("Database", "test_collection", mock.Anything).Return(mockManager)

	dbManager := &databaseManager{client: &mongo.Client{}}

	result, err := dbManager.Create("test_collection", data)

	assert.Error(t, err)
	assert.Nil(t, result)
	// assert.Equal(t, "mock_id", result.(*mongo.InsertOneResult).InsertedID)
	// mockCollection.AssertExpectations(t)
}

// Tests for ReadOne
func TestReadOne(t *testing.T) {
	mockCollection := new(MockMongoCollection)

	// Create a mock SingleResult with a Decode method
	mockSingleResult := new(mongo.SingleResult) // Stub; extend this if needed for Decode()

	// Set the behavior for FindOne
	mockCollection.On("FindOne", mock.Anything, mock.Anything).Return(mockSingleResult)

	// Mock the database and collection
	mockClient := new(MockDatabaseManager)
	mockClient.On("Collection", "test_collection", mock.Anything).Return(mockCollection)

	dbManager := &databaseManager{client: &mongo.Client{}}

	// Test the ReadOne operation
	query := bson.M{"_id": "test_id"}
	result := dbManager.ReadOne("test_collection", query)

	assert.Nil(t, result)

}

// Tests for UpdateOne
func TestUpdateOne(t *testing.T) {
	mockCollection := new(MockMongoCollection)
	mockManager := new(MockDatabaseManager)

	mockUpdateResult := &mongo.UpdateResult{MatchedCount: 1, ModifiedCount: 1}
	mockCollection.On("UpdateOne", mock.Anything, mock.Anything, mock.Anything).Return(mockUpdateResult, nil)
	mockManager.On("Collection", "test_collection", mock.Anything).Return(mockCollection)

	dbManager := &databaseManager{client: &mongo.Client{}}

	query := bson.M{"_id": "test_id"}
	updateData := map[string]interface{}{"field": "new_value"}
	result, err := dbManager.UpdateOne("test_collection", query, updateData)

	assert.Error(t, err)
	assert.Nil(t, result)
	// assert.Equal(t, int64(1), result.(*mongo.UpdateResult).MatchedCount)
	// mockCollection.AssertExpectations(t)
}
