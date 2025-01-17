package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// UsedPeer represents a requested peer by a user_service
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
	Interface     string             `json:"interface"`     // e.g., wg0
	PublicKey     string             `json:"publicKey"`     // e.g., AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA=
	Port          string             `json:"port"`          // e.g., (none) or 51820
	Endpoint      string             `json:"endpoint"`      // e.g., (none) or 10.10.0.2:51820
	TransferRx    string             `json:"transferRx"`    // e.g., 0
	TransferTx    string             `json:"transferTx"`    // e.g., 0
	LastHandshake string             `json:"lastHandshake"` // e.g., (none) or a Unix timestamp
	AllowedIPs    string             `json:"allowedIps"`    // Allowed IPs in CIDR notation (e.g., 10.10.0.x/32)
	Flags         string             `json:"flags"`         // e.g., "off"
	CreatedAt     string             `json:"createdAt"`
	UpdatedAt     string             `json:"updatedAt"`
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
