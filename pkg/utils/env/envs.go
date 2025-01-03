package env

import "os"

func GetPeerResourcePath() string {
	peersResourcePath := os.Getenv("PEERS_RESOURCE_PATH")
	if peersResourcePath == "" {
		return "/etc/wireguard/"
	}
	return peersResourcePath
}
