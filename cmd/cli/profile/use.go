package profile

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

	profile := config.FindProfile(args.Name)
	if profile == nil {
		log.Fatalf("Profile %s not found", args.Name)
	}

	config.CurrentProfile = args.Name

	shared.SaveConfig(config)
}

func (api *apiIml) usage(exit int) {
	fmt.Println("Usage of use profile:")
	fmt.Println("  --name - The name of the profile to use.")
	fmt.Println("  --help - View help information.")
	os.Exit(exit)
}
