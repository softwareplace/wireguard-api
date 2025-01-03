package peer

import "github.com/eliasmeireles/wireguard-api/pkg/models"

func (s *serviceImpl) GetAvailablePeer() (*models.Peer, error) {
	return s.repository().GetAvailablePeer()
}
