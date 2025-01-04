package peer

import "github.com/softwareplace/wireguard-api/pkg/models"

type Repository interface {
	SaveAll(peers []models.Peer) error
	Save(peer models.Peer) error
	GetAvailablePeer() (*models.Peer, error)
}

type repositoryImpl struct{}

func GetRepository() Repository {
	return &repositoryImpl{}
}
