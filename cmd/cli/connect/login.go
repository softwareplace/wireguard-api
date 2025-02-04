package connect

import (
	"fmt"
	"github.com/softwareplace/http-utils/api_context"
	"github.com/softwareplace/http-utils/request"
	"github.com/softwareplace/wireguard-api/cmd/cli/shared"
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
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

	api := request.NewService()

	config := request.Build(server.Host).
		WithPath("/login").
		WithBody(reqBody).
		WithHeader(api_context.XApiKey, server.ApiKey).
		WithExpectedStatusCode(http.StatusOK)

	_, err := api.Post(config)

	if err != nil {
		log.Fatalf("Failed to login: %v", err)
	}

	loginResp := LoginResponse{}
	err = api.BodyDecode(&loginResp)

	if err != nil {
		log.Fatalf("Failed to decode login response: %v", err)
	}

	if loginResp.AccessToken == "" || loginResp.Expires == "" {
		log.Fatalf("Failed to get authorization key")
	}

	profile.AuthorizationKey = loginResp.AccessToken
	profile.Expires = loginResp.Expires
}
