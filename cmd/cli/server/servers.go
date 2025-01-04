package server

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"log"
	"os"
)

func (api *apiIml) servers(args *shared.Args) {
	if args.H {
		serversUsage(0)
	}

	config := shared.LoadConfig()

	// If servers is empty, print a message and return
	if len(config.Servers) == 0 {
		log.Println("\nAny server available. You can add a server using the 'add' command.")
		log.Printf("\t%s --server --add --help\n\n", args.Command)
		return
	}

	log.Printf("\nCurrent server: %s\n\n", config.CurrentServer)

	for _, server := range config.Servers {
		if args.N {
			log.Printf("Name: %s\n", server.Name)
		} else {
			log.Printf("Name: %s\n\tAPI Key: %s\n\tDescription: %s\n\t", server.Name, server.ApiKey, server.Description)
		}
	}
}

func serversUsage(exit int) {
	fmt.Println("Usage of view servers:")
	fmt.Println("  -n - View only the name of the servers.")
	fmt.Println("  -h - View help information.")
	os.Exit(exit)
}
