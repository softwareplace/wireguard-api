package shared

import (
	"flag"
)

type Args struct {
	Profile     bool
	Server      bool
	Name        string
	Host        string
	ApiKey      string
	Description string
	Command     string
	H           bool
	Rm          bool
	Add         bool
	Ls          bool
	N           bool
	Use         bool
	All         bool
}

func Load() *Args {
	args := new(Args)
	args.parse()
	return args
}

func (a *Args) parse() {

	command := flag.CommandLine.Name()

	profile := flag.Bool("profile", false, "Run in profile mode.")
	server := flag.Bool("server", false, "Run server process.")
	h := flag.Bool("help", false, "Display help information.")
	n := flag.Bool("n", false, "Display only the resource name.")
	ls := flag.Bool("ls", false, "Lists all from a resource.")
	add := flag.Bool("add", false, "Adds a new resource.")
	name := flag.String("name", "", "Name of the resource.")
	description := flag.String("description", "", "The description of the resource.")
	apiKey := flag.String("api-key", "", "The API key of the resource.")
	host := flag.String("host", "", "Expected host.")
	rm := flag.Bool("rm", false, "Remove an existing resource.")
	use := flag.Bool("use", false, "Use a specific resource.")
	all := flag.Bool("all", false, "Apply to all resources.")

	flag.Parse()

	a.Profile = *profile
	a.Server = *server
	a.H = *h
	a.Ls = *ls
	a.Add = *add
	a.Name = *name
	a.Host = *host
	a.Rm = *rm
	a.Use = *use
	a.N = *n
	a.Description = *description
	a.ApiKey = *apiKey
	a.Command = command
	a.All = *all

}
