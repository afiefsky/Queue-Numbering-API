package model

import (
	"context"
	"time"

	"queue-numbering-api/config"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Order struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	CustomerName string             `bson:"customer_name" json:"customer_name"`
	Product      string             `bson:"product" json:"product"`
	Status       string             `bson:"status" json:"status"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
}

func (o *Order) Insert(ctx context.Context) (*mongo.InsertOneResult, error) {
	collection := config.MongoDB.Collection("orders")
	return collection.InsertOne(ctx, o)
}
