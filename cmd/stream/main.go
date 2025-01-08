package main

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/stream/parse"
	"github.com/softwareplace/wireguard-api/cmd/stream/spec"
	"os"
	"os/exec"
)

func main() {
	streamEnv := spec.New()

	cmd := exec.Command(streamEnv.WgCommand, streamEnv.Args...)

	// Run the command and capture output
	output, err := cmd.CombinedOutput()

	if err != nil {
		fmt.Printf("Error executing command: %v\n", err)
		return
	}

	dump, err := parse.WgDump(string(output))

	if err != nil {
		fmt.Printf("Error parsing wg dump: %v\n", err)
		os.Exit(1)
	}

	// Print parsed structs
	for _, line := range dump {
		fmt.Printf("%+v\n", line)
	}
}
