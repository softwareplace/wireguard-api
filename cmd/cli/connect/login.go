package connect

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"golang.org/x/crypto/ssh/terminal"
	"io"
	"log"
	"net/http"
	"strconv"
	"syscall"
	"time"
)

type LoginResponse struct {
	AccessToken string `json:"token"`
	Expires     string `json:"expires"`
}

func userAuthenticate(args *shared.Args, profile *spec.Profile, config *spec.Config, server *spec.Server) {
	if profile.AuthorizationKey != "" && profile.Expires != "" {
		now := time.Now()

		timestamp, err := strconv.ParseInt(profile.Expires, 10, 64)

		if err != nil {
			fmt.Println("Failed to check current expiration token")
		} else {
			timestamp := time.Unix(timestamp, 0)
			if timestamp.After(now) {
				return
			}
		}
	}

	config.RemoveProfile(profile.Name)
	Login(args, profile, *server)
	config.Profiles = append(config.Profiles, *profile)

	shared.SaveConfig(config)
	fmt.Printf("Login successful for %s\n", profile.Name)
}

func getPassword(args *shared.Args) string {
	if args.Pass != "" {
		return args.Pass
	}

	fmt.Print("Enter password for quest new authorization key: ")
	passwordBytes, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("Failed to read password: %v", err)
	}
	fmt.Println() // Print a newline to separate password input from other logs
	password := string(passwordBytes)
	passwordEncrypted, err := sec.Encrypt(password, []byte(sec.SampleEncryptKey))
	if err != nil {
		log.Fatalf("Failed to encrypt password: %v", err)
	}
	return passwordEncrypted
}

func Login(args *shared.Args, profile *spec.Profile, server spec.Server) {
	password := getPassword(args)

	reqBody := map[string]string{
		"username": profile.Name,
		"password": password,
	}
	body, err := json.Marshal(reqBody)
	if err != nil {
		log.Fatalf("Failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/login", server.Host), bytes.NewBuffer(body))
	if err != nil {
		log.Fatalf("Failed to create POST request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(request.XApiKey, server.ApiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to make POST request: %v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatalf("Failed to close response body: %v", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	// Parse the response body to the LoginResponse structure.
	var loginResp LoginResponse
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&loginResp); err != nil {
		log.Fatalf("Failed to decode response: %v", err)
	}

	profile.AuthorizationKey = loginResp.AccessToken
	profile.Expires = loginResp.Expires

}
