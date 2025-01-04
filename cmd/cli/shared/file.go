package shared

import (
	"github.com/softwareplace/wireguard-api/cmd/cli/spec"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func LoadConfig() *spec.Config {
	// Open the file
	file, err := os.Open(ContextPath + "conf.yaml")
	if os.IsNotExist(err) {
		defaultConfig := []byte("appVersion: v1\ncurrent-server: \"\"\ncurrent-profile: \"\"\nservers: []\nprofiles: []")
		if err := os.WriteFile(ContextPath+"conf.yaml", defaultConfig, 0644); err != nil {
			log.Fatal("failed to create default config file: %w", err)
		}
		file, err = os.Open(ContextPath + "conf.yaml")
	}
	if err != nil {
		log.Fatal("failed to open config file: %w", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Failed to close file: ", err)
		}
	}(file)

	// Decode the YAML into the Config struct
	var config spec.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal("failed to decode config YAML: %w", err)
	}

	return &config
}

func SaveConfig(config *spec.Config) {
	// Open the file
	file, err := os.Create(ContextPath + "conf.yaml")
	if err != nil {
		log.Fatal("failed to open config file: %w", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal("Failed to close file: ", err)
		}
	}(file)

	// Encode the Config struct into YAML
	encoder := yaml.NewEncoder(file)
	if err := encoder.Encode(config); err != nil {
		log.Fatal("failed to encode config YAML: %w", err)
	}
}
