package pipeline

import (
	"os"
	"path/filepath"
)

func CreateWorkspace(appName string) (string, error) {

	path := filepath.Join("tmp", appName)

	err := os.MkdirAll(path, 0755)

	if err != nil {
		return "", err
	}

	return path, nil
}

func WriteFile(path string, content string) error {

	return os.WriteFile(path, []byte(content), 0644)
}
