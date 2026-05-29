package builder

import (
	"fmt"
	"os/exec"
)

// LoadImageToKind loads a Docker image into Kind cluster
func LoadImageToKind(image string, clusterName string) error {
	cmd := exec.Command(
		"kind",
		"load",
		"docker-image",
		image,
		"--name",
		clusterName,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("kind load failed: %s | %v", string(output), err)
	}

	return nil
}
