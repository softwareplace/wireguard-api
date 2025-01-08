package connect

import (
	"fmt"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/http_api"
	"github.com/softwareplace/wireguard-api/pkg/utils/sec"
	"golang.org/x/crypto/ssh/terminal"
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
	passwordBytes, err := terminal.ReadPassword(syscall.Stdin)
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

	api := http_api.NewApi(LoginResponse{})

	config := http_api.Config(server.Host).
		WithPath("/login").
		WithBody(reqBody).
		WithHeader(request.XApiKey, server.ApiKey).
		WithExpectedStatusCode(http.StatusOK)

	loginResp, err := api.Post(config)

	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	profile.AuthorizationKey = loginResp.AccessToken
	profile.Expires = loginResp.Expires
}
