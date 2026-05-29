package builder

import (
	"fmt"
	"os/exec"
)

// BuildImage = docker build
func BuildImage(repoPath string, image string) error {

	cmd := exec.Command(
		"docker",
		"build",
		"-t",
		image,
		repoPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker build failed: %s | %v", string(output), err)
	}

	return nil
}

// TagImage = optionnel (si tu veux versioning futur)
func TagImage(image string, tag string) error {

	full := fmt.Sprintf("%s:%s", image, tag)

	cmd := exec.Command(
		"docker",
		"tag",
		image,
		full,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("docker tag failed: %s | %v", string(output), err)
	}

	return nil
}