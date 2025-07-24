package service

import (
	"context"
	"time"

	"queue-numbering-api/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateOrderRequest struct {
	CustomerName string `json:"customer_name"`
	Product      string `json:"product"`
}

func CreateOrder(ctx context.Context, req CreateOrderRequest) (*model.Order, error) {
	order := model.Order{
		ID:           primitive.NewObjectID(),
		CustomerName: req.CustomerName,
		Product:      req.Product,
		Status:       "waiting_for_payment",
		CreatedAt:    time.Now(),
	}

	_, err := order.Insert(ctx)
	if err != nil {
		return nil, err
	}

	return &order, nil
}
