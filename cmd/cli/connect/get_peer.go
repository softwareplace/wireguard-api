package connect

import (
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/request"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"net/http"
)

func GetPeer(profile *spec.Profile, server *spec.Server) models.Peer {
	api := request.NewService()

	apiConfig := request.Build(server.Host).
		WithPath("/peers").
		WithHeader(api_context.XApiKey, server.ApiKey).
		WithHeader("Authorization", profile.AuthorizationKey).
		WithExpectedStatusCode(http.StatusOK)

	_, err := api.Get(apiConfig)

	if err != nil {
		log.Fatalf("Failed to get peer: %v", err)
	}

	peer := models.Peer{}
	err = api.BodyDecode(&peer)

	if err != nil {
		log.Fatalf("Failed to decode peer: %v", err)
	}

	return peer
}
