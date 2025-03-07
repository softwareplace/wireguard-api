package env

import (
	"log"
	"os"
	"sync"
)

type ApplicationEnv struct {
	// ApiSecretAuthorization is the secret used for API authorization. Used to sec and decrypt authorization token claims.
	ApiSecretAuthorization string
	// Port is the port on which the application runs.
	Port string
	// ContextPath is the base path used for API routing.
	ContextPath string
	// PeerResourcePath is the resource path for peer connections.
	PeerResourcePath string
	// ApiSecretKey is the key used for API security.
	ApiSecretKey string
	// DBEnv holds the database environment configuration.
	DBEnv        DBEnv
	InitFilePath string
}

type DBEnv struct {
	// DatabaseName is the name of the MongoDB database.
	DatabaseName string
	// Username is the username for MongoDB authentication.
	Username string
	// Password is the password for MongoDB authentication.
	Password string
	// Uri is the connection URI for MongoDB.
	Uri string
}

var (
	instance   *ApplicationEnv
	appEnvOnce sync.Once
)

func AppEnv() ApplicationEnv {
	if os.Getenv("DEBUG_MODE") == "true" {
		log.SetFlags(log.LstdFlags | log.Llongfile)
	}

	appEnvOnce.Do(func() {
		if instance == nil {

			dbEnv := DBEnv{
				DatabaseName: GetRequiredEnv("MONGO_DATABASE"), // required
				Username:     GetRequiredEnv("MONGO_USERNAME"), // required
				Password:     GetRequiredEnv("MONGO_PASSWORD"), // required
				Uri:          GetRequiredEnv("MONGO_URI"),      // required
			}

			instance = &ApplicationEnv{
				ApiSecretAuthorization: GetRequiredEnv("API_SECRET_AUTHORIZATION"), // required
				Port:                   getServerPort(),
				ContextPath:            getServerContextPath(),
				PeerResourcePath:       getPeerResourcePath(),
				ApiSecretKey:           GetRequiredEnv("API_SECRET_KEY"),
				InitFilePath:           os.Getenv("API_INIT_FILE"),
				DBEnv:                  dbEnv,
			}
		}
	})

	return *instance
}

func getPeerResourcePath() string {
	if peersResourcePath := os.Getenv("PEERS_RESOURCE_PATH"); peersResourcePath != "" {
		return peersResourcePath
	}
	return "/etc/wireguard/"
}

func getServerPort() string {
	if envPort := os.Getenv("PORT"); envPort != "" {
		return envPort
	}
	return "1080"
}

func getServerContextPath() string {
	if contextPath := os.Getenv("CONTEXT_PATH"); contextPath != "" {
		return contextPath
	}
	return "/api/private-network/v1/"
}

func GetRequiredEnv(key string) string {
	envValue := os.Getenv(key)
	if envValue == "" {
		panic(key + " environment variable is required")
	}
	return envValue
}
