package main

import (
	"encoding/json"
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/stream/parse"
	"github.com/softwareplace/wireguard-api/cmd/stream/spec"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/http_api"
	"go/types"
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

	// Convert dump to JSON and print it
	jsonDump, err := json.MarshalIndent(dump, "", "  ")
	if err != nil {
		log.Fatalf("Error converting dump to JSON: %v", err)
	}

	// Copy JSON dump to clipboard
	if err := exec.Command("bash", "-c", fmt.Sprintf("echo '%s' | xclip -selection clipboard", string(jsonDump))).Run(); err != nil {
		log.Fatalf("Failed to copy JSON dump to clipboard: %v", err)
	}

	fmt.Println("Dump as JSON:")
	fmt.Println(string(jsonDump))

	api := http_api.NewApi(types.Nil{})
	config := http_api.Config(streamEnv.Server).
		WithPath("peers/stream").
		WithHeader("Authorization", streamEnv.Authorization).
		WithHeader(request.XApiKey, streamEnv.Authorization).
		WithBody(dump).
		WithExpectedStatusCode(http.StatusCreated)

	_, err = api.Post(config)
	if err != nil {
		log.Fatalf("Failed to post stream: %v", err)
	}

	log.Printf("Stream posted successfully")
	os.Exit(0)

}
