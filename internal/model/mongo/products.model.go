package mongo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Products struct {
	ID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"product_name,omitempty" bson:"product_name,omitempty"`
	Quantity int                `json:"quantity,omitempty" bson:"quantity,omitempty"`
}
