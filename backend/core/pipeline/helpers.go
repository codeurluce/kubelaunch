package pipeline

import (
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func cloneRepo(repoURL, branch string) (string, error) {
	path := "/tmp/kubelaunch-" + randomString()

	cmd := exec.Command("git", "clone", repoURL, path)

	if err := cmd.Run(); err != nil {
		return "", err
	}

	return path, nil
}

func randomString() string {
	const letters = "abcdefghijklmnopqrstuvwxyz0123456789"

	rand.Seed(time.Now().UnixNano())

	b := make([]byte, 8)

	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func scanFiles(root string) []string {
	var files []string

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files
}

func dockerBuild(path, image string) error {
	cmd := exec.Command("docker", "build", "-t", image, path)
	return cmd.Run()
}

func kindLoad(image string) error {
	cmd := exec.Command("kind", "load", "docker-image", image)
	return cmd.Run()
}

func kubectlApply(yaml string) error {
	cmd := exec.Command("kubectl", "apply", "-f", "-")

	cmd.Stdin = strings.NewReader(yaml)

	return cmd.Run()
}

func writeFile(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}
