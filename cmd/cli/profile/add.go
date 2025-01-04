package profile

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"log"
	"os"
)

func (api *apiIml) add(args *shared.Args) {

	if args.H {
		addUsage(0)
	}

	if args.Name == "" {
		addUsage(1)
	}

	config := shared.LoadConfig()

	profile := config.FindProfile(args.Name)
	if profile != nil {
		log.Fatalf("Profile %s already exists", args.Name)
	}

	config.Profiles = append(config.Profiles, spec.Profile{
		Name:        args.Name,
		Description: args.Description,
	})

	shared.SaveConfig(config)

	log.Printf("Profile %s added successfully\n", args.Name)
}

func addUsage(exit int) {
	fmt.Println("Add profile usage:")
	fmt.Println("  --name (required) - The name of the profile to add.")
	fmt.Println("  --description - The description of the profile to add.")
	fmt.Println("  --help - View help information.")
	os.Exit(exit)
}
