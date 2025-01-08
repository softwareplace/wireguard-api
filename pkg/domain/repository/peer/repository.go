package peer

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/db"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	SaveAll(peers []models.Peer) error
	Save(peer models.Peer) error
	Update(peer models.Peer) error
	FindByPublicKey(pubKey string) (*models.Peer, error)
	GetAvailablePeer() (*models.Peer, error)
}

type repositoryImpl struct {
	database *mongo.Database
}

func GetRepository() Repository {
	return &repositoryImpl{
		database: db.GetDB(),
	}
}

func (r *repositoryImpl) collection() *mongo.Collection {
	return r.database.Collection("peers")
}
