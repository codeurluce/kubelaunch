package k8s

import (
	"fmt"
	"os/exec"
)

// Apply manifests Kubernetes
func Apply(yamlPath string) error {

	cmd := exec.Command("kubectl", "apply", "-f", yamlPath)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("kubectl apply failed: %s", string(output))
	}

	return nil
}

// ApplyRaw YAML string (future version)
func ApplyRaw(yaml string) error {
	fmt.Println("Applying manifest...")
	fmt.Println(yaml)
	return nil
}
