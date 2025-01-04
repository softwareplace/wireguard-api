package profile

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"log"
	"os"
)

func (api *apiIml) profiles(args *shared.Args) {
	if args.H {
		profilesUsage(0)
	}

	config := shared.LoadConfig()

	// If profiles is empty, print a message and return
	if len(config.Profiles) == 0 {
		log.Println("\nAny profile available. You can add a profile using the 'add' command.")
		log.Printf("\t%s --profile --add --help\n\n", args.Command)
		return
	}

	log.Printf("\nCurrent profile: %s\n\n", config.CurrentProfile)

	for _, profile := range config.Profiles {
		if args.N {
			log.Printf("Name: %s\n", profile.Name)
		} else {
			log.Printf("Name: %s\n\tDescription: %s\n\t", profile.Name, profile.Description)
		}
	}
}

func profilesUsage(exit int) {
	fmt.Println("Usage of view profiles:")
	fmt.Println("  -n - View only the name of the profiles.")
	fmt.Println("  -h - View help information.")
	os.Exit(exit)
}
