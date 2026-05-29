package cluster

import (
	"os/exec"
	"strings"
)

const DefaultKindCluster = "kubelaunch"

func EnsureKindCluster() error {

	out, err := exec.Command("kind", "get", "clusters").Output()
	if err != nil {
		return err
	}

	if strings.Contains(string(out), DefaultKindCluster) {
		return nil
	}

	cmd := exec.Command(
		"kind",
		"create",
		"cluster",
		"--name",
		DefaultKindCluster,
	)

	return cmd.Run()
}

func KindClusterName() string {
	return DefaultKindCluster
}
