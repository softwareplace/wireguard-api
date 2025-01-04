package auth

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/api_secret"
	"github.com/softwareplace/wireguard-api/pkg/handlers/request"
	"github.com/softwareplace/wireguard-api/pkg/handlers/shared"
	"github.com/softwareplace/wireguard-api/pkg/utils/env"
	"log"
	"net/http"
	"os"
)

var (
	// apiSecret is expected to be an environment variable, adjust as needed
	apiSecret             any // Replace with logic to fetch from environment variables or similar
	mustValidatePublicKey = false
	appEnv                = env.AppEnv()
)

type ApiSecurityHandler interface {
	InitAPISecretKey()
	Middleware(next http.Handler) http.Handler
	ValidatePublicKey(ctx *request.ApiRequestContext) error
}

type apiSecurityHandlerImpl struct{}

func NewApiSecurityHandler() ApiSecurityHandler {
	return &apiSecurityHandlerImpl{}
}

func (a *apiSecurityHandlerImpl) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Validate the public key
		ctx := request.Build(w, r)

		if err := a.ValidatePublicKey(&ctx); err != nil {
			shared.MakeErrorResponse(w, "You are not allowed to access this resource", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, ctx.Request)
	})
}

// InitAPISecretKey initializes and validates the apiSecret environment variable.
// If apiSecret is provided, it ensures the private key from the specified path can be loaded.
// The application will crash if the private key cannot be loaded.
func (a *apiSecurityHandlerImpl) InitAPISecretKey() {
	if secretKey := appEnv.ApiSecretKey; secretKey != "" {
		// Load private key from the provided secretKey file path
		privateKeyData, err := os.ReadFile(secretKey)
		if err != nil {
			log.Fatalf("Failed to read private key file: %s", err.Error())
		}

		// Decode PEM block from the private key data
		block, _ := pem.Decode(privateKeyData)
		if block == nil || block.Type != "PRIVATE KEY" {
			log.Fatalf("Failed to decode private key PEM block")
		}

		// Parse the private key using ParsePKCS8PrivateKey
		privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
		if err != nil {
			log.Fatalf("Failed to parse private key: %s", err.Error())
		}
		apiSecret = privateKey
		mustValidatePublicKey = true

		switch key := apiSecret.(type) {
		case *ecdsa.PrivateKey:
			log.Println("Loaded ECDSA private key successfully")
		case *rsa.PrivateKey:
			log.Println("Loaded RSA private key successfully")
		default:
			log.Fatalf("Unsupported private key type: %T", key)
		}
	}
}

// ValidatePublicKey validates a given public key (in Base64 format) against the private key (apiSecret).
// This is performed only if mustValidatePublicKey is true.
func (a *apiSecurityHandlerImpl) ValidatePublicKey(ctx *request.ApiRequestContext) error {
	if !mustValidatePublicKey {
		return nil // No validation is required
	}

	// Decode the Base64-encoded public key
	claims, err := apiSecurityService.JWTClaims(ctx)

	if err != nil {
		return err
	}

	ctx.SetApiKeyClaims(claims)

	id, err := apiSecurityService.Decrypt(claims["key"].(string))

	if err != nil {
		return err
	}

	apiAccessSecret, err := api_secret.GetRepository().GetById(id)
	ctx.SetApiKeyId(id)
	if err != nil {
		return err
	}

	// Decode the PEM-encoded public key
	decryptKey, err := apiSecurityService.Decrypt(apiAccessSecret.Key)
	if err != nil {
		return err
	}
	block, _ := pem.Decode([]byte(decryptKey))
	if block == nil || block.Type != "PUBLIC KEY" {
		return fmt.Errorf("failed to decode PEM public key")
	}

	// Parse the public key
	parsedPublicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return fmt.Errorf("failed to parse public key: %w", err)
	}

	switch privateKey := apiSecret.(type) {
	case *ecdsa.PrivateKey:
		// Ensure the type of the public key matches ECDSA
		publicKey, ok := parsedPublicKey.(*ecdsa.PublicKey)
		if !ok {
			return fmt.Errorf("invalid public key type, expected ECDSA")
		}

		// Validate if the public key corresponds to the private key
		privateKeyPubKey := &privateKey.PublicKey
		if publicKey.X.Cmp(privateKeyPubKey.X) != 0 || publicKey.Y.Cmp(privateKeyPubKey.Y) != 0 {
			return fmt.Errorf("public key does not match the private key")
		}
	case *rsa.PrivateKey:
		// Ensure the type of the public key matches RSA
		publicKey, ok := parsedPublicKey.(*rsa.PublicKey)
		if !ok {
			return fmt.Errorf("invalid public key type, expected RSA")
		}

		// Validate if the public key corresponds to the private key
		privateKeyPubKey := &privateKey.PublicKey
		if publicKey.E != privateKeyPubKey.E || publicKey.N.Cmp(privateKeyPubKey.N) != 0 {
			return fmt.Errorf("public key does not match the private key")
		}
	default:
		return fmt.Errorf("unsupported private key type: %T", privateKey)
	}

	return nil
}
