package server

import (
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"log"
	"os"
)

func (api *apiIml) remove(args *shared.Args) {
	if args.H {
		rmUsage(0)
	}

	if args.Name == "" && !args.All {
		rmUsage(1)
	}

	config := shared.LoadConfig()

	if args.All {
		config.Servers = []spec.Server{}
	}

	if args.Name != "" {
		server := config.FindServer(args.Name)
		if server == nil {
			log.Fatalf("Server %s not found", args.Name)
		}

		config.RemoveServer(args.Name)
	}
	shared.SaveConfig(config)
	log.Println("Server(s) removed successfully")
}

func rmUsage(exit int) {
	log.Println("Usage of remove server:")
	log.Println("  --help - View help information.")
	log.Println("  --name - The name of the server to remove.")
	log.Println("  --all - Remove all servers.")
	os.Exit(exit)
}
