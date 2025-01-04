package server

import (
	"fmt"
	"os"
)

func (api *apiIml) help() {
	fmt.Println("Server usage:")
	fmt.Println("  -ls - Lists all the servers currently managed by the handler.")
	fmt.Println("  -add - Adds a new server.")
	fmt.Println("  -use - Selects a specific server for use.")
	fmt.Println("  -rm - Removes an existing server.")
	fmt.Println("  -help - Displays information about the available commands.")
	os.Exit(0)
}
