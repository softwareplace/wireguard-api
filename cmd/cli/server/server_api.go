package server

import (
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"os"
)

type Handler interface {
	servers(*shared.Args)
	use(args *shared.Args)
	add(*shared.Args)
	remove(args *shared.Args)
	help()
}

type apiIml struct {
}

func Run(args *shared.Args) {
	api := new(apiIml)

	switch {
	case args.Ls:
		api.servers(args)
		os.Exit(0)
	case args.Add:
		api.add(args)
		os.Exit(0)
	case args.Use:
		api.use(args)
		os.Exit(0)
	case args.Rm:
		api.remove(args)
		os.Exit(0)
	default:
		api.help()
	}
}
