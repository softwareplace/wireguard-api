package spec

import (
	"log"
	"os"
	"strings"
)

type StreamEnv struct {
	ApiKey        string
	Authorization string
	Server        string
	WgCommand     string
	Args          []string
}

func New() *StreamEnv {
	return &StreamEnv{
		ApiKey:        getEnvOrFatal("WG_API_KEY"),
		Authorization: getEnvOrFatal("WG_API_AUTHORIZATION"),
		Server:        getEnvOrFatal("WG_API_SERVER"),
		WgCommand:     getEnvOrFatal("WG_API_COMMAND"),
		Args:          strings.Split(getEnvOrFatal("WG_API_COMMAND_ARGS"), ","),
	}
}

func getEnvOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}
