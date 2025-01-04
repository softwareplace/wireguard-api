package profile

import (
	"fmt"
	"os"
)

func (api *apiIml) help() {
	fmt.Println("Profile usage:")
	fmt.Println("  -ls - Lists all the profiles currently managed by the handler.")
	fmt.Println("  -add - Adds a new profile.")
	fmt.Println("  -use - Selects a specific profile for use.")
	fmt.Println("  -rm - Removes an existing profile.")
	fmt.Println("  -help - Displays information about the available commands.")
	os.Exit(0)
}
