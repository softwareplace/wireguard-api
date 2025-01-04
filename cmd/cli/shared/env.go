package shared

import (
	"log"
	"os"
	"strings"
)

var ContextPath = contextPath()

func contextPath() string {
	if contextPath := os.Getenv("CONTEXT_PATH"); contextPath != "" {
		if !strings.HasSuffix(contextPath, "/") {
			return contextPath + "/"
		}
		return contextPath
	}
	if _, statErr := os.Stat(os.Getenv("HOME") + "/.wg-cli/"); os.IsNotExist(statErr) {
		err := os.MkdirAll(os.Getenv("HOME")+"/.wg-cli/", os.ModePerm)
		if err != nil {
			log.Fatal("Failed to create directory: ", err)
		}
	}
	return os.Getenv("HOME") + "/.wg-cli/"
}
