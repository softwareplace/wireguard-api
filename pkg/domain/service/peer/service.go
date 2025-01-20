package peer

import (
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/peer"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"sync"
)

type Service interface {
	repository() peer.Repository
	Load()
	GetAvailablePeer() (response *models.PeerResponse, error error, notfound bool)
	Stream([]models.Peer) error
}

type serviceImpl struct {
}

func (s *serviceImpl) repository() peer.Repository {
	return peer.GetRepository()
}

var (
	serviceInstance Service
	serviceOnce     sync.Once
)

func GetService() Service {
	serviceOnce.Do(func() {
		serviceInstance = &serviceImpl{}
	})
	return serviceInstance
}
