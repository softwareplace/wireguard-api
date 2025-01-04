package connect

import (
	"encoding/json"
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"io"
	"log"
	"net/http"
)

func GetPeer(args *shared.Args, profile *spec.Profile, config *spec.Config, server *spec.Server) models.Peer {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/peers", server.Host), nil)
	if err != nil {
		log.Fatalf("Failed to create GET request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(request.XApiKey, server.ApiKey)
	req.Header.Set("Authorization", profile.AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make GET request: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	var peer models.Peer
	if err := json.NewDecoder(resp.Body).Decode(&peer); err != nil {
		log.Fatalf("Failed to parse response body: %v", err)
	}

	return peer
}
