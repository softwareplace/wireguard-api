package peer

import (
	"github.com/eliasmeireles/wireguard-api/pkg/domain/repository/peer"
	"github.com/eliasmeireles/wireguard-api/pkg/models"
)

type Service interface {
	repository() peer.Repository
	Load()
	GetAvailablePeer() (*models.Peer, error)
}

type serviceImpl struct {
}

func (s *serviceImpl) repository() peer.Repository {
	return peer.GetRepository()
}

func GetService() Service {
	return &serviceImpl{}
}
