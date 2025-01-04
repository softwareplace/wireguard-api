package peer

import (
	"encoding/base64"
	"github.com/softwareplace/wireguard-api/pkg/models"
	envUtils "github.com/softwareplace/wireguard-api/pkg/utils/env"
	fileUtils "github.com/softwareplace/wireguard-api/pkg/utils/file"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func (s *serviceImpl) Load() {
	peerResourcePath := envUtils.AppEnv().PeerResourcePath

	// Define the pattern for matching file names
	pattern := regexp.MustCompile(`.*peer.*\.conf.yaml`)

	// List all matching files
	files, err := fileUtils.LoadMatchingFiles(peerResourcePath, pattern)

	if err != nil {
		log.Printf("Failed to locate peer configuration files: %v", err)
		return
	}

	var peerConfigs []models.Peer

	for _, file := range files {
		content, err := os.ReadFile(file)
		if err != nil {
			log.Printf("Failed to read peer configuration file %s", file)
			return
		}

		encodedContent := base64.StdEncoding.EncodeToString(content)
		peerConfig := models.Peer{
			PeerData: encodedContent,
			FileName: filepath.Base(file),
			Status:   "AVAILABLE",
		}
		peerConfigs = append(peerConfigs, peerConfig)
	}

	err = s.repository().SaveAll(peerConfigs)

	if err != nil {
		log.Printf("Failed to save peer configurations to the repository: %v", err)
		return
	}
}
