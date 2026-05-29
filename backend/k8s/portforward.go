package k8s

import (
	"fmt"
	"os/exec"
)

func StartPortForward(
	serviceName string,
	localPort string,
	targetPort string,
) error {

	cmd := exec.Command(
		"kubectl",
		"port-forward",
		"svc/"+serviceName,
		localPort+":"+targetPort,
	)

	// important -> background process
	err := cmd.Start()
	if err != nil {
		return fmt.Errorf("port-forward failed: %w", err)
	}

	return nil
}
