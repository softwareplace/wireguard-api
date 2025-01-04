package main

import (
	"github.com/softwareplace/wireguard-api/cmd/cli/profile"
	"github.com/softwareplace/wireguard-api/cmd/cli/server"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"log"
	"os"
)

func usage() {
	log.Printf("\nUsage of %s:\n", os.Args[0])
	log.Println("--profile - Manage profiles")
	log.Println("--server - Manage servers")
	log.Println("--connect - Connect to the current server")
	log.Println("--help")
}

func main() {
	log.SetFlags(0)
	log.Println("")

	args := shared.Load()

	if args.Profile {
		profile.Run(args)
		os.Exit(0)
	} else if args.Server {
		server.Run(args)
		os.Exit(0)
	} else if args.H {
		usage()
		os.Exit(0)

	} else {
		log.Println("No valid flag provided. Use --help for usage information.")
		usage()
		os.Exit(1)
	}
}
