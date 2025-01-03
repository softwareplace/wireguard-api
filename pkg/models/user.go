package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// User represents a user object
type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `json:"username"`
	Password  string             `json:"password"`
	Email     string             `json:"email"`
	Salt      string             `json:"salt"`
	Role      string             `json:"role"`
	Status    string             `json:"status"` // INACTIVE, ACTIVE, DELETED
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
}

type UserUpdate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
