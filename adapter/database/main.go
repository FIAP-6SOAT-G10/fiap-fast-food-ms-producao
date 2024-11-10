package database

type DatabaseManger interface {
	Create(collection string, data map[string]interface{}) (any, error)
	ReadOne(collection string, query map[string]interface{}) any
	UpdateOne(collection string, query any, data map[string]interface{}) (any, error)
	Disconnect() error
}
