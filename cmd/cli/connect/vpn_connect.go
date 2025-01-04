package connect

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"log"
	"strconv"
	"time"
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
}

func userAuthenticate(args *shared.Args, profile *spec.Profile, config *spec.Config, server *spec.Server) {
	if profile.AccessToken != "" && profile.Expires != "" {
		now := time.Now()

		timestamp, err := strconv.ParseInt(profile.Expires, 10, 64)

		if err != nil {
			fmt.Println("Failed to check current expiration token")
		} else {
			timestamp := time.Unix(timestamp, 0)
			if timestamp.After(now) {
				return
			}
		}
	}

	config.RemoveProfile(profile.Name)
	Login(args, profile, *server)
	config.Profiles = append(config.Profiles, *profile)

	shared.SaveConfig(config)
	fmt.Printf("Loggin in as %s\n", profile.Name)
}
