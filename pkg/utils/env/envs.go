package env

import (
	"os"
)

type ApplicationEnv struct {
	// ApiSecretAuthorization is the secret used for API authorization. Used to encrypt and decrypt authorization token claims.
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
	DBEnv DBEnv
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

var appEnv *ApplicationEnv

func AppEnv() ApplicationEnv {
	if appEnv == nil {
		dbEnv := DBEnv{
			DatabaseName: os.Getenv("MONGO_DATABASE"), // required
			Username:     os.Getenv("MONGO_USERNAME"), // required
			Password:     os.Getenv("MONGO_PASSWORD"), // required
			Uri:          os.Getenv("MONGO_URI"),      // required
		}

		appEnv = &ApplicationEnv{
			ApiSecretAuthorization: os.Getenv("API_SECRET_AUTHORIZATION"), // required
			Port:                   getServerPort(),
			ContextPath:            getServerContextPath(),
			PeerResourcePath:       getPeerResourcePath(),
			ApiSecretKey:           os.Getenv("API_SECRET_KEY"),
			DBEnv:                  dbEnv,
		}
	}
	return *appEnv
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
