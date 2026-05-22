package k8s

import (
	"fmt"
	"os/exec"
)

// GetPodLogs récupère les logs d’un pod
func GetPodLogs(podName string, namespace string) (string, error) {

	cmd := exec.Command("kubectl", "logs", podName, "-n", namespace)

	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to get logs: %s", string(output))
	}

	return string(output), nil
}
