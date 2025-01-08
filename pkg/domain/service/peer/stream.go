package peer

import "github.com/softwareplace/wireguard-api/pkg/models"

func (s *serviceImpl) Stream(peers []models.Peer) error {
	return s.repository().SaveAll(peers)
}
