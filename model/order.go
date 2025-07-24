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
	OrderCode    string             `bson:"order_code" json:"order_code"`
	CustomerName string             `bson:"customer_name" json:"customer_name"`
	Product      string             `bson:"product" json:"product"`
	Status       string             `bson:"status" json:"status"`
	QueueNumber  int64              `bson:"queue_number,omitempty" json:"queue_number,omitempty"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
}

func (o *Order) Insert(ctx context.Context) (*mongo.InsertOneResult, error) {
	collection := config.MongoDB.Collection("orders")
	return collection.InsertOne(ctx, o)
}

func UpdateQueueNumber(ctx context.Context, orderCode string, queueNumber int64) error {
	collection := config.MongoDB.Collection("orders")
	filter := map[string]interface{}{"order_code": orderCode}
	update := map[string]interface{}{
		"$set": map[string]interface{}{
			"queue_number": queueNumber,
		},
	}
	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}
