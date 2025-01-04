package connect

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"log"
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

	config.RemoveProfile(profile.Name)
	Login(args, profile, *server)
	config.Profiles = append(config.Profiles, *profile)

	shared.SaveConfig(config)

	fmt.Printf("Loggin in as %s\n", profile.Name)

}
