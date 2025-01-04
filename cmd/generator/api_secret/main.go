package main

import (
	"crypto/ecdsa"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"github.com/atotto/clipboard"
	"github.com/softwareplace/wireguard-api/pkg/domain/db"
	"github.com/softwareplace/wireguard-api/pkg/domain/repository/api_secret"
	"github.com/softwareplace/wireguard-api/pkg/domain/service/security"
	"github.com/softwareplace/wireguard-api/pkg/models"
	"log"
	"os"
	"time"
)

func main() {

	// Update usage to provide detailed help documentation
	flag.Usage = func() {
		log.SetFlags(0) // Disable time and date in log output
		log.Printf("\nUsage of %s:\n", os.Args[0])
		log.Println("  --client <value>	Client information for which the public key is generated (required)")
		log.Println("  --exp <value>	   Expiration time of the generated key in hours (must be a positive integer, required)")
		log.Println("Example:")
		log.Printf("  %s --client=\"exampleClient\" --exp=24\n", os.Args[0])
	}

	if secretKey := os.Getenv("API_SECRET_KEY"); secretKey != "" {
		// Define flags for the script
		clientInfo := flag.String("client", "", "Client information for which the public key is generated")
		expirationHours := flag.Int("exp", 0, "Expiration time of the generated key (in hours)")

		// Parse command-line flags
		flag.Parse()

		// Show usage if help is requested or if required flags are not provided
		if len(os.Args) < 2 || *clientInfo == "" || *expirationHours <= 0 {
			flag.Usage()
			os.Exit(1)
		}

		log.Printf("Generating public key with expiration (hours): %d for client: %s\n", *expirationHours, *clientInfo)

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

		// Generate and log the corresponding public key
		var publicKeyBytes []byte
		switch key := privateKey.(type) {
		case *ecdsa.PrivateKey:
			log.Println("Loaded ECDSA private key successfully")
			publicKeyBytes, err = x509.MarshalPKIXPublicKey(&key.PublicKey)
			if err != nil {
				log.Fatalf("Failed to marshal ECDSA public key: %s", err.Error())
			}
		case *rsa.PrivateKey:
			log.Println("Loaded RSA private key successfully")
			publicKeyBytes, err = x509.MarshalPKIXPublicKey(&key.PublicKey)
			if err != nil {
				log.Fatalf("Failed to marshal RSA public key: %s", err.Error())
			}
		default:
			log.Fatalf("Unsupported private key type: %T", key)
		}

		// Encode the public key in PEM format
		publicKeyPEM := pem.EncodeToMemory(&pem.Block{
			Type:  "PUBLIC KEY",
			Bytes: publicKeyBytes,
		})

		encryptedKey, err := security.GetApiSecurityService().Encrypt(string(publicKeyPEM))

		if err != nil {
			log.Fatalf("Failed to encrypt public key: %s", err)
			return
		}

		db.InitMongoDB()

		apiSecret := models.ApiSecret{
			Key:    encryptedKey,
			Client: *clientInfo,
			Status: "ACTIVE",
		}

		id, err := api_secret.GetRepository().Save(apiSecret)

		if err != nil {
			log.Fatalf("Failed to save api secret: %s, %d", err, id)
			return
		}

		apiJWTInfo := security.ApiJWTInfo{
			Client:     *clientInfo,
			Expiration: time.Duration(*expirationHours),
			Key:        *id,
		}

		apiSecretJWT, err := security.GetApiSecurityService().GenerateApiSecretJWT(apiJWTInfo)

		if err != nil {
			return
		}

		log.Printf("Api Secrte Key generated successfully:\n\n%s\n\n", apiSecretJWT)
		err = clipboard.WriteAll(apiSecretJWT)

		if err != nil {
			log.Fatalf("Failed to add jwt to the clipboard: %s", err)
		} else {
			log.Println("Api Secrete Key copied to the clipboard")
		}
	} else {
		log.Fatal("Generate public key failed: API_SECRET_KEY environment variable is required")
	}
}
