package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductionOrder struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Status     int                `bson:"status"`
	ExternalId string             `bson:"externalId"`
}
