package connect

import (
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/http_api"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"net/http"
)

func GetPeer(profile *spec.Profile, server *spec.Server) models.Peer {
	api := http_api.NewApi(models.Peer{})

	apiConfig := http_api.Config(server.Host).
		WithPath("/peers").
		WithHeader(request.XApiKey, server.ApiKey).
		WithHeader("Authorization", profile.AuthorizationKey).
		WithExpectedStatusCode(http.StatusOK)

	response, err := api.Get(apiConfig)

	if err != nil {
		log.Fatalf("Failed to get peer: %v", err)
	}

	return *response
}
