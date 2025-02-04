package main

import (
	"fmt"
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/request"
	"github.com/softwareplace/wireguard-api/cmd/stream/parse"
	"github.com/softwareplace/wireguard-api/cmd/stream/spec"
	"log"
	"net/http"
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

	api := request.NewService()
	config := request.Build(streamEnv.Server).
		WithPath("peers/stream").
		WithHeader("Authorization", streamEnv.Authorization).
		WithHeader(api_context.XApiKey, streamEnv.ApiKey).
		WithBody(dump).
		WithExpectedStatusCode(http.StatusOK)

	_, err = api.Post(config)
	if err != nil {
		log.Fatalf("Failed to post stream: %v", err)
	}

	log.Printf("Stream posted successfully")
	os.Exit(0)

}
