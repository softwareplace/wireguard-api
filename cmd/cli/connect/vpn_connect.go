package connect

import (
	"encoding/base64"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"log"
	"os"
)

func Run(args *shared.Args) {
	config := shared.LoadConfig()
	profile := config.GetProfile()
	server := config.GetServer()

	if profile == nil {
		log.Fatalf("Profile not found")
	}

	if server == nil {
		log.Fatalf("Server not found")
	}

	userAuthenticate(args, profile, config, server)

	peer := GetPeer(args, profile, config, server)

	if peer.FileName != "" && peer.PeerData != "" {
		// Decode peer.PeerData from base64
		decodedData, err := base64.StdEncoding.DecodeString(peer.PeerData)
		if err != nil {
			log.Fatalf("Failed to decode PeerData: %v", err)
		}

		// Create the temporary directory if it doesn't exist
		if args.PeerSourceDir == "" {
			args.PeerSourceDir = "/etc/wireguard"
		}

		if err := os.MkdirAll(args.PeerSourceDir, 0755); err != nil {
			log.Fatalf("Failed to create temp directory: %v", err)
		}

		// Write the decoded data to a file
		filePath := args.PeerSourceDir + "/" + "wg0.conf"
		if err := os.WriteFile(filePath, decodedData, 0644); err != nil {
			log.Fatalf("Failed to save file: %v", err)
		}

		log.Printf("Peer data successfully saved to %s", filePath)
	}
}
