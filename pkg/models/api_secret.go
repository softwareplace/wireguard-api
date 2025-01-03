package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ApiSecret struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Key       string             `json:"key" bson:"key"`
	Client    string             `json:"client" bson:"client"`
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
	Status    string             `json:"status"` // INACTIVE, ACTIVE, DELETED, EXPIRED
}
