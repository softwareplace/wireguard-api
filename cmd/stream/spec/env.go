package spec

import (
	"log"
	"os"
	"strings"
)

type StreamEnv struct {
	ApiServer string
	ApiSecret string
	WgCommand string
	Args      []string
}

func New() *StreamEnv {
	env := &StreamEnv{}
	env.loadFromEnv()
	return env
}

func (env *StreamEnv) loadFromEnv() {
	env.ApiServer = getEnvOrFatal("API_SERVER")
	env.ApiSecret = getEnvOrFatal("API_SECRET")
	env.WgCommand = getEnvOrFatal("WG_COMMAND")
	env.Args = strings.Split(getEnvOrFatal("WG_COMMAND_ARGS"), ",")
}

func getEnvOrFatal(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatalf("Environment variable %s is required but not set", key)
	}
	return value
}
