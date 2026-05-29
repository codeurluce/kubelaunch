package cluster

import (
	"os/exec"
	"strings"
)

type Provider string

const (
	Kind          Provider = "kind"
	Minikube      Provider = "minikube"
	DockerDesktop Provider = "docker-desktop"
	Cloud         Provider = "cloud"
	Unknown       Provider = "unknown"
)

func DetectProvider() Provider {

	out, err := exec.Command(
		"kubectl",
		"config",
		"current-context",
	).Output()

	if err != nil {
		return Unknown
	}

	ctx := strings.ToLower(strings.TrimSpace(string(out)))

	switch {

	case strings.Contains(ctx, "kind"):
		return Kind

	case strings.Contains(ctx, "minikube"):
		return Minikube

	case strings.Contains(ctx, "docker-desktop"):
		return DockerDesktop

	case strings.Contains(ctx, "gke"),
		strings.Contains(ctx, "eks"),
		strings.Contains(ctx, "aks"):
		return Cloud

	default:
		return Unknown
	}
}
