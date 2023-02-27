package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
	ID       primitive.ObjectID `json:"id"`
	Clothing *string            `json:"clothing"`
	Price    *float64           `json:"price"`
	Server   *string            `json:"server"`
	Table    *string            `json:"table"`
}
