package peer

import "github.com/softwareplace/wireguard-api/pkg/models"

func (s *serviceImpl) GetAvailablePeer() (response *models.Peer, error error, notfound bool) {
	peer, err := s.repository().GetAvailablePeer()

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			return nil, nil, true
		}
		return nil, err, false
	}
	return peer, nil, false
}
