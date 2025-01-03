package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// UsedPeer represents a requested peer by a user
type UsedPeer struct {
	Username string `json:"username"`
	PeerData string `json:"peerData"`
	Status   string `json:"status"` // INACTIVE, AVAILABLE, IN_USE
}

// Peer represents all WireGuard peer data in the system
type Peer struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	PeerData  string             `json:"peerData"`
	FileName  string             `json:"fileName"`
	Status    string             `json:"status"` // INACTIVE, AVAILABLE, IN_USE
	CreatedAt string             `json:"created_at"`
	UpdatedAt string             `json:"updated_at"`
}
