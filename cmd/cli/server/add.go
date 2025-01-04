package server

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/model"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"log"
	"os"
)

func (api *apiIml) add(args *shared.Args) {

	if args.H {
		addUsage(0)
	}

	if args.Name == "" || args.Host == "" {
		addUsage(1)
	}

	config := shared.LoadConfig()

	server := config.FindServer(args.Name)
	if server != nil {
		log.Fatalf("Server %s already exists", args.Name)
	}

	config.Servers = append(config.Servers, model.Server{
		Name:        args.Name,
		Host:        args.Host,
		Description: args.Description,
	})

	shared.SaveConfig(config)

	log.Printf("server %s added successfully\n", args.Name)
}

func addUsage(exit int) {
	fmt.Println("Add server usage:")
	fmt.Println("  --name (required) - The name of the server to add.")
	fmt.Println("  --host (required) - THe host of the server to add.")
	fmt.Println("  --description - The description of the server to add.")
	fmt.Println("  --help - View help information.")
	os.Exit(exit)
}
