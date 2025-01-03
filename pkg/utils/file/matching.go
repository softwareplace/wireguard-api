package file

import (
	"os"
	"path/filepath"
	"regexp"
)

func LoadMatchingFiles(dir string, pattern *regexp.Regexp) ([]string, error) {
	var matchingFiles []string

	// Walk through the directory
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Skip directories
		if info.IsDir() {
			return nil
		}
		// Check if the file name matches the pattern
		if pattern.MatchString(info.Name()) {
			matchingFiles = append(matchingFiles, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return matchingFiles, nil
}
