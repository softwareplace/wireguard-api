package peer

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/peer"
	"github.com/softwareplace/wireguard-api/pkg/models"
)

type Service interface {
	repository() peer.Repository
	Load()
	GetAvailablePeer() (response *models.PeerResponse, error error, notfound bool)
}

type serviceImpl struct {
}

func (s *serviceImpl) repository() peer.Repository {
	return peer.GetRepository()
}

func GetService() Service {
	return &serviceImpl{}
}
