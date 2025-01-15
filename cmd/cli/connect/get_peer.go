package connect

import (
	"github.com/softwareplace/http-utils/request"
	httputilsserver "github.com/softwareplace/http-utils/server"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"net/http"
)

func GetPeer(profile *spec.Profile, server *spec.Server) models.Peer {
	api := request.NewApi(models.Peer{})

	apiConfig := request.Build(server.Host).
		WithPath("/peers").
		WithHeader(httputilsserver.XApiKey, server.ApiKey).
		WithHeader("Authorization", profile.AuthorizationKey).
		WithExpectedStatusCode(http.StatusOK)

	response, err := api.Get(apiConfig)

	if err != nil {
		log.Fatalf("Failed to get peer: %v", err)
	}

	return *response
}
