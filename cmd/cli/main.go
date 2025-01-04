package main

import (
	"github.com/softwareplace/wireguard-api/cmd/cli/profile"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"log"
	"os"
)

func usage() {
	log.Printf("\nUsage of %s:\n", os.Args[0])
	log.Println("--profile")
	log.Println("--setup")
	log.Println("--help")
}

func main() {
	log.SetFlags(0)
	log.Println("")

	args := shared.Load()

	if args.Profile {
		profile.Run(args)
		os.Exit(0)
	} else if args.Setup {
		log.Println("Running setup process...")
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
