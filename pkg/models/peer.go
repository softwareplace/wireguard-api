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
	Id            primitive.ObjectID `bson:"_id,omitempty"`
	PeerData      string             `json:"peerData"`
	FileName      string             `json:"fileName"`
	Status        string             `json:"status"`
	Interface     string             `json:"interface"`
	PublicKey     string             `json:"publicKey"`
	PrivateKey    string             `json:"privateKey"`
	Port          string             `json:"port"`
	Endpoint      string             `json:"endpoint"`
	TransferRx    string             `json:"transferRx"`
	TransferTx    string             `json:"transferTx"`
	LastHandshake string             `json:"lastHandshake"`
	AllowedIPs    string             `json:"allowedIps"`
	Flags         string             `json:"flags"`
	CreatedAt     string             `json:"created_at"`
	UpdatedAt     string             `json:"updated_at"`
}

type PeerResponse struct {
	PeerData string `json:"peerData"`
	FileName string `json:"fileName"`
}

func (p *Peer) ToResponse() *PeerResponse {
	return &PeerResponse{
		PeerData: p.PeerData,
		FileName: p.FileName,
	}
}
