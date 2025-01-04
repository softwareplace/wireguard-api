package peer

import "github.com/softwareplace/wireguard-api/pkg/models"

func (s *serviceImpl) GetAvailablePeer() (*models.Peer, error) {
	return s.repository().GetAvailablePeer()
}
