package db

import (
	"context"
	"fiap-fast-food-ms-producao/adapter/context_manager"
	"fiap-fast-food-ms-producao/adapter/database"
	"fiap-fast-food-ms-producao/domain/models"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type databaseManager struct {
	client *mongo.Client
}

const fiap_database string = "fiap-tech-challenge"

func (d *databaseManager) Create(collection string, data map[string]interface{}) (any, error) {
	database := d.client.Database(fiap_database)
	c := database.Collection(collection)
	insertOne, err := c.InsertOne(context.TODO(), data)
	if err != nil {
		return nil, err
	}
	return insertOne, nil
}

func (d *databaseManager) ReadOne(collection string, query map[string]interface{}) any {
	c := d.client.Database(fiap_database).Collection(collection)
	findOne := c.FindOne(context.TODO(), query)
	var result models.ProductionOrder
	if err := findOne.Decode(&result); err != nil {
		fmt.Println(err)
		return nil
	}
	return result
}

func (d *databaseManager) UpdateOne(collection string, query any, data map[string]interface{}) (any, error) {
	c := d.client.Database(fiap_database).Collection(collection)

	update := bson.M{"$set": data}

	updateResult, err := c.UpdateOne(context.TODO(), query, update)
	if err != nil {
		return nil, err
	}
	return updateResult, nil
}

func (d *databaseManager) Disconnect() error {
	return nil
}

func NewDatabaseManager(ctx context_manager.ContextManager) (database.DatabaseManger, error) {
	uri := ctx.Get("MONGO_URL")
	clientOptions := options.Client().ApplyURI(uri.(string))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	if err := client.Ping(context.TODO(), nil); err != nil {
		return nil, err
	}
	fmt.Println("Connected to MongoDB")
	return &databaseManager{client: client}, nil
}
