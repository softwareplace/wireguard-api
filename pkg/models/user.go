package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// User represents a user_service object
type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username  string             `json:"username" yaml:"username"`
	Password  string             `json:"password" yaml:"password"`
	Email     string             `json:"email" yaml:"email"`
	Salt      string             `json:"salt" yaml:"salt"`
	Roles     []string           `json:"roles" yaml:"roles"`
	Status    string             `json:"status" yaml:"status"` // INACTIVE, ACTIVE, DELETED
	CreatedAt string             `json:"createdAt" yaml:"created-at"`
	UpdatedAt string             `json:"updatedAt" yaml:"updated-at"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
