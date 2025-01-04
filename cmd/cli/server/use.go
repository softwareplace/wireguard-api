package server

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"log"
	"os"
)

func (api *apiIml) use(args *shared.Args) {
	if args.H {
		api.usage(0)
	}

	if args.Name == "" {
		api.usage(1)
	}

	config := shared.LoadConfig()

	server := config.FindServer(args.Name)
	if server == nil {
		log.Fatalf("Server %s not found", args.Name)
	}

	config.CurrentServer = args.Name

	shared.SaveConfig(config)

	log.Printf("Server %s is now in use\n", args.Name)
}

func (api *apiIml) usage(exit int) {
	fmt.Println("Server usage:")
	fmt.Println("  --name - The name of the server to use.")
	fmt.Println("  --help - View help information.")
	os.Exit(exit)
}
