package shared

import (
	"github.com/softwareplace/wireguard-api/cmd/cli/model"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

func LoadConfig() *model.Config {
	// Open the file
	file, err := os.Open(ContextPath + "conf.yaml")
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
	var config model.Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal("failed to decode config YAML: %w", err)
	}

	return &config
}

func SaveConfig(config *model.Config) {
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
