package profile

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
		config.Profiles = []spec.Profile{}
	}

	if args.Name != "" {
		profile := config.FindProfile(args.Name)
		if profile == nil {
			log.Fatalf("Profile %s not found", args.Name)
		}

		config.RemoveProfile(args.Name)
	}
	shared.SaveConfig(config)
	log.Println("Profile(s) removed successfully")
}

func rmUsage(exit int) {
	log.Println("Usage of remove profile:")
	log.Println("  --help - View help information.")
	log.Println("  --name - The name of the profile to remove.")
	log.Println("  --all - Remove all profiles.")
	os.Exit(exit)
}
